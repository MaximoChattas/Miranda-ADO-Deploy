package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)

type amenityClient struct{}

type amenityClientInterface interface {
	InsertAmenity(amenity model.Amenity) model.Amenity
	GetAmenityById(id int) model.Amenity
	GetAmenityByName(name string) model.Amenity
	GetAmenities() model.Amenities
	//DeleteAmenityById(id string) error
}

var AmenityClient amenityClientInterface

func init() {
	AmenityClient = &amenityClient{}
}

func (c amenityClient) InsertAmenity(amenity model.Amenity) model.Amenity {

	result := Db.Create(&amenity)

	if result.Error != nil {
		log.Error("Failed to insert amenity.")
		return amenity
	}

	log.Debug("Amenity created:", amenity.Id)
	return amenity
}

func (c amenityClient) GetAmenityById(id int) model.Amenity {
	var amenity model.Amenity

	Db.Where("id = ?", id).First(&amenity)
	log.Debug("Amenity: ", amenity)

	return amenity
}

func (c amenityClient) GetAmenityByName(name string) model.Amenity {
	var amenity model.Amenity

	Db.Where("name = ?", name).First(&amenity)
	log.Debug("Amenity: ", amenity)

	return amenity
}

func (c amenityClient) GetAmenities() model.Amenities {
	var amenities model.Amenities
	Db.Find(&amenities)

	log.Debug("Amenities: ", amenities)

	return amenities
}
