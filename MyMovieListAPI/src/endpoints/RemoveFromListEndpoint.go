package endpoints

import (
	"net/http"
	//"encoding/json"
	//"errors"
)

func RemoveFromListEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	res := []byte("removefromlist")
	return res, err
}