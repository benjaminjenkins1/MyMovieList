package endpoints


import (
	"net/http"
	"encoding/json"
	m "model"
	u "util"
)


func UserEndpoint(r *http.Request, id string) ([]byte, error) {
	var err error
	request_id := r.URL.Query().Get("id")
	if request_id == "" {
		request_id = id	
	}
	user, err := m.ReadUser(request_id)
	u.Check(err)
	if err != nil {
		return nil, err
	}
	if id != request_id && user.Public == false {
		return []byte("{}"), nil
	}
	res, err := json.Marshal(user)
	u.Check(err)
	return res, err
}