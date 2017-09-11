package routes

import (
	M "Go-User-System/models"
	"html/template"
	"log"
	"net/http"
)

// Page type as a structure
type Page struct {
	Title    string
	UserID   interface{}
	Username interface{}
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, "Oops, Page not found!!", http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func loggedIn(w http.ResponseWriter, url string, r *http.Request) {
	var URL string
	if url == "" {
		URL = "/login"
	} else {
		URL = url
	}
	id, _ := M.AllSessions(r)
	if id == nil {
		http.Redirect(w, r, URL, http.StatusFound)
	}
}

func notLoggedIn(w http.ResponseWriter, r *http.Request) {
	id, _ := M.AllSessions(r)
	if id != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
