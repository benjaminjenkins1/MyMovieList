package util

import (
	"fmt"
	"os"
)

func ConnStr() string {

	dbname   := os.Getenv("DB_NAME")
	user 		 := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host     := os.Getenv("DB_HOST")
	port     := os.Getenv("DB_PORT")
	connStr  := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s", dbname, user, password, host, port)

	return connStr

}

