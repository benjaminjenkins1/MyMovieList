package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	e "endpoints"
	m "model"
	"github.com/kataras/go-sessions"
	"github.com/kataras/go-sessions/sessiondb/redis"
	"github.com/kataras/go-sessions/sessiondb/redis/service"
	"github.com/dghubble/gologin"
	"github.com/dghubble/gologin/facebook"
	"golang.org/x/oauth2"
	facebookOAuth2 "golang.org/x/oauth2/facebook"
)

var redisDB = redis.New(service.Config{
	Addr: service.DefaultRedisAddr,
	Password:    "",
	Database:    "",
	MaxIdle:     0,
	MaxActive: 	 0,
	IdleTimeout: service.DefaultRedisIdleTimeout,
	Prefix:      "",
})

var sess = sessions.New(sessions.Config{Cookie: "mymovielistcookie"})

func main() {

	err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
	}

	sess.UseDatabase(redisDB)

	oauth2Config := &oauth2.Config{
		ClientID:     os.Getenv("FB_CLIENT_ID"),
		ClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		RedirectURL:  "https://mymovielist.benjen.me/facebook/callback",
		Endpoint:     facebookOAuth2.Endpoint,
		Scopes:       []string{"email"},
	}
	stateConfig := gologin.DefaultCookieConfig

	http.HandleFunc("/user",			     		Endpoint(e.UserEndpoint,           http.MethodGet ))
  http.HandleFunc("/updateuser",     	  Endpoint(e.UpdateUserEndpoint,	   http.MethodPost))
	http.HandleFunc("/lists", 				 		Endpoint(e.ListsEndpoint,					 http.MethodGet ))
	http.HandleFunc("/addtolist", 		 		Endpoint(e.AddToListEndpoint,			 http.MethodPost))
	http.HandleFunc("/removefromlist", 		Endpoint(e.RemoveFromListEndpoint, http.MethodPost))
  http.HandleFunc("/createlist", 		 		Endpoint(e.CreateListEndpoint,		 http.MethodPost))
  http.HandleFunc("/deletelist", 		 	  Endpoint(e.DeleteListEndpoint,		 http.MethodPost))
	http.HandleFunc("/listoptions", 	    Endpoint(e.ListOptionsEndpoint,		 http.MethodPost))
	http.HandleFunc("/logout",						logoutHandler)

	http.Handle("/facebook/login", 	  facebook.StateHandler(stateConfig, facebook.LoginHandler(oauth2Config, nil)))
	http.Handle("/facebook/callback", facebook.StateHandler(stateConfig, facebook.CallbackHandler(oauth2Config, fbIssueSession(), nil)))

	serverAddress :=  os.Getenv("HOST") + ":" + os.Getenv("PORT")
	fmt.Printf("Listening on %s...\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))

}


func Endpoint(fn func(*http.Request, string) ([]byte, error), method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// log stuff
		log.Printf("- %s - ENDPOINT - %s - %s", r.RemoteAddr, r.Method, r.URL)

		// check method 
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		//authenticate
		if !isAuthenticated(w, r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		s := sess.Start(w, r)
		id := s.GetString("id")

		// do the thing
		res, err := fn(r, id)

		// write the response
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.Write(res)
	}
}


func isAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	s := sess.Start(w, r)
	if _, err := s.GetBoolean("loggedIn"); err == nil {
		return true	
	}
	return false
}


func fbIssueSession() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		FBUser, err := facebook.UserFromContext(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user := m.NewUser(FBUser.ID)
		user.Username = FBUser.Name
		user.LoginType = "fb"
		err = user.WriteUser()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		s := sess.Start(w, r)
		s.Set("loggedIn", true)
		s.Set("id", user.Id)
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}


func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	sess.Destroy(w, r)
	w.WriteHeader(http.StatusOK)
}