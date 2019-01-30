package endpoints

import (
	"net/http"
	"errors"
	"encoding/json"
	u "util"
	m "model"
)

type DeleteListRequest struct {
	Id int `json: "id"`
}

func DeleteListEndpoint(r *http.Request, id string) ([]byte, error) {
	
	deleteListRequest := DeleteListRequest{}
	err := json.NewDecoder(r.Body).Decode(&deleteListRequest)
	if err != nil {
		return []byte("{}"), errors.New("Delete list: error decoding request")
	}
	l, err := m.ReadList(deleteListRequest.Id)
	if err != nil {
		return []byte("{}"), errors.New("Delete list: error reading list")
	}
	listOwner := l.Owner
	if listOwner == id {
		err := l.Delete()
		u.Check(err)
		res, err := ListsEndpoint(r, id)
		return res, err
	} else {
		return []byte("{}"), errors.New("Delete list: list not owned by user")
	}
	
}