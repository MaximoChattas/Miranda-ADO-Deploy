package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"project/client"
	"project/dto"
	"project/model"
	"testing"
	"time"
)

type TestReservation struct{}

func init() {
	client.ReservationClient = &TestReservation{}
	client.HotelClient = &TestHotel{}
	client.UserClient = &TestUser{}
}

func (t TestReservation) InsertReservation(reservation model.Reservation) model.Reservation {

	if reservation.StartDate == "" {
		reservation.Id = 0
	} else {
		reservation.Id = 1
	}

	return reservation
}

func (t TestReservation) GetReservationById(id int) model.Reservation {

	var reservation model.Reservation

	if id > 10 {
		reservation.Id = 0
	} else {
		reservation.Id = id

		if id == 2 {
			reservation.StartDate = time.Now().Add(96 * time.Hour).Format("02-01-2006 15:04")
		} else if id == 3 {
			reservation.StartDate = time.Now().Add(24 * time.Hour).Format("02-01-2006 15:04")
		}
	}

	return reservation
}

func (t TestReservation) GetReservations() model.Reservations {

	return model.Reservations{
		model.Reservation{
			Id:        1,
			StartDate: "01-01-2024 10:00",
			EndDate:   "01-02-2024 10:00",
			UserId:    1,
			HotelId:   1,
			Amount:    45000,
		},

		model.Reservation{
			Id:        2,
			StartDate: "01-01-2024 10:00",
			EndDate:   "01-02-2024 10:00",
			UserId:    2,
			HotelId:   1,
			Amount:    45000,
		},
	}
}

func (t TestReservation) GetReservationsByUser(userId int) model.Reservations {

	if userId > 10 {
		return model.Reservations{}
	} else {
		return model.Reservations{
			model.Reservation{
				Id:        1,
				StartDate: "01-01-2024 10:00",
				EndDate:   "01-02-2024 10:00",
				UserId:    userId,
				HotelId:   1,
				Amount:    45000,
			},

			model.Reservation{
				Id:        2,
				StartDate: "01-04-2024 10:00",
				EndDate:   "03-04-2024 10:00",
				UserId:    userId,
				HotelId:   1,
				Amount:    10000,
			},
		}
	}
}

func (t TestReservation) GetReservationsByHotel(hotelId int) model.Reservations {

	if hotelId > 10 {
		return model.Reservations{}
	} else {
		return model.Reservations{
			model.Reservation{
				Id:        1,
				StartDate: "01-01-2024 10:00",
				EndDate:   "01-02-2024 10:00",
				UserId:    1,
				HotelId:   hotelId,
				Amount:    45000,
			},

			model.Reservation{
				Id:        2,
				StartDate: "01-01-2024 10:00",
				EndDate:   "01-02-2024 10:00",
				UserId:    1,
				HotelId:   hotelId,
				Amount:    10000,
			},
		}
	}
}

func (t TestReservation) DeleteReservation(reservation model.Reservation) error {

	if reservation.Id > 10 {
		return errors.New("failed to delete reservation")
	}

	return nil
}

func TestInsertReservation_Service_UserNotFound(t *testing.T) {

	a := assert.New(t)

	reservation := dto.ReservationDto{
		StartDate: "01-02-2024 10:00",
		EndDate:   "10-02-2024 10:00",
		UserId:    15,
		HotelId:   1,
	}

	_, err := ReservationService.InsertReservation(reservation)

	expectedResult := "user not found"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())
}

func TestInsertReservation_Service_HotelNotFound(t *testing.T) {

	a := assert.New(t)

	reservation := dto.ReservationDto{
		StartDate: "01-02-2024 10:00",
		EndDate:   "10-02-2024 10:00",
		UserId:    1,
		HotelId:   15,
	}

	_, err := ReservationService.InsertReservation(reservation)

	expectedResult := "hotel not found"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())
}

func TestInsertReservation_Service_Error(t *testing.T) {

	a := assert.New(t)

	reservation := dto.ReservationDto{
		StartDate: "01-02-2024 10:00",
		EndDate:   "10-01-2024 10:00",
		UserId:    1,
		HotelId:   1,
	}

	_, err := ReservationService.InsertReservation(reservation)

	expectedResult := "a reservation cant end before it starts"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())
}

func TestInsertReservation_Service_NotAvailable(t *testing.T) {

	a := assert.New(t)

	reservation := dto.ReservationDto{
		StartDate: "01-01-2024 10:00",
		EndDate:   "10-01-2024 10:00",
		UserId:    1,
		HotelId:   1,
	}

	_, err := ReservationService.InsertReservation(reservation)

	expectedResult := "there are no rooms available"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())
}

func TestInsertReservation_Service_Success(t *testing.T) {

	a := assert.New(t)

	reservation := dto.ReservationDto{
		StartDate: "01-02-2024 10:00",
		EndDate:   "11-02-2024 10:00",
		UserId:    1,
		HotelId:   1,
	}

	result, err := ReservationService.InsertReservation(reservation)

	reservation.Id = 1
	reservation.Amount = 100000

	a.Nil(err)
	a.Equal(reservation, result)
}

func TestGetReservationById_Service_NotFound(t *testing.T) {

	a := assert.New(t)

	_, err := ReservationService.GetReservationById(12)

	expectedResult := "reservation not found"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())
}

func TestGetReservationById_Service_Found(t *testing.T) {

	a := assert.New(t)

	result, err := ReservationService.GetReservationById(1)

	expectedResult := dto.ReservationDto{Id: 1}

	a.Nil(err)
	a.Equal(expectedResult, result)
}

