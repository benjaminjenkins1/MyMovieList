package endpoints


import (
	"net/http"
	"encoding/json"
	u "util"
	m "model"
)


/*
A user can update their profile privacy
*/
func UpdateUserEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	user, err := m.ReadUser(id)
	u.Check(err)
	newPublic  	:= r.FormValue("public")
	if newPublic == "true" {
		user.Public = true
	} else if newPublic == "false" {
		user.Public = false
	}
	err = user.WriteUser()
	u.Check(err)
	res, err := json.Marshal(user)
	u.Check(err)
	return res, err
}