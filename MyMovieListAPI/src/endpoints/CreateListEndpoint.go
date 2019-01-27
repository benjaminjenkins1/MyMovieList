package endpoints

import (
	"net/http"
	//"encoding/json"
	//"errors"
)

func CreateListEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	res := []byte("createlist")
	return res, err
}