package service

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"project/client"
	"project/dto"
	"project/model"
	"testing"
)

type TestUser struct{}

func init() {
	client.UserClient = &TestUser{}
}

func (t TestUser) InsertUser(user model.User) model.User {

	if user.Name == "" {
		user.Id = 0
	} else {
		user.Id = 1
	}

	return user
}

func (t TestUser) GetUserById(id int) model.User {
	var user model.User

	if id > 10 {
		user.Id = 0
	} else {
		user.Id = id
	}

	return user
}

func (t TestUser) GetUserByEmail(email string) model.User {
	var user model.User

	if email == "" {
		user.Id = 0
	} else {
		user.Id = 1
		user.Email = email

		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("password1"), bcrypt.DefaultCost)
		user.Password = string(encryptedPassword)
	}

	return user
}

func (t TestUser) GetUsers() model.Users {

	return model.Users{
		model.User{
			Id:       1,
			Name:     "John",
			LastName: "Doe",
			Dni:      "123456",
			Email:    "johndoe@email.com",
			Password: "password1",
			Role:     "Customer",
		},

		model.User{
			Id:       2,
			Name:     "Jane",
			LastName: "Doe",
			Dni:      "654321",
			Email:    "janedoe@email.com",
			Password: "password2",
			Role:     "Customer",
		},
	}
}

func TestInsertUser_Service_Error(t *testing.T) {

	a := assert.New(t)
	var user dto.UserDto

	_, err := UserService.InsertUser(user)

	expectedResponse := "error creating user"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestInsertUser_Service_Success(t *testing.T) {

	a := assert.New(t)
	user := dto.UserDto{
		Name:     "John",
		LastName: "Doe",
		Dni:      "123456",
		Email:    "johndoe@email.com",
		Password: "password1",
	}

	result, err := UserService.InsertUser(user)

	a.Nil(err)
	a.NotEqual(user, result)

	user.Id = 1
	user.Role = "Customer"
	user.Password = result.Password

	a.Equal(user, result)
}

func TestGetUserById_Service_NotFound(t *testing.T) {

	a := assert.New(t)

	_, err := UserService.GetUserById(12)

	expectedResponse := "user not found"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestGetUserById_Service_Found(t *testing.T) {

	a := assert.New(t)

	result, err := UserService.GetUserById(1)

	expectedResponse := dto.UserDto{Id: 1}

	a.Nil(err)
	a.Equal(expectedResponse, result)
}

func TestGetUsers_Service(t *testing.T) {

	a := assert.New(t)

	result, err := UserService.GetUsers()

	expectedResponse := dto.UsersDto{
		dto.UserDto{
			Id:       1,
			Name:     "John",
			LastName: "Doe",
			Dni:      "123456",
			Email:    "johndoe@email.com",
			Role:     "Customer",
		},

		dto.UserDto{
			Id:       2,
			Name:     "Jane",
			LastName: "Doe",
			Dni:      "654321",
			Email:    "janedoe@email.com",
			Role:     "Customer",
		},
	}

	a.Nil(err)
	a.Equal(expectedResponse, result)
}

func TestUserLogin_Service_NotRegistered(t *testing.T) {

	a := assert.New(t)
	var user dto.UserDto

	_, err := UserService.UserLogin(user)

	expectedResponse := "user not registered"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestUserLogin_Service_IncorrectPassword(t *testing.T) {

	a := assert.New(t)
	user := dto.UserDto{Email: "email@email.com", Password: "password"}

	_, err := UserService.UserLogin(user)

	expectedResponse := "incorrect password"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestUserLogin_Service_Success(t *testing.T) {

	a := assert.New(t)
	user := dto.UserDto{Email: "email@email.com", Password: "password1"}

	result, err := UserService.UserLogin(user)

	expectedResponse := dto.UserDto{Id: 1, Email: user.Email}

	a.Nil(err)
	a.Equal(expectedResponse, result)
}
