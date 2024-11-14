package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"project/dto"
	"project/service"
	"strings"
	"testing"
	"time"
)

type TestReservation struct{}

func init() {
	service.ReservationService = &TestReservation{}
}

func (t TestReservation) InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error) {

	if reservationDto.StartDate == "" {
		return reservationDto, errors.New("error creating reservation")
	}

	reservationDto.Id = 1

	return reservationDto, nil
}

func (t TestReservation) GetReservationById(id int) (dto.ReservationDto, error) {

	if id > 10 {
		return dto.ReservationDto{}, errors.New("reservation not found")
	}

	return dto.ReservationDto{Id: id}, nil
}

func (t TestReservation) GetReservations() (dto.ReservationsDto, error) {

	return dto.ReservationsDto{
		dto.ReservationDto{Id: 1},
		dto.ReservationDto{Id: 2},
	}, nil
}

func (t TestReservation) GetReservationsByUser(userId int) (dto.UserReservationsDto, error) {

	if userId > 10 {
		return dto.UserReservationsDto{}, errors.New("user not found")
	}

	return dto.UserReservationsDto{
		UserId: userId,
		Reservations: dto.ReservationsDto{
			dto.ReservationDto{Id: 1},
			dto.ReservationDto{Id: 2},
		},
	}, nil
}

func (t TestReservation) GetReservationsByUserRange(userId int, startDate string, endDate string) (dto.ReservationsDto, error) {

	rangeStart, _ := time.Parse("02-01-2006 15:04", startDate)
	rangeEnd, _ := time.Parse("02-01-2006 15:04", endDate)

	if rangeStart.After(rangeEnd) {
		return dto.ReservationsDto{}, errors.New("a reservation cant end before it starts")
	}

	return dto.ReservationsDto{dto.ReservationDto{Id: 1, UserId: userId}, dto.ReservationDto{Id: 2, UserId: userId}}, nil

}

func (t TestReservation) GetReservationsByHotel(hotelId int) (dto.HotelReservationsDto, error) {

	if hotelId > 10 {
		return dto.HotelReservationsDto{}, errors.New("hotel not found")
	}

	return dto.HotelReservationsDto{
		HotelId: hotelId,
		Reservations: dto.ReservationsDto{
			dto.ReservationDto{Id: 1},
			dto.ReservationDto{Id: 2},
		},
	}, nil
}

func (t TestReservation) DeleteReservation(id int) error {

	if id > 10 {
		return errors.New("reservation not found")
	}

	return nil
}

func TestInsertReservation_Controller_Error(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.POST("/reserve", InsertReservation)

	body := `{
        "start_date": "",
        "end_date": "",
        "user_id": 0,
        "hotel_id": 0,
        "amount": 0
    }`

	req, err := http.NewRequest(http.MethodPost, "/reserve", strings.NewReader(body))
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"error":"error creating reservation"}`

	a.Equal(http.StatusBadRequest, w.Code)
	a.Equal(expectedResponse, w.Body.String())
}

func TestInsertReservation_Controller_Success(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.POST("/reserve", InsertReservation)

	body := `{
       "start_date": "01-01-2024",
       "end_date": "01-02-2024",
       "user_id": 1,
       "hotel_id": 1,
       "amount": 123
   }`

	req, err := http.NewRequest(http.MethodPost, "/reserve", strings.NewReader(body))
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.ReservationDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.ReservationDto{
		Id:        1,
		StartDate: "01-01-2024",
		EndDate:   "01-02-2024",
		UserId:    1,
		HotelId:   1,
		Amount:    123,
	}

	a.Equal(http.StatusCreated, w.Code)
	a.Equal(expectedResponse, response)
}

func TestGetReservationById_Controller_NotFound(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/reservation/:id", GetReservationById)

	req, err := http.NewRequest(http.MethodGet, "/reservation/400", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"error":"reservation not found"}`

	a.Equal(http.StatusNotFound, w.Code)
	a.Equal(expectedResponse, w.Body.String())
}

func TestGetReservationById_Controller_Found(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/reservation/:id", GetReservationById)

	req, err := http.NewRequest(http.MethodGet, "/reservation/1", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.ReservationDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.ReservationDto{
		Id: 1,
	}

	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, response)
}

func TestGetReservations_Controller(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/reservation", GetReservations)

	req, err := http.NewRequest(http.MethodGet, "/reservation", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.ReservationsDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.ReservationsDto{
		dto.ReservationDto{
			Id: 1,
		},

		dto.ReservationDto{
			Id: 2,
		},
	}

	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, response)
}

func TestGetReservationsByUser_Controller_NotFound(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/user/reservations/:id", GetReservationsByUser)

	req, err := http.NewRequest(http.MethodGet, "/user/reservations/400", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"error":"user not found"}`

	a.Equal(http.StatusNotFound, w.Code)
	a.Equal(expectedResponse, w.Body.String())
}

func TestGetReservationsByUser_Controller_Found(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/user/reservations/:id", GetReservationsByUser)

	req, err := http.NewRequest(http.MethodGet, "/user/reservations/1", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.UserReservationsDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.UserReservationsDto{
		UserId: 1,
		Reservations: dto.ReservationsDto{
			dto.ReservationDto{Id: 1},
			dto.ReservationDto{Id: 2},
		},
	}
	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, response)
}

func TestGetReservationsByUserRange_Controller_Error(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/user/reservations/:id/range", GetReservationsByUserRange)

	req, err := http.NewRequest(http.MethodGet, "/user/reservations/1/range?start_date=01-02-2024+10:00&end_date=01-01-2024+10:00", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"error":"a reservation cant end before it starts"}`

	a.Equal(http.StatusBadRequest, w.Code)
	a.Equal(expectedResponse, w.Body.String())
}

func TestGetReservationsByUserRange_Controller_Success(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/user/reservations/:id/range", GetReservationsByUserRange)

	req, err := http.NewRequest(http.MethodGet, "/user/reservations/1/range?start_date=01-01-2024+10:00&end_date=01-02-2024+10:00", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.ReservationsDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.ReservationsDto{dto.ReservationDto{Id: 1, UserId: 1}, dto.ReservationDto{Id: 2, UserId: 1}}

	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, response)
}

func TestGetReservationsByHotel_Controller_NotFound(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/hotel/reservations/:id", GetReservationsByHotel)

	req, err := http.NewRequest(http.MethodGet, "/hotel/reservations/400", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"error":"hotel not found"}`

	a.Equal(http.StatusNotFound, w.Code)
	a.Equal(expectedResponse, w.Body.String())
}

func TestGetReservationsByHotel_Controller_Found(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/hotel/reservations/:id", GetReservationsByHotel)

	req, err := http.NewRequest(http.MethodGet, "/hotel/reservations/1", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.HotelReservationsDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.HotelReservationsDto{
		HotelId: 1,
		Reservations: dto.ReservationsDto{
			dto.ReservationDto{Id: 1},
			dto.ReservationDto{Id: 2},
		},
	}
	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, response)
}

func TestDeleteReservation_Controller_NotFound(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.DELETE("/reservation/:id", DeleteReservation)

	req, err := http.NewRequest(http.MethodDelete, "/reservation/400", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"error":"reservation not found"}`

	a.Equal(http.StatusBadRequest, w.Code)
	a.Equal(expectedResponse, w.Body.String())
}

func TestDeleteReservation_Controller_Found(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.DELETE("/reservation/:id", DeleteReservation)

	req, err := http.NewRequest(http.MethodDelete, "/reservation/1", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"message":"Reservation deleted"}`

	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, w.Body.String())
}
