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

func TestInsertHotel_Client(t *testing.T) {
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

func TestGetHotelById_Client(t *testing.T) {
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

	mock.ExpectQuery(`SELECT * FROM "hotels" WHERE id = @p1 ORDER BY "hotels"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY`).
		WithArgs(hotel.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "room_amount", "description", "street_name", "street_number", "rate"}).
			AddRow(hotel.Id, hotel.Name, hotel.RoomAmount, hotel.Description, hotel.StreetName, hotel.StreetNumber, hotel.Rate))

	result := HotelClient.GetHotelById(hotel.Id)

	a.Equal(hotel, result)
	a.Equal(1, hotel.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetHotels_Client(t *testing.T) {
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

	hotels := model.Hotels{
		model.Hotel{
			Id:           1,
			Name:         "Sample Hotel",
			RoomAmount:   10,
			Description:  "Sample description",
			StreetName:   "Sample Street",
			StreetNumber: 123,
			Rate:         4.5,
		},

		model.Hotel{
			Id:           2,
			Name:         "Sample Hotel 2",
			RoomAmount:   10,
			Description:  "Sample description 2",
			StreetName:   "Sample Street",
			StreetNumber: 456,
			Rate:         104.5,
		},
	}

	mock.ExpectQuery(`SELECT * FROM "hotels"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "room_amount", "description", "street_name", "street_number", "rate"}).
			AddRow(hotels[0].Id, hotels[0].Name, hotels[0].RoomAmount, hotels[0].Description, hotels[0].StreetName, hotels[0].StreetNumber, hotels[0].Rate).
			AddRow(hotels[1].Id, hotels[1].Name, hotels[1].RoomAmount, hotels[1].Description, hotels[1].StreetName, hotels[1].StreetNumber, hotels[1].Rate))

	result := HotelClient.GetHotels()

	a.Equal(hotels, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestDeleteHotel_Client(t *testing.T) {
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
	mock.ExpectExec(`DELETE FROM "hotel_amenities" WHERE "hotel_amenities"."hotel_id" = @p1`).
		WithArgs(hotel.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "hotels" WHERE "hotels"."id" = @p1`).
		WithArgs(hotel.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = HotelClient.DeleteHotel(hotel)

	a.Nil(err)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestUpdateHotel_Client(t *testing.T) {
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
		Amenities:    model.Amenities{},
		Images:       model.Images{},
	}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "hotels" SET "name"=@p1,"room_amount"=@p2,"description"=@p3,"street_name"=@p4,"street_number"=@p5,"rate"=@p6 WHERE "id" = @p7`).
		WithArgs(hotel.Name, hotel.RoomAmount, hotel.Description, hotel.StreetName, hotel.StreetNumber, hotel.Rate, hotel.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result := HotelClient.UpdateHotel(hotel)

	a.Equal(hotel, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
