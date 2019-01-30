package endpoints

import (
	"net/http"
	"encoding/json"
	"errors"
	u "util"
	m "model"
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

	l, err := m.ReadList(listOptionsRequest.Id)
	if err != nil {
		return []byte("{}"), errors.New("List options: error reading list")
	}
	if l.Owner != id {
		return []byte("{}"), errors.New("List options: user does not own list")
	}

	l.Name   = listOptionsRequest.Name
	l.Public = listOptionsRequest.Public

	err = l.WriteList()
	u.Check(err)

	listBytes, err := json.Marshal(l)
	return listBytes, err

}