package endpoints

import (
	"net/http"
	"encoding/json"
	"errors"
	u "util"
	m "model"
	"log"
)

type ListOptionsRequest struct {
	Id     int    `json: "id"`
	Name   string `json: "name"`
	Public bool   `json: "public"`
}

/*
A user can make a list public or private and change its name
*/
func ListOptionsEndpoint(r *http.Request, id string) ([]byte, error) {
	
	listOptionsRequest := ListOptionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&listOptionsRequest)
	if err != nil {
		return []byte("{}"), errors.New("List options: error decoding request")
	}

	log.Println(listOptionsRequest.Id)
	log.Println(listOptionsRequest.Name)	
	log.Println(listOptionsRequest.Public)

	l, err := m.ReadList(listOptionsRequest.Id)
	if err != nil {
		return []byte("{}"), errors.New("List options: error reading list")
	}

	l.Name   = listOptionsRequest.Name
	l.Public = listOptionsRequest.Public

	err = l.WriteList()
	u.Check(err)

	res, err := ListsEndpoint(r, id)
	return res, err

}