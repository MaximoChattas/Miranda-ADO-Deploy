package client

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"project/model"
	"testing"
)

func TestInsertUser_Client(t *testing.T) {
	a := assert.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock database")
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		DriverName: "sqlserver",
		Conn:       db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Connection failed to open")
	}

	Db = gormDB
	UserClient = &userClient{}

	user := model.User{
		Id:       1,
		Name:     "John",
		LastName: "Doe",
		Dni:      "123456",
		Email:    "johndoe@email.com",
		Password: "password",
		Role:     "Customer",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`SET IDENTITY_INSERT "users" ON;INSERT INTO "users" ("name","last_name","dni","email","password","role","id") OUTPUT INSERTED."id" VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7);SET IDENTITY_INSERT "users" OFF;`).
		WithArgs(user.Name, user.LastName, user.Dni, user.Email, user.Password, user.Role, user.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	result := UserClient.InsertUser(user)

	a.Equal(user, result)
	a.Equal(1, result.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetUserById_Client(t *testing.T) {
	a := assert.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock database")
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		DriverName: "sqlserver",
		Conn:       db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Connection failed to open")
	}

	Db = gormDB
	UserClient = &userClient{}

	user := model.User{
		Id:       1,
		Name:     "John",
		LastName: "Doe",
		Dni:      "123456",
		Email:    "johndoe@email.com",
		Password: "password",
		Role:     "Customer",
	}

	mock.ExpectQuery(`SELECT * FROM "users" WHERE id = @p1 ORDER BY "users"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY`).
		WithArgs(user.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "last_name", "dni", "email", "password", "role"}).
			AddRow(user.Id, user.Name, user.LastName, user.Dni, user.Email, user.Password, user.Role))

	result := UserClient.GetUserById(user.Id)

	a.Equal(user, result)
	a.Equal(1, result.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetUserByEmail_Client(t *testing.T) {
	a := assert.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock database")
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		DriverName: "sqlserver",
		Conn:       db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Connection failed to open")
	}

	Db = gormDB
	UserClient = &userClient{}

	user := model.User{
		Id:       1,
		Name:     "John",
		LastName: "Doe",
		Dni:      "123456",
		Email:    "johndoe@email.com",
		Password: "password",
		Role:     "Customer",
	}

	mock.ExpectQuery(`SELECT * FROM "users" WHERE email = @p1 ORDER BY "users"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY`).
		WithArgs(user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "last_name", "dni", "email", "password", "role"}).
			AddRow(user.Id, user.Name, user.LastName, user.Dni, user.Email, user.Password, user.Role))

	result := UserClient.GetUserByEmail(user.Email)

	a.Equal(user, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetUsers_Client(t *testing.T) {
	a := assert.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock database")
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		DriverName: "sqlserver",
		Conn:       db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Connection failed to open")
	}

	Db = gormDB
	UserClient = &userClient{}

	users := model.Users{
		model.User{
			Id:       1,
			Name:     "John",
			LastName: "Doe",
			Dni:      "123456",
			Email:    "johndoe@email.com",
			Password: "password",
			Role:     "Customer",
		},

		model.User{
			Id:       2,
			Name:     "Jane",
			LastName: "Doe",
			Dni:      "654321",
			Email:    "janedoe@email.com",
			Password: "password",
			Role:     "Customer",
		},
	}

	mock.ExpectQuery(`SELECT * FROM "users"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "last_name", "dni", "email", "password", "role"}).
			AddRow(users[0].Id, users[0].Name, users[0].LastName, users[0].Dni, users[0].Email, users[0].Password, users[0].Role).
			AddRow(users[1].Id, users[1].Name, users[1].LastName, users[1].Dni, users[1].Email, users[1].Password, users[1].Role))

	result := UserClient.GetUsers()

	a.Equal(users, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
