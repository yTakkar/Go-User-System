package main

import (
	R "Go-User-System/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func init() {
	godotenv.Load()
}

func main() {

	router := httprouter.New()

	router.GET("/", R.Home)
	router.GET("/welcome", R.Welcome)
	router.GET("/error", R.Error)
	router.GET("/register", R.Register)
	router.GET("/login", R.Login)
	router.GET("/profile/:user", R.Profile)
	router.GET("/profile", R.ProfileParamless)
	router.GET("/logout", R.Logout)

	router.POST("/user/register", R.UserRegister)
	router.POST("/user/login", R.UserLogin)

	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(os.Getenv("PORT"))
}
