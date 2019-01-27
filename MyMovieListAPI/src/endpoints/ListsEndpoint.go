package endpoints

import (
	"net/http"
	//"encoding/json"
	//"errors"
)

func ListsEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	res := []byte("lists")
	return res, err
}