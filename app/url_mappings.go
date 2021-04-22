package app

import (
	"resthttpservice/controllers"
)

func mapRoutesUsingGin() {
	router.GET("/users/:user_id", controllers.GetUser)
}

// binds routers using plain native net/http
func mapRoutes() {
	//http.HandleFunc("/user", controllers.GetUser)
}
