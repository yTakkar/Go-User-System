package routes

import (
	M "Go-User-System/models"
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
	R := make(map[string]interface{})

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
		R["mssg"] = "Some values are missing!"
	} else if len(username) < 4 || len(username) > 32 {
		R["mssg"] = "Username should be between 4 and 32"
	} else if mailErr != nil {
		R["mssg"] = "Invalid Format!"
	} else if password != passwordAgain {
		R["mssg"] = "Passwords don't match"
	} else if userCount > 0 {
		R["mssg"] = "Username already exists!"
	} else if emailCount > 0 {
		R["mssg"] = "Email already exists!"
	} else {

		hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if hashErr != nil {
			log.Fatal(hashErr)
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

		R["mssg"] = username + ", you can now login!!"
		R["success"] = true

	}

	M.JSON(w, r, R)

}

// UserLogin function
func UserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	R := make(map[string]interface{})

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

	encErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(rpassword))

	if rusername == "" || rpassword == "" {
		R["mssg"] = "Some values are missing!"
	} else if userCount == 0 {
		R["mssg"] = "Invalid username!"
	} else if encErr != nil {
		R["mssg"] = "Invalid password!"
	} else {

		session := M.GetSession(r)
		session.Values["id"] = id
		session.Values["username"] = username
		session.Save(r, w)

		R["success"] = true
		R["mssg"] = "Hello, " + username + "!"
		R["user"] = id

	}

	M.JSON(w, r, R)

}