func TestGetReservations_Service(t *testing.T) {

	a := assert.New(t)

	result, err := ReservationService.GetReservations()

	expectedResult := dto.ReservationsDto{
		dto.ReservationDto{
			Id:        1,
			StartDate: "01-01-2024 10:00",
			EndDate:   "01-02-2024 10:00",
			UserId:    1,
			HotelId:   1,
			Amount:    45000,
		},

		dto.ReservationDto{
			Id:        2,
			StartDate: "01-01-2024 10:00",
			EndDate:   "01-02-2024 10:00",
			UserId:    2,
			HotelId:   1,
			Amount:    45000,
		},
	}

	a.Nil(err)
	a.Equal(expectedResult, result)
}

func TestGetReservationsByUser_Service_UserNotFound(t *testing.T) {

	a := assert.New(t)

	_, err := ReservationService.GetReservationsByUser(12)

	expectedResult := "user not found"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())
}

func TestGetReservationsByUser_Service_Success(t *testing.T) {

	a := assert.New(t)

	userId := 1
	result, err := ReservationService.GetReservationsByUser(userId)

	reservations := dto.ReservationsDto{
		dto.ReservationDto{
			Id:        1,
			StartDate: "01-01-2024 10:00",
			EndDate:   "01-02-2024 10:00",
			UserId:    userId,
			HotelId:   1,
			Amount:    45000,
		},

		dto.ReservationDto{
			Id:        2,
			StartDate: "01-04-2024 10:00",
			EndDate:   "03-04-2024 10:00",
			UserId:    userId,
			HotelId:   1,
			Amount:    10000,
		},
	}

	expectedResult := dto.UserReservationsDto{
		UserId:       1,
		Reservations: reservations,
	}

	a.Nil(err)
	a.Equal(expectedResult, result)
}

func TestGetReservationsByUserRange_Service_Error(t *testing.T) {

	a := assert.New(t)

	userId := 1
	startDate := "02-01-2024 10:00"
	endDate := "01-01-2024 10:00"

	_, err := ReservationService.GetReservationsByUserRange(userId, startDate, endDate)

	expectedResponse := "a reservation cant end before it starts"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestGetReservationsByUserRange_Service_NoReservations(t *testing.T) {

	a := assert.New(t)

	userId := 1
	startDate := "02-11-2024 10:00"
	endDate := "03-11-2024 10:00"

	result, err := ReservationService.GetReservationsByUserRange(userId, startDate, endDate)

	var expectedResponse dto.ReservationsDto

	a.Nil(err)
	a.Equal(expectedResponse, result)
}

func TestGetReservationsByUserRange_Service_Success(t *testing.T) {

	a := assert.New(t)

	userId := 1
	startDate := "01-01-2024 00:00"
	endDate := "31-12-2024 23:59"

	result, err := ReservationService.GetReservationsByUserRange(userId, startDate, endDate)

	expectedResponse := dto.ReservationsDto{
		dto.ReservationDto{
			Id:        1,
			StartDate: "01-01-2024 10:00",
			EndDate:   "01-02-2024 10:00",
			UserId:    userId,
			HotelId:   1,
			Amount:    45000,
		},

		dto.ReservationDto{
			Id:        2,
			StartDate: "01-04-2024 10:00",
			EndDate:   "03-04-2024 10:00",
			UserId:    userId,
			HotelId:   1,
			Amount:    10000,
		},
	}

	a.Nil(err)
	a.Equal(expectedResponse, result)
}

func TestGetReservationsByHotel_Service_HotelNotFound(t *testing.T) {

	a := assert.New(t)

	_, err := ReservationService.GetReservationsByHotel(12)

	expectedResult := "hotel not found"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())
}

func TestGetReservationsByHotel_Service_Success(t *testing.T) {

	a := assert.New(t)

	hotelId := 1
	result, err := ReservationService.GetReservationsByHotel(hotelId)

	reservations := dto.ReservationsDto{
		dto.ReservationDto{
			Id:        1,
			StartDate: "01-01-2024 10:00",
			EndDate:   "01-02-2024 10:00",
			UserId:    1,
			HotelId:   hotelId,
			Amount:    45000,
		},

		dto.ReservationDto{
			Id:        2,
			StartDate: "01-01-2024 10:00",
			EndDate:   "01-02-2024 10:00",
			UserId:    1,
			HotelId:   hotelId,
			Amount:    10000,
		},
	}

	expectedResult := dto.HotelReservationsDto{
		HotelId:           hotelId,
		HotelName:         "Hotel 1",
		HotelRoomAmount:   2,
		HotelDescription:  "Hotel 1 Description",
		HotelStreetName:   "Hotel 1 Street",
		HotelStreetNumber: 10,
		HotelRate:         10000,
		Reservations:      reservations,
	}

	a.Nil(err)
	a.Equal(expectedResult, result)
}

func TestDeleteReservation_Service_NotFound(t *testing.T) {

	a := assert.New(t)

	err := ReservationService.DeleteReservation(0)

	expectedResponse := "reservation not found"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestDeleteReservation_Service_Error(t *testing.T) {

	a := assert.New(t)

	err := ReservationService.DeleteReservation(3)

	expectedResponse := "can't delete a reservation 48hs before it starts"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestDeleteReservation_Service_Success(t *testing.T) {

	a := assert.New(t)

	err := ReservationService.DeleteReservation(2)

	a.Nil(err)
}
