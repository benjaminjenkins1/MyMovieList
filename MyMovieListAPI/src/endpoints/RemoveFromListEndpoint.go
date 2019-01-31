package endpoints

import (
	"net/http"
	m "model"
	u "util"
	"encoding/json"
	"errors"
)

type RemoveFromListRequest struct {
	Id    int 				 `json: "id"`
	Items []m.ListItem `json: "items"`
}

func RemoveFromListEndpoint(r *http.Request, id string) ([]byte, error) {
	
	removeFromListRequest := RemoveFromListRequest{}
	err := json.NewDecoder(r.Body).Decode(&removeFromListRequest)
	if err != nil {
		return []byte("{}"), errors.New("Remove from list: error decoding request")
	}

	l, err := m.ReadList(removeFromListRequest.Id)
	if err != nil {
		return []byte("{}"), errors.New("Remove from list: error reading list")
	}
	if l.Owner != id {
		return []byte("{}"), errors.New("Remove from list: user does not own list")
	}

	l = m.RemoveItems(removeFromListRequest.Items, l)
	err = l.WriteList()
	u.Check(err)

	listBytes, err := json.Marshal(l)
	return listBytes, err

}