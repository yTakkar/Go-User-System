package routes

import (
	M "Go-User-System/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/badoux/checkmail"

	// For mysql
	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// UserRegister function
func UserRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Response := make(map[string]interface{})

	username := strings.TrimSpace(r.PostFormValue("username"))
	email := strings.TrimSpace(r.PostFormValue("email"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	passwordAgain := strings.TrimSpace(r.PostFormValue("password_again"))

	mailErr := checkmail.ValidateFormat(email)

	db := M.DB()

	var (
		userCount  int
		emailCount int
	)

	db.QueryRow("SELECT COUNT(id) AS userCount FROM users WHERE username=?", username).Scan(&userCount)
	db.QueryRow("SELECT COUNT(id) AS emailCount FROM users WHERE email=?", email).Scan(&emailCount)

	if username == "" || email == "" || password == "" || passwordAgain == "" {
		Response["mssg"] = "Some values are missing!"
	} else if len(username) < 4 || len(username) > 32 {
		Response["mssg"] = "Username should be between 4 and 32"
	} else if mailErr != nil {
		Response["mssg"] = "Invalid Format!"
	} else if password != passwordAgain {
		Response["mssg"] = "Passwords don't match"
	} else if userCount > 0 {
		Response["mssg"] = "Username already exists!"
	} else if emailCount > 0 {
		Response["mssg"] = "Email already exists!"
	} else {

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		_, insErr := db.Exec(
			"INSERT INTO users(username, email, password) VALUES(?, ?, ?) ",
			username,
			email,
			hash,
		)
		if insErr != nil {
			log.Fatal(insErr)
		}

		Response["mssg"] = "You are now registered!"
		Response["success"] = true

	}

	final, err := json.Marshal(Response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(final)

}

// UserLogin function
func UserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Response := make(map[string]interface{})

	rusername := strings.TrimSpace(r.PostFormValue("username"))
	rpassword := strings.TrimSpace(r.PostFormValue("password"))

	db := M.DB()
	var (
		userCount int
		id        int
		username  string
		password  string
	)

	db.QueryRow("SELECT COUNT(id) AS userCount, id, username, password FROM users WHERE username=?", rusername).Scan(&userCount, &id, &username, &password)

	if rusername == "" || rpassword == "" {
		Response["mssg"] = "Some values are missing!"
	} else if userCount == 0 {
		Response["mssg"] = "Invalid username!"
	} else if encErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(rpassword)); encErr != nil {
		Response["mssg"] = "Invalid password!"
	} else {

		session := M.GetSession(r)
		session.Values["id"] = id
		session.Values["username"] = username
		session.Save(r, w)

		Response["success"] = true
		Response["mssg"] = "You are now logged in"
		Response["user"] = id

	}

	final, err := json.Marshal(Response)
	if err != nil {
		panic(err)
	}

	fmt.Println(userCount)
	w.Header().Set("Content-Type", "application/json")
	w.Write(final)

}
