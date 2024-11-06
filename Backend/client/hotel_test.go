package client

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm/logger"
	"project/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInsertHotel(t *testing.T) {
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
	HotelClient = &hotelClient{}

	hotel := model.Hotel{
		Id:           1,
		Name:         "Sample Hotel",
		RoomAmount:   10,
		Description:  "Sample description",
		StreetName:   "Sample Street",
		StreetNumber: 123,
		Rate:         4.5,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`SET IDENTITY_INSERT "hotels" ON;INSERT INTO "hotels" ("name","room_amount","description","street_name","street_number","rate","id") OUTPUT INSERTED."id" VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7);SET IDENTITY_INSERT "hotels" OFF;`).
		WithArgs(hotel.Name, hotel.RoomAmount, hotel.Description, hotel.StreetName, hotel.StreetNumber, hotel.Rate, hotel.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	result := HotelClient.InsertHotel(hotel)

	a.Equal(hotel, result)
	a.Equal(1, hotel.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
