package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"resthttpservice/services"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
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
}
