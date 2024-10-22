package usecase

import (
	"meeting-room-booking/domain"
	"meeting-room-booking/repository"
)

type BookingUseCase interface {
	GetAll() ([]domain.Booking, error)
	GetBookingByID(id int) (*domain.Booking, error)
	CreateBooking(booking domain.Booking) error
	UpdateBooking(booking domain.Booking) error
	DeleteBooking(id int) error
}

type bookingUseCase struct {
	bookingRepo repository.BookingRepository
}

func NewBookingUseCase(bookingRepo repository.BookingRepository) BookingUseCase {
	return &bookingUseCase{bookingRepo: bookingRepo}
}

func (u *bookingUseCase) GetAll() ([]domain.Booking, error) {
	return u.bookingRepo.GetAll()
}

func (u *bookingUseCase) GetBookingByID(id int) (*domain.Booking, error) {
	return u.bookingRepo.GetByID(id)
}

func (u *bookingUseCase) CreateBooking(booking domain.Booking) error {
	return u.bookingRepo.Create(booking)
}

func (u *bookingUseCase) UpdateBooking(booking domain.Booking) error {
	return u.bookingRepo.Update(booking)
}

func (u *bookingUseCase) DeleteBooking(id int) error {
	return u.bookingRepo.Delete(id)
}
