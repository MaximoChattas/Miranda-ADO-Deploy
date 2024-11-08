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

func TestInsertReservation_Client(t *testing.T) {
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
	ReservationClient = &reservationClient{}

	reservation := model.Reservation{
		Id:        1,
		StartDate: "10-11-2024 15:00",
		EndDate:   "11-11-2024 11:00",
		UserId:    1,
		HotelId:   1,
		Amount:    20500,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`SET IDENTITY_INSERT "reservations" ON;INSERT INTO "reservations" ("start_date","end_date","user_id","hotel_id","amount","id") OUTPUT INSERTED."id" VALUES (@p1,@p2,@p3,@p4,@p5,@p6);SET IDENTITY_INSERT "reservations" OFF;`).
		WithArgs(reservation.StartDate, reservation.EndDate, reservation.UserId, reservation.HotelId, reservation.Amount, reservation.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	result := ReservationClient.InsertReservation(reservation)

	a.Equal(reservation, result)
	a.Equal(1, result.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetReservationById_Client(t *testing.T) {
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
	ReservationClient = &reservationClient{}

	reservation := model.Reservation{
		Id:        1,
		StartDate: "10-11-2024 15:00",
		EndDate:   "11-11-2024 11:00",
		UserId:    1,
		HotelId:   1,
		Amount:    20500,
	}

	mock.ExpectQuery(`SELECT * FROM "reservations" WHERE id = @p1 ORDER BY "reservations"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY`).
		WithArgs(reservation.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "start_date", "end_date", "user_id", "hotel_id", "amount"}).
			AddRow(reservation.Id, reservation.StartDate, reservation.EndDate, reservation.UserId, reservation.HotelId, reservation.Amount))

	result := ReservationClient.GetReservationById(reservation.Id)

	a.Equal(reservation, result)
	a.Equal(1, result.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetReservations_Client(t *testing.T) {
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
	ReservationClient = &reservationClient{}

	reservations := model.Reservations{
		model.Reservation{
			Id:        1,
			StartDate: "10-11-2024 15:00",
			EndDate:   "11-11-2024 11:00",
			UserId:    1,
			HotelId:   1,
			Amount:    20500,
		},

		model.Reservation{
			Id:        2,
			StartDate: "10-11-2024 15:00",
			EndDate:   "11-11-2024 11:00",
			UserId:    1,
			HotelId:   1,
			Amount:    20500,
		},
	}

	mock.ExpectQuery(`SELECT * FROM "reservations"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "start_date", "end_date", "user_id", "hotel_id", "amount"}).
			AddRow(reservations[0].Id, reservations[0].StartDate, reservations[0].EndDate, reservations[0].UserId, reservations[0].HotelId, reservations[0].Amount).
			AddRow(reservations[1].Id, reservations[1].StartDate, reservations[1].EndDate, reservations[1].UserId, reservations[1].HotelId, reservations[1].Amount))

	result := ReservationClient.GetReservations()

	a.Equal(reservations, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetReservationByUser_Client(t *testing.T) {
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
	ReservationClient = &reservationClient{}

	reservations := model.Reservations{
		model.Reservation{
			Id:        1,
			StartDate: "10-11-2024 15:00",
			EndDate:   "11-11-2024 11:00",
			UserId:    1,
			HotelId:   1,
			Amount:    20500,
		},

		model.Reservation{
			Id:        2,
			StartDate: "10-11-2024 15:00",
			EndDate:   "11-11-2024 11:00",
			UserId:    1,
			HotelId:   1,
			Amount:    20500,
		},
	}

	userId := 1

	mock.ExpectQuery(`SELECT * FROM "reservations" WHERE user_id = @p1`).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "start_date", "end_date", "user_id", "hotel_id", "amount"}).
			AddRow(reservations[0].Id, reservations[0].StartDate, reservations[0].EndDate, reservations[0].UserId, reservations[0].HotelId, reservations[0].Amount).
			AddRow(reservations[1].Id, reservations[1].StartDate, reservations[1].EndDate, reservations[1].UserId, reservations[1].HotelId, reservations[1].Amount))

	result := ReservationClient.GetReservationsByUser(userId)

	a.Equal(reservations, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetReservationByHotel_Client(t *testing.T) {
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
	ReservationClient = &reservationClient{}

	reservations := model.Reservations{
		model.Reservation{
			Id:        1,
			StartDate: "10-11-2024 15:00",
			EndDate:   "11-11-2024 11:00",
			UserId:    1,
			HotelId:   1,
			Amount:    20500,
		},

		model.Reservation{
			Id:        2,
			StartDate: "10-11-2024 15:00",
			EndDate:   "11-11-2024 11:00",
			UserId:    1,
			HotelId:   1,
			Amount:    20500,
		},
	}

	hotelId := 2

	mock.ExpectQuery(`SELECT * FROM "reservations" WHERE hotel_id = @p1`).
		WithArgs(hotelId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "start_date", "end_date", "user_id", "hotel_id", "amount"}).
			AddRow(reservations[0].Id, reservations[0].StartDate, reservations[0].EndDate, reservations[0].UserId, reservations[0].HotelId, reservations[0].Amount).
			AddRow(reservations[1].Id, reservations[1].StartDate, reservations[1].EndDate, reservations[1].UserId, reservations[1].HotelId, reservations[1].Amount))

	result := ReservationClient.GetReservationsByHotel(hotelId)

	a.Equal(reservations, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestDeleteReservation_Client(t *testing.T) {
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
	ReservationClient = &reservationClient{}

	reservation := model.Reservation{
		Id:        1,
		StartDate: "10-11-2024 15:00",
		EndDate:   "11-11-2024 11:00",
		UserId:    1,
		HotelId:   1,
		Amount:    20500,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "reservations" WHERE "reservations"."id" = @p1`).
		WithArgs(reservation.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = ReservationClient.DeleteReservation(reservation)

	a.Nil(err)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
