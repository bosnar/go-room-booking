package repository

import (
	"database/sql"
	"meeting-room-booking/domain"
)

type BookingRepository interface {
	GetAll() ([]domain.Booking, error)
	GetByID(id int) (*domain.Booking, error)
	Create(booking domain.Booking) error
	Update(booking domain.Booking) error
	Delete(id int) error
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) GetAll() ([]domain.Booking, error) {

	rows, err := r.db.Query("SELECT id, room_name, start_time, end_time, reserved_by FROM bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []domain.Booking
	for rows.Next() {
		var booking domain.Booking
		if err := rows.Scan(&booking.ID, &booking.RoomName, &booking.StartTime, &booking.EndTime, &booking.ReservedBy); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (r *bookingRepository) GetByID(id int) (*domain.Booking, error) {

	row := r.db.QueryRow("SELECT id, room_name, start_time, end_time, reserved_by FROM bookings WHERE id = $1", id)

	var booking domain.Booking
	if err := row.Scan(&booking.ID, &booking.RoomName, &booking.StartTime, &booking.EndTime, &booking.ReservedBy); err != nil {
		return nil, err
	}

	return &booking, nil
}

func (r *bookingRepository) Create(booking domain.Booking) error {

	_, err := r.db.Exec("INSERT INTO bookings (room_name, start_time, end_time, reserved_by) VALUES ($1, $2, $3, $4)", booking.RoomName, booking.StartTime, booking.EndTime, booking.ReservedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *bookingRepository) Update(booking domain.Booking) error {

	_, err := r.db.Exec("UPDATE bookings SET room_name = $1, start_time = $2, end_time = $3, reserved_by = $4 WHERE id = $5", booking.RoomName, booking.StartTime, booking.EndTime, booking.ReservedBy, booking.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *bookingRepository) Delete(id int) error {

	_, err := r.db.Exec("DELETE FROM bookings WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
