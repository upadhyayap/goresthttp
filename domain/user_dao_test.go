package domain

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserUserFound(t *testing.T) {
	user, err := UserDao.GetUser(123)
	if err != nil {
		t.Fail()
		log.Println(err.Error())
	}

	assert.NotNil(t, user)
	assert.Equal(t, user.Id, uint(123))
	assert.Equal(t, user.FirstName, "Anand")
	assert.Equal(t, user.LastName, "bhai")
}

func TestGetUserNoUserFound(t *testing.T) {
	_, err := UserDao.GetUser(124)
	if err == nil {
		t.Fail()
		log.Println(err.Error())
	}

}

/* func ExampleGetUser() {
	UserDao.GetUser(123)
	//Output
	//*User
} */
