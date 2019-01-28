package endpoints

import (
	"net/http"
	m "model"
	u "util"
	//"encoding/json"
	//"errors"
)

func CreateListEndpoint(r *http.Request, id string) ([]byte, error) {
	l := m.NewList(id)
	err := l.InsertAsNewList()
	u.Check(err)
	res, err := ListsEndpoint(r, id)
	return res, err
}