package endpoints

import (
	"net/http"
	//"encoding/json"
	//"errors"
)

func DeleteListEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	res := []byte("deletelist")
	return res, err
}