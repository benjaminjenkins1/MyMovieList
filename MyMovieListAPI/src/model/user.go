package model


import (
	u "util"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)


type User struct {
	Id          string    `json: "id"`
	Username    string    `json: "username"`
	DateCreated time.Time `json: "dateCreated"`
	LastLogin   time.Time `json: "lastLogin"`
	LoginType   string    `json: "losginType"`
	Public      bool      `json: "public"`
}


func NewUser(id string) *User {
	user := &User{Id: id}
	return user
}


func ReadUser(id string) (*User, error) {
	var err error
	queryStr := "SELECT id, username, dateCreated, lastLogin, loginType, public FROM users WHERE id=$1"
	db, err := sql.Open("postgres", u.ConnStr())
	if err != nil {
		return nil, err
	}
	user := &User{}
	err = db.QueryRow(queryStr, id).Scan(&user.Id, &user.Username, &user.DateCreated, &user.LastLogin, &user.LoginType, &user.Public)
	return user, err
}


func (usr *User) WriteUser() error {
	var err error
	queryStr := `
	INSERT INTO users (id, username, loginType) VALUES
	($1, $2, $3)
	ON CONFLICT (id) DO UPDATE
	SET username=$2,
	    lastLogin=NOW(),
			public=$4
	`
	db, err := sql.Open("postgres", u.ConnStr())
	if err != nil {
		return  err
	}
	_, err = db.Exec(queryStr, usr.Id, usr.Username, usr.LoginType, usr.Public)
	return err
}
