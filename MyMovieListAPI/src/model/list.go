package model

import (
	"time"
	"encoding/json"
	"database/sql"
	_ "github.com/lib/pq"
	u "util"
	"log"
)

type listItem struct {
	Id 				 int 		   `json:"id"`
	Title 		 string    `json:"title"`
	PosterPath string    `json:"posterPath"`
	Status     string    `json:"status"`
	Runtime    int		   `json:"runtime"`
	Overview   string    `json:"overview"`
	DateAdded  string    `json:"dateAdded"`
}

type List struct {
	Id     int				`json:"id"`
	Name   string 		`json:"name"`
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
		item.DateAdded = time.Now().String()
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

/*
Read a list from the database by id
*/
func ReadList(id int) (List, error) {
	queryStr := `
	SELECT owner, name, public, items FROM lists WHERE
	id=$1
	`
	db, err := sql.Open("postgres", u.ConnStr())
	if err != nil {
		return List{}, err
	}
	l := List{ Id: id }
	var itemsJSON string
	err = db.QueryRow(queryStr, id).Scan(&l.Owner, &l.Name, &l.Public, &itemsJSON)
	err = json.Unmarshal([]byte(itemsJSON), &l.Items)
	return l, err
}

/*
Write a list to the database
*/
func (l List) WriteList() error {
	itemsBytes, err := json.Marshal(l.Items)
	if err != nil {
		return err
	}
	itemsStr := string(itemsBytes[:])
	queryStr := `
	UPDATE lists
	SET name=$1,
	    public=$2,
			items=$3
	WHERE
	id=$4
	`
	db, err := sql.Open("postgres", u.ConnStr())
	if err != nil {
		return err
	}
	_, err = db.Exec(queryStr, l.Name, l.Public, itemsStr, l.Id)
	return err
}

/*
Insert a new list into the database with a unique id and any items the list struct contains
List is always private by default
*/
func (l List) InsertAsNewList() error {

	log.Println(l.Name)
	
	queryStr := `
	INSERT INTO lists (owner, name, items) VALUES
	($1, $2, $3)
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
	_, err = db.Exec(queryStr, l.Owner, l.Name, itemsStr)
	return err
	
}


/*
Delete a List by id
*/
func (l List) Delete() error {

	id := l.Id
	queryStr := `
	DELETE FROM lists WHERE id=$1
	`
	db, err := sql.Open("postgres", u.ConnStr())
		if err != nil {
			return err
	}
	_, err = db.Exec(queryStr, id)
	return err

}
