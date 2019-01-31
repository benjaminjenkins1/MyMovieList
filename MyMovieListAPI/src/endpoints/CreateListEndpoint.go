package endpoints

import (
	"net/http"
	m "model"
	u "util"
	"encoding/json"
	"errors"
)

type CreateListRequest struct {
	Name string `json: "name"`
}

func CreateListEndpoint(r *http.Request, id string) ([]byte, error) {

	createListRequest := CreateListRequest{}
	err := json.NewDecoder(r.Body).Decode(&createListRequest)
	if err != nil {
		return []byte("{}"), errors.New("Create list: error decoding request")
	}

	l := m.NewList(id)
	l.Name = createListRequest.Name
	err = l.InsertAsNewList()
	u.Check(err)
	res, err := ListsEndpoint(r, id)
	return res, err
	
}