package domain

import (
	"errors"
	"fmt"
)

//You can return as many vlues you want but last one has to be error
var (
	users = map[int64]*User{
		123: {Id: uint(123), FirstName: "Anand", LastName: "bhai", email: "anand@github.com"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, error)
}

type userDao struct{}

func (ud *userDao) GetName() string {
	return "user Dao"
}

func (ud *userDao) GetUser(userId int64) (*User, error) {
	fmt.Println("I am called from user dao")
	user := users[userId]
	if user == nil {
		return nil, errors.New("No user found")
	}

	return user, nil
}
