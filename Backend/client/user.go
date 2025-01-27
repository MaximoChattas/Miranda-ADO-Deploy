package client

import (
	"project/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userClient struct{}

type userClientInterface interface {
	InsertUser(user model.User) model.User
	GetUserById(id int) model.User
	GetUserByEmail(email string) model.User
	GetUsers() model.Users
}

var UserClient userClientInterface

func init() {
	UserClient = userClient{}
}

var Db *gorm.DB

func (c userClient) InsertUser(user model.User) model.User {

	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("Failed to insert user.")
		return user
	}

	log.Debug("User created:", user.Id)
	return user
}

func (c userClient) GetUserById(id int) model.User {
	var user model.User

	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func (c userClient) GetUserByEmail(email string) model.User {
	var user model.User

	Db.Where("email = ?", email).First(&user)
	log.Debug("User: ", user)

	return user
}

func (c userClient) GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}
