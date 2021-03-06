package endpoints


import (
	"net/http"
	"encoding/json"
	u "util"
	m "model"
)


type UserOptionsRequest struct {
	Public bool `json: "public"`
}

/*
A user can make their profile public or private
*/
func UserOptionsEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	user, err := m.ReadUser(id)
	u.Check(err)
	updateRequest := UserOptionsRequest{}
	err = json.NewDecoder(r.Body).Decode(&updateRequest)
	u.Check(err)
	user.Public = updateRequest.Public
	err = user.WriteUser()
	u.Check(err)
	res, err := json.Marshal(user)
	u.Check(err)
	return res, err
}