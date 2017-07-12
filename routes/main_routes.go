package routes

import (
	M "Go-User-System/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Home function
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id, username := M.AllSessions(r)
	loggedIn(w, "/welcome", r)
	renderTemplates(w, "index", &Page{"Home", id, username})
}

// Welcome function
func Welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notLoggedIn(w, r)
	renderTemplates(w, "welcome", &Page{"Welcome", nil, nil})
}

// Error function
func Error(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	renderTemplates(w, "error", &Page{"Oops! Something went wrong", nil, nil})
}

// Register function
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notLoggedIn(w, r)
	renderTemplates(w, "register", &Page{"Register to continue!", nil, nil})
}

// Login function
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notLoggedIn(w, r)
	renderTemplates(w, "login", &Page{"Login to continue!", nil, nil})
}

// Profile function
func Profile(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	loggedIn(w, "", r)
	id, username := M.AllSessions(r)
	renderTemplates(w, "profile", &Page{"Profile Page", id, username})
}

// ProfileParamless function
func ProfileParamless(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loggedIn(w, "", r)
	http.Redirect(w, r, "/", http.StatusFound)
}

// Logout function
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loggedIn(w, "", r)
	session := M.GetSession(r)
	delete(session.Values, "id")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
