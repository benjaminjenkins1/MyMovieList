package endpoints

import (
	"net/http"
	//"encoding/json"
	//"errors"
)

type UpdateListRequest struct {
	Public bool `json: "public"`
}

/*
A user can make a list public or private
*/
func ListOptionsEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	res := []byte("listoptions")
	return res, err
}