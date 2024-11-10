package service

import (
	"github.com/stretchr/testify/assert"
	"project/client"
	"project/dto"
	"project/model"
	"testing"
)

type TestAmenity struct{}

func init() {
	client.AmenityClient = &TestAmenity{}
}

func (t TestAmenity) InsertAmenity(amenity model.Amenity) model.Amenity {
	if amenity.Name == "" {
		amenity.Id = 0
	} else {
		amenity.Id = 1
	}

	return amenity
}

func (t TestAmenity) GetAmenityById(id int) model.Amenity {
	var amenity model.Amenity

	if id > 10 {
		amenity.Id = 0
	} else {
		amenity.Id = id
	}

	return amenity
}

func (t TestAmenity) GetAmenityByName(name string) model.Amenity {
	return model.Amenity{
		Id:   1,
		Name: name,
	}
}

func (t TestAmenity) GetAmenities() model.Amenities {
	return model.Amenities{
		model.Amenity{
			Id:   1,
			Name: "Breakfast",
		},

		model.Amenity{
			Id:   2,
			Name: "Pool",
		},
	}
}

func TestInsertAmenity_Service_Error(t *testing.T) {

	a := assert.New(t)
	var amenity dto.AmenityDto

	_, err := AmenityService.InsertAmenity(amenity)

	expectedResult := "error creating amenity"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())
}

func TestInsertAmenity_Service_Success(t *testing.T) {

	a := assert.New(t)
	amenity := dto.AmenityDto{Name: "Example"}

	result, err := AmenityService.InsertAmenity(amenity)

	expectedResult := dto.AmenityDto{
		Id:   1,
		Name: amenity.Name,
	}

	a.Nil(err)
	a.Equal(expectedResult, result)
}

func TestGetAmenities_Service(t *testing.T) {

	a := assert.New(t)

	expectedResult := dto.AmenitiesDto{
		dto.AmenityDto{
			Id:   1,
			Name: "Breakfast",
		},

		dto.AmenityDto{
			Id:   2,
			Name: "Pool",
		},
	}

	result, err := AmenityService.GetAmenities()

	a.Nil(err)
	a.Equal(expectedResult, result)
}
