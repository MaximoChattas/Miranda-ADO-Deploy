package service

import (
	"errors"
	"math"
	"project/client"
	"project/dto"
	"project/model"
	"time"
)

type reservationService struct{}

type reservationServiceInterface interface {
	InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error)
	GetReservationById(id int) (dto.ReservationDto, error)
	GetReservations() (dto.ReservationsDto, error)
	GetReservationsByUser(userId int) (dto.UserReservationsDto, error)
	GetReservationsByUserRange(userId int, startDate string, endDate string) (dto.ReservationsDto, error)
	GetReservationsByHotel(hotelId int) (dto.HotelReservationsDto, error)
	DeleteReservation(id int) error
}

var ReservationService reservationServiceInterface

func init() {
	ReservationService = &reservationService{}
}

func (s *reservationService) InsertReservation(reservationDto dto.ReservationDto) (dto.ReservationDto, error) {

	userDto := client.UserClient.GetUserById(reservationDto.UserId)
	hotelDto := client.HotelClient.GetHotelById(reservationDto.HotelId)

	if userDto.Id == 0 {
		return reservationDto, errors.New("user not found")
	}

	if hotelDto.Id == 0 {
		return reservationDto, errors.New("hotel not found")
	}

	timeStart, _ := time.Parse("02-01-2006 15:04", reservationDto.StartDate)
	timeEnd, _ := time.Parse("02-01-2006 15:04", reservationDto.EndDate)

	if timeStart.After(timeEnd) {
		return reservationDto, errors.New("a reservation cant end before it starts")
	}

	if HotelService.CheckAvailability(reservationDto.HotelId, timeStart, timeEnd) {
		var reservation model.Reservation

		reservation.StartDate = reservationDto.StartDate
		reservation.EndDate = reservationDto.EndDate
		reservation.HotelId = reservationDto.HotelId
		reservation.UserId = reservationDto.UserId

		hoursAmount := timeEnd.Sub(timeStart).Hours()
		nightsAmount := math.Ceil(hoursAmount / 24)
		rate := hotelDto.Rate

		reservation.Amount = rate * nightsAmount

		reservation = client.ReservationClient.InsertReservation(reservation)

		reservationDto.Id = reservation.Id
		reservationDto.Amount = reservation.Amount

		return reservationDto, nil
	}

	return reservationDto, errors.New("there are no rooms available")
}

func (s *reservationService) GetReservationById(id int) (dto.ReservationDto, error) {
	var reservation model.Reservation
	var reservationDto dto.ReservationDto

	reservation = client.ReservationClient.GetReservationById(id)

	if reservation.Id == 0 {
		return reservationDto, errors.New("reservation not found")
	}

	reservationDto.Id = reservation.Id
	reservationDto.StartDate = reservation.StartDate
	reservationDto.EndDate = reservation.EndDate
	reservationDto.HotelId = reservation.HotelId
	reservationDto.UserId = reservation.UserId
	reservationDto.Amount = reservation.Amount

	return reservationDto, nil
}

func (s *reservationService) GetReservations() (dto.ReservationsDto, error) {

	var reservations model.Reservations = client.ReservationClient.GetReservations()
	var reservationsDto dto.ReservationsDto

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto

		reservationDto.Id = reservation.Id
		reservationDto.StartDate = reservation.StartDate
		reservationDto.EndDate = reservation.EndDate
		reservationDto.HotelId = reservation.HotelId
		reservationDto.UserId = reservation.UserId
		reservationDto.Amount = reservation.Amount

		reservationsDto = append(reservationsDto, reservationDto)
	}

	return reservationsDto, nil
}

func (s *reservationService) GetReservationsByUser(userId int) (dto.UserReservationsDto, error) {
	var user model.User = client.UserClient.GetUserById(userId)
	var userReservationsDto dto.UserReservationsDto
	var reservationsDto dto.ReservationsDto

	if user.Id == 0 {
		return userReservationsDto, errors.New("user not found")
	}
	var reservations model.Reservations = client.ReservationClient.GetReservationsByUser(userId)

	userReservationsDto.UserId = user.Id
	userReservationsDto.UserName = user.Name
	userReservationsDto.UserLastName = user.LastName
	userReservationsDto.UserDni = user.Dni
	userReservationsDto.UserEmail = user.Email
	userReservationsDto.UserPassword = user.Password

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto

		reservationDto.Id = reservation.Id
		reservationDto.StartDate = reservation.StartDate
		reservationDto.EndDate = reservation.EndDate
		reservationDto.HotelId = reservation.HotelId
		reservationDto.UserId = reservation.UserId
		reservationDto.Amount = reservation.Amount

		reservationsDto = append(reservationsDto, reservationDto)
	}

	userReservationsDto.Reservations = reservationsDto

	return userReservationsDto, nil
}

