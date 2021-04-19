package app

import (
	"net/http"
	"resthttpservice/controllers"
)

func StartApp() {
	bindRoutes()
	StartAppInternal()
}

func bindRoutes() {
	http.HandleFunc("/user", controllers.GetUser)
}

func StartAppInternal() {
	// scope of err is limited in if statement
	if err := http.ListenAndServe(":6061", nil); err != nil {
		panic(err)
	}
}
