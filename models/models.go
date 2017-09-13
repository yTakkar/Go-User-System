package models

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("It's vvvVery-secret"))

// GetSession to return the session
func GetSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "session")
	if err != nil {
		panic(err)
	}
	return session
}

// AllSessions function to return all the sessions
func AllSessions(r *http.Request) (interface{}, interface{}) {
	session := GetSession(r)
	id := session.Values["id"]
	username := session.Values["username"]
	return id, username
}

// JSON function which returns response as json
func JSON(w http.ResponseWriter, r *http.Request, resp interface{}) {
	final, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(final)
}
