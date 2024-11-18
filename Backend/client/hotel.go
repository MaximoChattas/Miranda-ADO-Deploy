package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)

type hotelClient struct{}

type hotelClientInterface interface {
	InsertHotel(hotel model.Hotel) model.Hotel
	GetHotelById(id int) model.Hotel
	GetHotels() model.Hotels
	DeleteHotel(hotel model.Hotel) error
	UpdateHotel(hotel model.Hotel) model.Hotel
}

var HotelClient hotelClientInterface

func init() {
	HotelClient = &hotelClient{}
}

func (c hotelClient) InsertHotel(hotel model.Hotel) model.Hotel {

	result := Db.Create(&hotel)

	if result.Error != nil {
		log.Error("Failed to insert hotel.")
		return hotel
	}

	log.Debug("Hotel created:", hotel.Id)
	return hotel
}

func (c hotelClient) GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).Preload("Amenities").Preload("Images").First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func (c hotelClient) GetHotels() model.Hotels {
	var hotels model.Hotels
	Db.Preload("Images").Find(&hotels)

	log.Debug("Hotels: ", hotels)

	return hotels
}

func (c hotelClient) DeleteHotel(hotel model.Hotel) error {

	Db.Model(&hotel).Association("Amenities").Clear()
	Db.Model(&hotel).Association("Images").Clear()

	err := Db.Delete(&hotel).Error

	if err != nil {
		log.Debug("Failed to delete hotel")
	} else {
		log.Debug("Hotel deleted: ", hotel.Id)
	}
	return err
}

func (c hotelClient) UpdateHotel(hotel model.Hotel) model.Hotel {

	var newAmenities model.Amenities

	for _, amenity := range hotel.Amenities {
		newAmenities = append(newAmenities, amenity)
	}
	result := Db.Save(&hotel)

	Db.Model(&hotel).Association("Amenities").Replace(newAmenities)

	if result.Error != nil {
		log.Debug("Failed to update hotel")
		return model.Hotel{}
	}

	log.Debug("Updated hotel: ", hotel.Id)
	return hotel
}
