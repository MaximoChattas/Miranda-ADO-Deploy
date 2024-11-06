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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to Create Mock Database")
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		DriverName: "sqlserver",
		Conn:       db, // Use the mocked *sql.DB connection
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Silence GORM logs for tests
	})
	if err != nil {
		t.Fatalf("Connection Failed to Open")
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
	mock.ExpectExec(`SET IDENTITY_INSERT "hotels" ON;INSERT INTO "hotels" $begin:math:text$"name", "room_amount", "description", "street_name", "street_number", "rate", "id"$end:math:text$ OUTPUT INSERTED\."id" VALUES $begin:math:text$\\?, \\?, \\?, \\?, \\?, \\?$end:math:text$;SET IDENTITY_INSERT "hotels" OFF`).
		WithArgs(hotel.Name, hotel.RoomAmount, hotel.Description, hotel.StreetName, hotel.StreetNumber, hotel.Rate).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result := HotelClient.InsertHotel(hotel)

	// Assert the result
	assert.Equal(t, hotel, result)
	assert.Equal(t, 1, hotel.Id)

	// Assert that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
