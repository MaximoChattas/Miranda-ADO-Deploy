package service

import (
	"github.com/stretchr/testify/assert"
	"project/client"
	"project/dto"
	"project/model"
	"testing"
)

type TestImage struct{}

func init() {
	client.ImageClient = &TestImage{}
}

func (t TestImage) InsertImage(image model.Image) model.Image {

	if image.Path == "" {
		image.Id = 0
	} else {
		image.Id = 1
	}

	return image
}

func (t TestImage) InsertImages(images model.Images) model.Images {

	if len(images) == 0 {
		images = append(images, model.Image{
			Id:      0,
			Path:    "",
			HotelId: 0,
		})
	} else {
		for i := range images {
			images[i].Id = i + 1
		}
	}

	return images
}

func (t TestImage) GetImageById(id int) model.Image {
	var image model.Image

	if id > 10 {
		image.Id = 0
	} else {
		image.Id = id
	}

	return image
}

func (t TestImage) GetImages() model.Images {
	return model.Images{}
}

func (t TestImage) GetImagesByHotelId(hotelId int) model.Images {
	return model.Images{}
}

func (t TestImage) DeleteImage(image model.Image) error { return nil }

func TestInsertImages_Service_Error(t *testing.T) {

	a := assert.New(t)
	var images dto.ImagesDto

	_, err := ImageService.InsertImages(images)

	expectedResponse := "failed to insert images"

	a.NotNil(err)
	a.Equal(expectedResponse, err.Error())
}

func TestInsertImages_Service_Success(t *testing.T) {

	a := assert.New(t)
	images := dto.ImagesDto{
		dto.ImageDto{Path: "image1.jpg"},
	}

	result, err := ImageService.InsertImages(images)

	images[0].Id = 1

	a.Nil(err)
	a.Equal(images, result)
}

func TestGetImageById_Service_Found(t *testing.T) {

	a := assert.New(t)

	result, err := ImageService.GetImageById(1)

	expectedResult := dto.ImageDto{Id: 1}

	a.Nil(err)
	a.Equal(expectedResult, result)
}

func TestGetImageById_Service_NotFound(t *testing.T) {

	a := assert.New(t)

	_, err := ImageService.GetImageById(12)

	expectedResult := "image not found"

	a.NotNil(err)
	a.Equal(expectedResult, err.Error())
}
