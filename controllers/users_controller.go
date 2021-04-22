package controllers

import (
	"log"
	"net/http"
	"resthttpservice/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		// throw 400 Bad request
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	log.Println("user id is ", userId)

	user, err := services.UserService.GetUser(123)
	if err != nil {
		// handle error
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}

	userJson, _ := json.Marshal(user)
	res.Write([]byte(userJson))
} */

// using gin-gonic, controller method should have param *gin.Context
func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		// throw 400 Bad request
		/* res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error())) */
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Println("user id is ", userId)

	user, err := services.UserService.GetUser(userId)
	if err != nil {
		// handle error
		// res.WriteHeader(http.StatusNotFound)
		// res.Write([]byte(err.Error()))
		// return

		c.JSON(http.StatusBadRequest, err)
		return
	}

	// using c.JSON will also set the contentType
	c.JSON(http.StatusOK, user)
	/* userJson, _ := json.Marshal(user)
	res.Write([]byte(userJson)) */
}
