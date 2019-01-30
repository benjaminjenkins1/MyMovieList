package endpoints

import (
	"net/http"
	m "model"
	"encoding/json"
	"errors"
)

type ListRequest struct {
	Id int `json: "id"`
}

func ListEndpoint(r *http.Request, id string) ([]byte, error) {
	listRequest := ListRequest{}
	err := json.NewDecoder(r.Body).Decode(&listRequest)
	if err != nil {
		return []byte("{}"), errors.New("Get list: error decoding request")
	}
	l, err := m.ReadList(listRequest.Id)
	if err != nil {
		return []byte("{}"), errors.New("Get list: error reading list")
	}
	listOwner := l.Owner
	if listOwner == id || l.Public == true {
		listBytes, err := json.Marshal(l)
		return listBytes, err
	} else {
		return []byte("{}"), errors.New("Get list: list not owned by user or not public")
	}
}