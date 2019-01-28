package endpoints

import (
	"net/http"
	m "model"
	"encoding/json"
	u "util"
	"database/sql"
	_ "github.com/lib/pq"
	//"errors"
)

func ListsEndpoint(r *http.Request, id string) ([]byte, error) {
	lists, err := getListsForUser(id)
	u.Check(err)
	if err != nil {
		return []byte("{}"), err
	}
	res, err := json.Marshal(lists)
	u.Check(err)
	return res, err
}


/*
Get the lists for a user
*/
func getListsForUser(uid string) ([]m.List, error) {
	queryStr := `
	SELECT id, public, items FROM lists WHERE
	owner=$1
	`
	db, err := sql.Open("postgres", u.ConnStr())
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(queryStr, uid)
	var userLists []m.List
	for rows.Next() {
		var (
			id     int
			public bool
			items  string
		)
		if err := rows.Scan(&id, &public, &items); err != nil {
			return nil, err
		}
		thisList := m.List{ Id:     id,
												Owner:  uid,
			                  Public: public }
		err := json.Unmarshal([]byte(items), &thisList.Items)
		if err != nil {
			return nil, err
		}
		userLists = append(userLists, thisList)
	}
	return userLists, err
}