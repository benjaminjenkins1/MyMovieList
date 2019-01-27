package endpoints

import (
	"net/http"
	//"encoding/json"
	//"errors"
)

func AddToListEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	res := []byte("addtolist")
	return res, err
}