func (s *reservationService) GetReservationsByUserRange(userId int, startDate string, endDate string) (dto.ReservationsDto, error) {

	var reservationsInRange dto.ReservationsDto

	rangeStart, _ := time.Parse("02-01-2006 15:04", startDate)
	rangeEnd, _ := time.Parse("02-01-2006 15:04", endDate)

	if rangeStart.After(rangeEnd) {
		return reservationsInRange, errors.New("a reservation cant end before it starts")
	}

	reservations := client.ReservationClient.GetReservationsByUser(userId)

	for _, reservation := range reservations {

		reservationStart, _ := time.Parse("02-01-2006 15:04", reservation.StartDate)
		reservationEnd, _ := time.Parse("02-01-2006 15:04", reservation.EndDate)

		if reservationStart.After(rangeStart) && reservationEnd.Before(rangeEnd) ||
			reservationStart.Before(rangeStart) && reservationEnd.After(rangeEnd) ||
			reservationStart.Before(rangeEnd) && reservationEnd.After(rangeEnd) ||
			reservationStart.Before(rangeStart) && reservationEnd.After(rangeEnd) ||
			reservationStart.Equal(rangeStart) || reservationEnd.Equal(rangeEnd) {

			var reservationDto dto.ReservationDto

			reservationDto.Id = reservation.Id
			reservationDto.StartDate = reservation.StartDate
			reservationDto.EndDate = reservation.EndDate
			reservationDto.Amount = reservation.Amount
			reservationDto.UserId = reservation.UserId
			reservationDto.HotelId = reservation.HotelId

			reservationsInRange = append(reservationsInRange, reservationDto)
		}
	}

	return reservationsInRange, nil
}

func (s *reservationService) GetReservationsByHotel(hotelId int) (dto.HotelReservationsDto, error) {
	var hotel model.Hotel = client.HotelClient.GetHotelById(hotelId)
	var hotelReservations dto.HotelReservationsDto
	var reservationsDto dto.ReservationsDto

	if hotel.Id == 0 {
		return hotelReservations, errors.New("hotel not found")
	}

	var reservations model.Reservations = client.ReservationClient.GetReservationsByHotel(hotelId)

	hotelReservations.HotelId = hotel.Id
	hotelReservations.HotelName = hotel.Name
	hotelReservations.HotelDescription = hotel.Description
	hotelReservations.HotelRoomAmount = hotel.RoomAmount
	hotelReservations.HotelStreetName = hotel.StreetName
	hotelReservations.HotelStreetNumber = hotel.StreetNumber
	hotelReservations.HotelRate = hotel.Rate

	for _, reservation := range reservations {
		var reservationDto dto.ReservationDto
		reservationDto.Id = reservation.Id
		reservationDto.StartDate = reservation.StartDate
		reservationDto.EndDate = reservation.EndDate
		reservationDto.HotelId = reservation.HotelId
		reservationDto.UserId = reservation.UserId
		reservationDto.Amount = reservation.Amount

		reservationsDto = append(reservationsDto, reservationDto)
	}

	hotelReservations.Reservations = reservationsDto

	return hotelReservations, nil
}

func (s *reservationService) DeleteReservation(id int) error {

	reservation := client.ReservationClient.GetReservationById(id)

	if reservation.Id == 0 {
		return errors.New("reservation not found")
	}

	reservationStart, _ := time.Parse("02-01-2006 15:04", reservation.StartDate)

	if reservationStart.Before(time.Now().Add(-48 * time.Hour)) {
		return errors.New("can't delete a reservation 48hs before it starts")
	}

	err := client.ReservationClient.DeleteReservation(reservation)

	return err

}
