package service

import (
	"errors"
	"project/client"
	"project/dto"
	"project/model"
	"time"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto.HotelDto, error)
	GetHotels() (dto.HotelsDto, error)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)
	CheckAvailability(hotelId int, startDate time.Time, endDate time.Time) bool
	CheckAllAvailability(startDate string, endDate string) (dto.HotelsDto, error)
	DeleteHotel(id int) error
	UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)
}

var HotelService hotelServiceInterface

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {
	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.RoomAmount = hotelDto.RoomAmount
	hotel.StreetName = hotelDto.StreetName
	hotel.StreetNumber = hotelDto.StreetNumber
	hotel.Rate = hotelDto.Rate

	for _, amenityName := range hotelDto.Amenities {
		amenity := client.AmenityClient.GetAmenityByName(amenityName)

		if amenity.Id == 0 {
			return hotelDto, errors.New("amenity not found")
		}

		hotel.Amenities = append(hotel.Amenities, amenity)

	}

	hotel = client.HotelClient.InsertHotel(hotel)

	hotelDto.Id = hotel.Id

	if hotel.Id == 0 {
		return hotelDto, errors.New("error creating hotel")
	}

	return hotelDto, nil
}

func (s *hotelService) GetHotels() (dto.HotelsDto, error) {

	var hotels model.Hotels = client.HotelClient.GetHotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto
		hotelDto.Id = hotel.Id
		hotelDto.Name = hotel.Name
		hotelDto.RoomAmount = hotel.RoomAmount
		hotelDto.Description = hotel.Description
		hotelDto.StreetName = hotel.StreetName
		hotelDto.StreetNumber = hotel.StreetNumber
		hotelDto.Rate = hotel.Rate

		if len(hotel.Images) > 0 {
			var imageDto dto.ImageDto
			imageDto.Id = hotel.Images[0].Id
			imageDto.Path = hotel.Images[0].Path
			imageDto.HotelId = hotel.Images[0].HotelId

			hotelDto.Images = append(hotelDto.Images, imageDto)
		}

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *hotelService) GetHotelById(id int) (dto.HotelDto, error) {

	var hotel model.Hotel = client.HotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.Id == 0 {
		return hotelDto, errors.New("hotel not found")
	}
	hotelDto.Id = hotel.Id
	hotelDto.Name = hotel.Name
	hotelDto.RoomAmount = hotel.RoomAmount
	hotelDto.Description = hotel.Description
	hotelDto.StreetName = hotel.StreetName
	hotelDto.StreetNumber = hotel.StreetNumber
	hotelDto.Rate = hotel.Rate

	for _, amenity := range hotel.Amenities {
		hotelDto.Amenities = append(hotelDto.Amenities, amenity.Name)
	}

	for _, image := range hotel.Images {
		var imageDto dto.ImageDto
		imageDto.Id = image.Id
		imageDto.Path = image.Path
		imageDto.HotelId = image.HotelId

		hotelDto.Images = append(hotelDto.Images, imageDto)
	}

	return hotelDto, nil
}

func (s *hotelService) CheckAvailability(hotelId int, startDate time.Time, endDate time.Time) bool {

	hotel := client.HotelClient.GetHotelById(hotelId)
	reservations := client.ReservationClient.GetReservationsByHotel(hotelId)

	roomsAvailable := hotel.RoomAmount

	for _, reservation := range reservations {

		reservationStart, _ := time.Parse("02-01-2006 15:04", reservation.StartDate)
		reservationEnd, _ := time.Parse("02-01-2006 15:04", reservation.EndDate)

		if reservationStart.After(startDate) && reservationEnd.Before(endDate) ||
			reservationStart.Before(startDate) && reservationEnd.After(startDate) ||
			reservationStart.Before(endDate) && reservationEnd.After(endDate) ||
			reservationStart.Before(startDate) && reservationEnd.After(endDate) ||
			reservationStart.Equal(startDate) || reservationEnd.Equal(endDate) {
			roomsAvailable--
		}
		if roomsAvailable == 0 {
			return false
		}
	}

	return true
}

func (s *hotelService) CheckAllAvailability(startDate string, endDate string) (dto.HotelsDto, error) {

	var hotelsAvailable dto.HotelsDto

	reservationStart, _ := time.Parse("02-01-2006 15:04", startDate)
	reservationEnd, _ := time.Parse("02-01-2006 15:04", endDate)

	if reservationStart.After(reservationEnd) {
		return hotelsAvailable, errors.New("a reservation cant end before it starts")
	}

	hotels := client.HotelClient.GetHotels()

	for _, hotel := range hotels {
		if s.CheckAvailability(hotel.Id, reservationStart, reservationEnd) {
			var hotelDto dto.HotelDto
			hotelDto.Id = hotel.Id
			hotelDto.Name = hotel.Name
			hotelDto.StreetName = hotel.StreetName
			hotelDto.StreetNumber = hotel.StreetNumber
			hotelDto.RoomAmount = hotel.RoomAmount
			hotelDto.Rate = hotel.Rate
			hotelDto.Description = hotel.Description

			if len(hotel.Images) > 0 {
				var imageDto dto.ImageDto
				imageDto.Id = hotel.Images[0].Id
				imageDto.Path = hotel.Images[0].Path
				imageDto.HotelId = hotel.Images[0].HotelId

				hotelDto.Images = append(hotelDto.Images, imageDto)
			}

			hotelsAvailable = append(hotelsAvailable, hotelDto)
		}
	}

	return hotelsAvailable, nil
}

func (s *hotelService) DeleteHotel(id int) error {

	hotel := client.HotelClient.GetHotelById(id)

	if hotel.Id == 0 {
		return errors.New("hotel not found")
	}

	for _, image := range hotel.Images {
		err := client.ImageClient.DeleteImage(image)

		if err != nil {
			return err
		}
	}

	err := client.HotelClient.DeleteHotel(hotel)

	return err
}

func (s *hotelService) UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {

	hotel := client.HotelClient.GetHotelById(hotelDto.Id)

	if hotel.Id == 0 {
		return hotelDto, errors.New("hotel not found")
	}

	hotel.Name = hotelDto.Name
	hotel.StreetName = hotelDto.StreetName
	hotel.StreetNumber = hotelDto.StreetNumber
	hotel.Rate = hotelDto.Rate
	hotel.Description = hotelDto.Description
	hotel.RoomAmount = hotelDto.RoomAmount
	hotel.Amenities = model.Amenities{}

	for _, amenityName := range hotelDto.Amenities {
		amenity := client.AmenityClient.GetAmenityByName(amenityName)

		if amenity.Id == 0 {
			return hotelDto, errors.New("amenity not found")
		}

		hotel.Amenities = append(hotel.Amenities, amenity)
	}

	hotel = client.HotelClient.UpdateHotel(hotel)

	if hotel.Id == 0 {
		return hotelDto, errors.New("error updating hotel")
	}

	return hotelDto, nil

}
