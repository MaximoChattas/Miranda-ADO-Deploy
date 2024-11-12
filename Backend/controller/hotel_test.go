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

type TestHotel struct{}

func init() {
	service.HotelService = &TestHotel{}
}

func (t TestHotel) GetHotelById(id int) (dto.HotelDto, error) {

	if id > 10 {
		return dto.HotelDto{}, errors.New("hotel not found")
	}

	return dto.HotelDto{Id: id}, nil
}

func (t TestHotel) GetHotels() (dto.HotelsDto, error) {
	return dto.HotelsDto{dto.HotelDto{Id: 1}, dto.HotelDto{Id: 2}}, nil
}

func (t TestHotel) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {
	hotelDto.Id = 1
	return hotelDto, nil
}

func (t TestHotel) CheckAvailability(hotelId int, startDate time.Time, endDate time.Time) bool {

	if hotelId > 10 {
		return false
	}

	return true
}

func (t TestHotel) CheckAllAvailability(startDate string, endDate string) (dto.HotelsDto, error) {
	reservationStart, _ := time.Parse("02-01-2006 15:04", startDate)
	reservationEnd, _ := time.Parse("02-01-2006 15:04", endDate)

	if reservationStart.After(reservationEnd) {
		return dto.HotelsDto{}, errors.New("a reservation cant end before it starts")
	}
	return dto.HotelsDto{dto.HotelDto{Id: 1}, dto.HotelDto{Id: 2}}, nil
}

func (t TestHotel) DeleteHotel(id int) error {

	if id > 10 {
		return errors.New("hotel not found")
	}

	return nil
}

func (t TestHotel) UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {

	if hotelDto.Id > 10 {
		return hotelDto, errors.New("hotel not found")
	}

	return hotelDto, nil
}

func TestInsertHotel_Controller(t *testing.T) {
	a := assert.New(t)

	r := gin.Default()
	r.POST("/hotel", InsertHotel)

	body := `{
        "name": "Hotel Test",
        "room_amount": 10,
        "description": "Test hotel description",
        "street_name": "Test Street",
        "street_number": 123,
        "rate": 4.5
    }`

	req, err := http.NewRequest(http.MethodPost, "/hotel", strings.NewReader(body))
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)
	var response dto.HotelDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	a.Equal(http.StatusCreated, w.Code)
	a.NotZero(response.Id)
}

func TestGetHotelById_Controller_NotFound(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/hotel/:id", GetHotelById)

	req, err := http.NewRequest(http.MethodGet, "/hotel/400", nil)

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

func TestGetHotelById_Controller_Found(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/hotel/:id", GetHotelById)

	req, err := http.NewRequest(http.MethodGet, "/hotel/1", nil)

	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.HotelDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.HotelDto{Id: 1}

	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, response)

}

func TestGetHotels_Controller(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/hotel", GetHotels)

	req, err := http.NewRequest(http.MethodGet, "/hotel", nil)

	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.HotelsDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.HotelsDto{dto.HotelDto{Id: 1}, dto.HotelDto{Id: 2}}

	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, response)

}

func TestCheckAllAvailability_Controller_Error(t *testing.T) {
	a := assert.New(t)

	r := gin.Default()
	r.GET("/availability", CheckAllAvailability)

	req, err := http.NewRequest(http.MethodGet, "/availability?start_date=16-06-2023+15:00&end_date=15-06-2023+11:00", nil)

	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	a.Equal(http.StatusBadRequest, w.Code)

	expectedResponse := `{"error":"a reservation cant end before it starts"}`

	a.Equal(expectedResponse, w.Body.String())
}

func TestCheckAllAvailability_Controller_Success(t *testing.T) {
	a := assert.New(t)

	r := gin.Default()
	r.GET("/availability", CheckAllAvailability)

	req, err := http.NewRequest(http.MethodGet, "/availability?start_date=16-06-2023+15:00&end_date=17-06-2023+11:00", nil)

	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.HotelsDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.HotelsDto{dto.HotelDto{Id: 1}, dto.HotelDto{Id: 2}}

	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, response)
}

func TestDeleteHotel_Controller_NotFound(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.DELETE("/hotel/:id", DeleteHotel)

	req, err := http.NewRequest(http.MethodDelete, "/hotel/400", nil)

	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"error":"hotel not found"}`

	a.Equal(http.StatusBadRequest, w.Code)
	a.Equal(expectedResponse, w.Body.String())

}

func TestDeleteHotel_Controller_Found(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.DELETE("/hotel/:id", DeleteHotel)

	req, err := http.NewRequest(http.MethodDelete, "/hotel/1", nil)

	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"message":"Hotel deleted"}`

	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, w.Body.String())

}

func TestUpdateHotel_Controller_NotFound(t *testing.T) {
	a := assert.New(t)

	r := gin.Default()
	r.PUT("/hotel/:id", UpdateHotel)

	body := `{
		"id": 400,
        "name": "Hotel Test",
        "room_amount": 10,
        "description": "Test hotel description",
        "street_name": "Test Street",
        "street_number": 123,
        "rate": 4.5
    }`

	req, err := http.NewRequest(http.MethodPut, "/hotel/400", strings.NewReader(body))
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"error":"hotel not found"}`

	a.Equal(http.StatusBadRequest, w.Code)
	a.Equal(expectedResponse, w.Body.String())
}

func TestUpdateHotel_Controller_Found(t *testing.T) {
	a := assert.New(t)

	r := gin.Default()
	r.PUT("/hotel/:id", UpdateHotel)

	body := `{
		"id": 1,
        "name": "Hotel Test",
        "room_amount": 10,
        "description": "Test hotel description",
        "street_name": "Test Street",
        "street_number": 123,
        "rate": 4.5
    }`

	req, err := http.NewRequest(http.MethodPut, "/hotel/1", strings.NewReader(body))
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := dto.HotelDto{
		Id:           1,
		Name:         "Hotel Test",
		RoomAmount:   10,
		Description:  "Test hotel description",
		StreetName:   "Test Street",
		StreetNumber: 123,
		Rate:         4.5,
	}

	var response dto.HotelDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	a.Equal(http.StatusOK, w.Code)

	a.Equal(expectedResponse, response)
}
