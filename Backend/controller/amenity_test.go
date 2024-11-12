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
)

type TestAmenity struct{}

func init() {
	service.AmenityService = &TestAmenity{}
}

func (t TestAmenity) InsertAmenity(amenityDto dto.AmenityDto) (dto.AmenityDto, error) {

	if amenityDto.Name == "" {
		return amenityDto, errors.New("error creating amenity")
	}

	amenityDto.Id = 1

	return amenityDto, nil
}

func (t TestAmenity) GetAmenities() (dto.AmenitiesDto, error) {

	return dto.AmenitiesDto{
		dto.AmenityDto{
			Id:   1,
			Name: "Breakfast",
		},

		dto.AmenityDto{
			Id:   2,
			Name: "Pool",
		},
	}, nil
}

func TestInsertAmenity_Controller_Error(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.POST("/amenity", InsertAmenity)

	body := `{
		"name": ""
	}`

	req, err := http.NewRequest(http.MethodPost, "/amenity", strings.NewReader(body))
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	expectedResponse := `{"error":"error creating amenity"}`

	a.Equal(http.StatusBadRequest, w.Code)
	a.Equal(expectedResponse, w.Body.String())
}

func TestInsertAmenity_Controller_Success(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.POST("/amenity", InsertAmenity)

	body := `{
		"name": "Breakfast"
	}`

	req, err := http.NewRequest(http.MethodPost, "/amenity", strings.NewReader(body))
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.AmenityDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.AmenityDto{Id: 1, Name: "Breakfast"}

	a.Equal(http.StatusCreated, w.Code)
	a.Equal(expectedResponse, response)
}

func TestGetAmenities_Controller(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/amenity", GetAmenities)

	req, err := http.NewRequest(http.MethodGet, "/amenity", nil)
	if err != nil {
		log.Fatalf("New request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	var response dto.AmenitiesDto
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedResponse := dto.AmenitiesDto{
		dto.AmenityDto{
			Id:   1,
			Name: "Breakfast",
		},

		dto.AmenityDto{
			Id:   2,
			Name: "Pool",
		},
	}

	a.Equal(http.StatusOK, w.Code)
	a.Equal(expectedResponse, response)
}
