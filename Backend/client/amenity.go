package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)

func InsertAmenity(amenity model.Amenity) model.Amenity {

	result := Db.Create(&amenity)

	if result.Error != nil {
		log.Error("Failed to insert amenity.")
		return amenity
	}

	log.Debug("Amenity created:", amenity.Id)
	return amenity
}

func GetAmenityById(id int) model.Amenity {
	var amenity model.Amenity

	Db.Where("id = ?", id).First(&amenity)
	log.Debug("Amenity: ", amenity)

	return amenity
}

func GetAmenityByName(name string) model.Amenity {
	var amenity model.Amenity

	Db.Where("name = ?", name).First(&amenity)
	log.Debug("Amenity: ", amenity)

	return amenity
}

func GetAmenities() model.Amenities {
	var amenities model.Amenities
	Db.Find(&amenities)

	log.Debug("Amenities: ", amenities)

	return amenities
}
