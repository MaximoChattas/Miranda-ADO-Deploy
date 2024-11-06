package service

import (
	"errors"
	"project/client"
	"project/dto"
	"project/model"
)

type amenityService struct{}

type amenityServiceInterface interface {
	InsertAmenity(amenityDto dto.AmenityDto) (dto.AmenityDto, error)
	GetAmenities() (dto.AmenitiesDto, error)
}

var AmenityService amenityServiceInterface

func init() {
	AmenityService = &amenityService{}
}

func (s *amenityService) InsertAmenity(amenityDto dto.AmenityDto) (dto.AmenityDto, error) {
	var amenity model.Amenity

	amenity.Name = amenityDto.Name

	amenity = client.AmenityClient.InsertAmenity(amenity)

	if amenity.Id == 0 {
		return amenityDto, errors.New("error creating amenity")
	}

	amenityDto.Id = amenity.Id

	return amenityDto, nil
}

func (s *amenityService) GetAmenities() (dto.AmenitiesDto, error) {
	var amenities model.Amenities = client.AmenityClient.GetAmenities()
	var amenitiesDto dto.AmenitiesDto

	for _, amenity := range amenities {
		var amenityDto dto.AmenityDto
		amenityDto.Id = amenity.Id
		amenityDto.Name = amenity.Name

		amenitiesDto = append(amenitiesDto, amenityDto)
	}

	return amenitiesDto, nil
}
