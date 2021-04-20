package services

import (
	"log"
	"resthttpservice/domain"
)

type userService struct{}

var UserService userService

func init() {
	UserService = userService{}
}

func (us *userService) GetName() string {
	return "userService"
}

func (us *userService) GetUser(userId int64) (*domain.User, error) {
	log.Println("I am called from", us.GetName())
	return domain.UserDao.GetUser(userId)
}
