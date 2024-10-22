package usecase

import (
	"meeting-room-booking/domain"
	"meeting-room-booking/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Repository
type MockBookingRepository struct {
	mock.Mock
}

func (m *MockBookingRepository) GetAll() ([]domain.Booking, error) {
	args := m.Called()
	return args.Get(0).([]domain.Booking), args.Error(1)
}

func (m *MockBookingRepository) GetByID(id int) (*domain.Booking, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Booking), args.Error(1)
}

func (m *MockBookingRepository) Create(booking domain.Booking) error {
	args := m.Called(booking)
	return args.Error(0)
}

func (m *MockBookingRepository) Update(booking domain.Booking) error {
	args := m.Called(booking)
	return args.Error(0)
}

func (m *MockBookingRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllBookings(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	bookingUseCase := usecase.NewBookingUseCase(mockRepo)

	mockBookings := []domain.Booking{
		{ID: 1, RoomName: "Room A", StartTime: "09:00", EndTime: "10:00", ReservedBy: "John"},
		{ID: 2, RoomName: "Room B", StartTime: "11:00", EndTime: "12:00", ReservedBy: "Jane"},
	}

	// Mock repository response
	mockRepo.On("GetAll").Return(mockBookings, nil)

	result, err := bookingUseCase.GetAll()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "Room A", result[0].RoomName)
	mockRepo.AssertExpectations(t)
}

func TestGetBookingByID(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	bookingUseCase := usecase.NewBookingUseCase(mockRepo)

	mockBooking := &domain.Booking{ID: 1, RoomName: "Room A", StartTime: "09:00", EndTime: "10:00", ReservedBy: "John"}

	// Mock repository response
	mockRepo.On("GetByID", 1).Return(mockBooking, nil)

	result, err := bookingUseCase.GetBookingByID(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	mockRepo.AssertExpectations(t)
}

func TestCreateBooking(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	bookingUseCase := usecase.NewBookingUseCase(mockRepo)

	newBooking := domain.Booking{RoomName: "Room C", StartTime: "13:00", EndTime: "14:00", ReservedBy: "Bob"}

	// Mock repository response
	mockRepo.On("Create", newBooking).Return(nil)

	err := bookingUseCase.CreateBooking(newBooking)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "Room C", newBooking.RoomName)

	mockRepo.AssertExpectations(t)
}

func TestUpdateBooking(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	bookingUseCase := usecase.NewBookingUseCase(mockRepo)

	updatedBooking := domain.Booking{ID: 1, RoomName: "Room A", StartTime: "09:00", EndTime: "10:00", ReservedBy: "John"}

	// Mock repository response
	mockRepo.On("Update", updatedBooking).Return(nil)

	err := bookingUseCase.UpdateBooking(updatedBooking)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBooking(t *testing.T) {
	mockRepo := new(MockBookingRepository)
	bookingUseCase := usecase.NewBookingUseCase(mockRepo)

	// Mock repository response
	mockRepo.On("Delete", 1).Return(nil)

	err := bookingUseCase.DeleteBooking(1)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
