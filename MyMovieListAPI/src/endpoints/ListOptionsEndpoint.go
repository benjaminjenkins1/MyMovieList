package endpoints

import (
	"net/http"
	//"encoding/json"
	//"errors"
)

func ListOptionsEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	res := []byte("listoptions")
	return res, err
}