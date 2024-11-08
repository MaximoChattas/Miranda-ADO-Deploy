package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"project/client"
	"project/dto"
	"project/model"
	"testing"
)

type TestHotel struct{}

func init() {
	client.HotelClient = &TestHotel{}
}

func (t TestHotel) InsertHotel(hotel model.Hotel) model.Hotel {
	if hotel.Name == "" {
		hotel.Id = 0
	} else {
		hotel.Id = 1
	}

	return hotel
}

func (t TestHotel) GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	if id > 10 {
		hotel.Id = 0
	} else {
		hotel.Id = id
	}

	return hotel
}

func (t TestHotel) GetHotels() model.Hotels {

	return model.Hotels{
		model.Hotel{
			Id:           1,
			Name:         "Hotel 1",
			RoomAmount:   10,
			Description:  "Hotel 1 Description",
			StreetName:   "Hotel 1 Street",
			StreetNumber: 10,
			Rate:         10000,
			Amenities:    nil,
			Images:       nil,
		},

		model.Hotel{
			Id:           2,
			Name:         "Hotel 2",
			RoomAmount:   10,
			Description:  "Hotel 2 Description",
			StreetName:   "Hotel 2 Street",
			StreetNumber: 10,
			Rate:         10000,
			Amenities:    nil,
			Images:       nil,
		},
	}

}

func (t TestHotel) DeleteHotel(hotel model.Hotel) error {
	if hotel.Id > 10 {
		return errors.New("failed to delete hotel")
	}

	return nil
}

func (t TestHotel) UpdateHotel(hotel model.Hotel) model.Hotel {

	return hotel
}

func TestInsertHotel_Service_Error(t *testing.T) {

	a := assert.New(t)
	var hotelDto dto.HotelDto

	_, err := HotelService.InsertHotel(hotelDto)

	expectedResponse := "error creating hotel"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())

}

func TestInsertHotel_Service_Success(t *testing.T) {

	a := assert.New(t)
	hotelDto := dto.HotelDto{
		Name: "Hotel",
	}

	result, err := HotelService.InsertHotel(hotelDto)

	hotelDto.Id = 1

	a.Nil(err)
	a.Equal(hotelDto, result)

}

func TestGetHotelById_Service_Found(t *testing.T) {

	a := assert.New(t)

	_, err := HotelService.GetHotelById(1)

	a.Nil(err)
}

func TestGetHotelById_Service_NotFound(t *testing.T) {

	a := assert.New(t)

	_, err := HotelService.GetHotelById(20)

	expectedResponse := "hotel not found"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestGetHotels_Service(t *testing.T) {

	a := assert.New(t)

	result, err := HotelService.GetHotels()

	expectedResult := dto.HotelsDto{
		dto.HotelDto{
			Id:           1,
			Name:         "Hotel 1",
			RoomAmount:   10,
			Description:  "Hotel 1 Description",
			StreetName:   "Hotel 1 Street",
			StreetNumber: 10,
			Rate:         10000,
			Amenities:    nil,
			Images:       nil,
		},

		dto.HotelDto{
			Id:           2,
			Name:         "Hotel 2",
			RoomAmount:   10,
			Description:  "Hotel 2 Description",
			StreetName:   "Hotel 2 Street",
			StreetNumber: 10,
			Rate:         10000,
			Amenities:    nil,
			Images:       nil,
		},
	}

	a.Nil(err)
	a.Equal(expectedResult, result)
}

func TestDeleteHotel_Service_NotFound(t *testing.T) {

	a := assert.New(t)

	hotelId := 12

	err := HotelService.DeleteHotel(hotelId)

	expectedResponse := "hotel not found"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestDeleteHotel_Service_Found(t *testing.T) {

	a := assert.New(t)

	hotelId := 1

	err := HotelService.DeleteHotel(hotelId)

	a.Nil(err)
}

func TestUpdateHotel_Service_NotFound(t *testing.T) {

	a := assert.New(t)

	hotel := dto.HotelDto{
		Id:           12,
		Name:         "Hotel 1",
		RoomAmount:   10,
		Description:  "Hotel 1 Description",
		StreetName:   "Hotel 1 Street",
		StreetNumber: 10,
		Rate:         10000,
		Amenities:    nil,
		Images:       nil,
	}

	_, err := HotelService.UpdateHotel(hotel)

	expectedResult := "hotel not found"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())

}

func TestUpdateHotel_Service_Found(t *testing.T) {

	a := assert.New(t)

	hotel := dto.HotelDto{
		Id:           1,
		Name:         "Hotel 1",
		RoomAmount:   10,
		Description:  "Hotel 1 Description",
		StreetName:   "Hotel 1 Street",
		StreetNumber: 10,
		Rate:         10000,
		Amenities:    nil,
		Images:       nil,
	}

	result, err := HotelService.UpdateHotel(hotel)

	a.Nil(err)
	a.Equal(hotel, result)
}
