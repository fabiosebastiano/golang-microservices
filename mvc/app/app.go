package app

import (
	"net/http"

	"github.com/fabiosebastiano/golang-microservices/mvc/controllers"
)

//StartApplication is the entry point of the app
func StartApplication() {

	http.HandleFunc("/user", controllers.GetUser)
	http.HandleFunc("/users", controllers.GetUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
