package repository

import (
	"meeting-room-booking/domain"
	"meeting-room-booking/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// สร้าง mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// จำลองข้อมูลที่คืนจากฐานข้อมูล
	rows := sqlmock.NewRows([]string{"id", "room_name", "start_time", "end_time", "reserved_by"}).
		AddRow(1, "Room A", "09:00", "10:00", "John").
		AddRow(2, "Room B", "11:00", "12:00", "Jane")

	// Mock query
	mock.ExpectQuery("SELECT id, room_name, start_time, end_time, reserved_by FROM bookings").
		WillReturnRows(rows)

	repo := repository.NewBookingRepository(db)
	bookings, err := repo.GetAll()

	// ตรวจสอบว่าไม่มี error และได้ผลลัพธ์ตามที่คาดไว้
	assert.NoError(t, err)
	assert.Len(t, bookings, 2)
	assert.Equal(t, "Room A", bookings[0].RoomName)
	assert.Equal(t, "Room B", bookings[1].RoomName)

	// ตรวจสอบว่า mock ได้ถูกเรียกใช้ตามที่คาดหวัง
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateBooking(t *testing.T) {
	// สร้าง mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock expected result for insert
	mock.ExpectExec("INSERT INTO bookings").
		WithArgs("Room C", "13:00", "14:00", "Bob").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewBookingRepository(db)
	newBooking := domain.Booking{
		RoomName:   "Room C",
		StartTime:  "13:00",
		EndTime:    "14:00",
		ReservedBy: "Bob",
	}

	err = repo.Create(newBooking)

	// ตรวจสอบว่าไม่มี error
	assert.NoError(t, err)

	// ตรวจสอบว่า mock ได้ถูกเรียกใช้ตามที่คาดหวัง
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteBooking(t *testing.T) {
	// สร้าง mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock expected result for delete
	mock.ExpectExec("DELETE FROM bookings WHERE id = \\$1").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewBookingRepository(db)
	err = repo.Delete(1)

	// ตรวจสอบว่าไม่มี error
	assert.NoError(t, err)

	// ตรวจสอบว่า mock ได้ถูกเรียกใช้ตามที่คาดหวัง
	assert.NoError(t, mock.ExpectationsWereMet())
}
