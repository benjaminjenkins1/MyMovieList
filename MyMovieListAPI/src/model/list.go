package model

import (
	"time"
	"encoding/json"
	"database/sql"
	_ "github.com/lib/pq"
	u "util"
)

type listItem struct {
	Id 				 int 		   `json:"id"`
	Title 		 string    `json:"title"`
	PosterPath string    `json:"posterPath"`
	Status     string    `json:"status"`
	Runtime    int		   `json:"runtime"`
	Overview   string    `json:"overview"`
	DateAdded  time.Time `json:"dateAdded"`
}

type List struct {
	Id     int				`json:"id"`
	Owner  string  		`json:"owner"`
	Public bool				`json:"public"`
	Items  []listItem `json:"items"`
}


func NewList(owner string) List {
	list := List{ Owner: owner }
	return list
}

/*
Add an item to a list's Items if an item with its Id does not already eexist
*/
func AddItems(items []listItem, l List) List {
	for _, item := range items {
		item.DateAdded = time.Now()
	}
	for _, item := range items {
		exists := false
		for _, existingItem := range l.Items {
			if item.Id == existingItem.Id {
				exists = true
			}
		}
		if !exists {
			l.Items = append(l.Items, item)
		}
	}
	return l
}

func ReadList(id int) (List, error) {
	queryStr := `
	SELECT (owner, public, items) FROM lists WHERE
	id=$1
	`
	db, err := sql.Open("postgres", u.ConnStr())
	if err != nil {
		return List{}, err
	}
	l := List{ Id: id }
	var itemsJSON string
	err = db.QueryRow(queryStr, id).Scan(l.Owner, l.Public, itemsJSON)
	err = json.Unmarshal([]byte(itemsJSON), l.Items)
	return l, err
}

func (l List) WriteList() error {
	itemsBytes, err := json.Marshal(l.Items)
	if err != nil {
		return err
	}
	itemsStr := string(itemsBytes[:])
	queryStr := `
	UPDATE lists
	SET public=$1,
			items=$2
	WHERE
	id=$3
	`
	db, err := sql.Open("postgres", u.ConnStr())
	if err != nil {
		return err
	}
	_, err = db.Exec(queryStr, l.Public, itemsStr, l.Id)
	return err
}

/*
Insert a new list into the database with a unique id and any items the list struct contains
List is always private by default
*/
func (l List) InsertAsNewList() error {
	var queryStr string

	if l.Items == nil {

		queryStr = `
		INSERT INTO lists (owner) VALUES
		($1)
		`
		db, err := sql.Open("postgres", u.ConnStr())
		if err != nil {
			return err
		}
		_, err = db.Exec(queryStr, l.Owner)
		return err

	} else {

		queryStr := `
		INSERT INTO lists (owner, items) VALUES
		($1, $2)
		`
		itemsBytes, err := json.Marshal(l.Items)
		if err != nil {
			return err
		}
		itemsStr := string(itemsBytes[:])
		db, err := sql.Open("postgres", u.ConnStr())
		if err != nil {
			return err
		}
		_, err = db.Exec(queryStr, l.Owner, itemsStr)
		return err

	}
	
}
