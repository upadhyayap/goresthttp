package domain

import "errors"

//You can return as many avlues you want but last one has to be error
var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Anand", LastName: "bhai", email: "anand@github.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	user := users[userId]
	if user == nil {
		return nil, errors.New("No user found")
	}

	return user, nil
}
