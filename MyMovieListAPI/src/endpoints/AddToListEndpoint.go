package endpoints

import (
	"net/http"
	m "model"
	u "util"
	"encoding/json"
	"errors"
)

type AddToListRequest struct {
	Id    int          `json: "id"`
	Items []m.ListItem `json: "items"`
}

func AddToListEndpoint(r *http.Request, id string) ([]byte, error) {
	
	addToListRequest := AddToListRequest{}
	err := json.NewDecoder(r.Body).Decode(&addToListRequest)
	if err != nil {
		return []byte("{}"), errors.New("Add to list: error decoding request")
	}

	l, err := m.ReadList(addToListRequest.Id)
	if err != nil {
		return []byte("{}"), errors.New("Add to list: error reading list")
	}
	if l.Owner != id {
		return []byte("{}"), errors.New("Add to list: user does not own list")
	}

	l = m.AddItems(addToListRequest.Items, l)
	err = l.WriteList()
	u.Check(err)

	listBytes, err := json.Marshal(l)
	return listBytes, err

}