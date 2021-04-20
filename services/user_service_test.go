package services

import (
	"log"
	"resthttpservice/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	UserDaoMock userDaoMock
	getuserFunc func(userId int64) (*domain.User, error)
)

type userDaoMock struct{}

func init() {
	domain.UserDao = &userDaoMock{}
}

func (udm *userDaoMock) GetUser(userId int64) (*domain.User, error) {
	log.Println("Mock is called")
	return getuserFunc(userId)
}

func TestUserNotFound(t *testing.T) {
	getuserFunc = func(userId int64) (*domain.User, error) { return nil, nil }
	user, err := UserService.GetUser(12)
	if err != nil {
		t.FailNow()
		log.Println("Error in fetching user")
	}

	if user != nil {
		t.Fail()
		log.Println("Expected no user but found", user.FirstName)
	}
}

func TestUserFound(t *testing.T) {
	getuserFunc = func(userId int64) (*domain.User, error) { return &domain.User{Id: 123, FirstName: "Anand"}, nil }
	user, err := UserService.GetUser(123)
	if err != nil {
		t.FailNow()
		log.Println("Error in fetching user")
	}

	if user == nil {
		t.Fail()
		log.Println("Expected a user but found no user")
	}

	assert.Equal(t, uint(123), user.Id)
	assert.Equal(t, "Anand", user.FirstName)
}
