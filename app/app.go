package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	//router = gin.New()
}

func StartApp() {
	mapRoutesUsingGin()
	//bindRoutes()
	StartAppInternal()
}

func StartAppInternal() {
	// scope of err is limited in if statement
	// using net/http
	/* if err := http.ListenAndServe(":6061", nil); err != nil {
		panic(err)
	} */

	//using gin-gonic

	if err := router.Run(":6061"); err != nil {
		panic(err)
	}
}
