package endpoints

import (
	"net/http"
	m "model"
	"encoding/json"
	"errors"
	"strconv"
)

type ListRequest struct {
	Id int `json: "id"`
}

func ListEndpoint(r *http.Request, id string) ([]byte, error) {
	
	listIdString := r.URL.Query()["id"][0]
	
	listId, _ := strconv.Atoi(listIdString)

	l, err := m.ReadList(listId)
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