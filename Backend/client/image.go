package client

import (
	log "github.com/sirupsen/logrus"
	"project/model"
)

type imageClient struct{}

type imageClientInterface interface {
	InsertImage(image model.Image) model.Image
	InsertImages(images model.Images) model.Images
	GetImageById(id int) model.Image
	GetImages() model.Images
	GetImagesByHotelId(hotelId int) model.Images
}

var ImageClient imageClientInterface

func init() {
	ImageClient = &imageClient{}
}

func (c imageClient) InsertImage(image model.Image) model.Image {

	result := Db.Create(&image)

	if result.Error != nil {
		log.Error("Failed to insert image.")
		return image
	}

	log.Debug("Image created:", image.Id)
	return image
}

func (c imageClient) InsertImages(images model.Images) model.Images {

	for i := range images {
		result := Db.Create(&images[i])

		if result.Error != nil {
			log.Error("Failed to insert image.")
			return images
		}

		id := images[i].Id
		Db.First(&images[i], id)
	}

	return images
}

func (c imageClient) GetImageById(id int) model.Image {
	var image model.Image

	Db.Where("id = ?", id).First(&image)
	log.Debug("Image: ", image)

	return image
}

func (c imageClient) GetImages() model.Images {
	var images model.Images
	Db.Find(&images)

	log.Debug("Images: ", images)

	return images
}

func (c imageClient) GetImagesByHotelId(hotelId int) model.Images {
	var images model.Images

	Db.Where("hotel_id = ?", hotelId).Find(&images)
	log.Debug("Images: ", images)

	return images
}
