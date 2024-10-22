package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"meeting-room-booking/controller"
	"meeting-room-booking/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (m *MockBookingUseCase) GetAll() ([]domain.Booking, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Booking), args.Error(1)
}

func (m *MockBookingUseCase) GetBookingByID(id int) (*domain.Booking, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Booking), args.Error(1)
}

func (m *MockBookingUseCase) CreateBooking(booking domain.Booking) error {
	args := m.Called(booking)
	return args.Error(0)
}

func (m *MockBookingUseCase) UpdateBooking(booking domain.Booking) error {
	args := m.Called(booking)
	return args.Error(0)
}

func (m *MockBookingUseCase) DeleteBooking(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Mock UseCase
type MockBookingUseCase struct {
	mock.Mock
}

// 1. Test: Get All Bookings
func TestGetAllBookings(t *testing.T) {
	// Setup Gin and Mock
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockUseCase := new(MockBookingUseCase)
	bookingCtrl := controller.NewBookingController(mockUseCase)

	// Mock response
	mockBookings := []domain.Booking{
		{ID: 1, RoomName: "Room A", StartTime: "09:00", EndTime: "10:00", ReservedBy: "John"},
		{ID: 2, RoomName: "Room B", StartTime: "11:00", EndTime: "12:00", ReservedBy: "Jane"},
	}
	mockUseCase.On("GetAll").Return(mockBookings, nil)

	r.GET("/v1/bookings", bookingCtrl.GetAll)
	req, _ := http.NewRequest(http.MethodGet, "/v1/bookings", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// // ถ้าอยากจะเช็คว่า message ที่ได้มาจาก error นั้นตรงกับที่คาดหวัง
	// assert.Equal(t, w.Body.String(), gin.H{"message": "Booking created successfully"})

	mockUseCase.AssertExpectations(t)
}

// 2. Test: Get Booking by ID
func TestGetBookingByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockUseCase := new(MockBookingUseCase)
	bookingCtrl := controller.NewBookingController(mockUseCase)

	mockBooking := &domain.Booking{ID: 1, RoomName: "Room A", StartTime: "09:00", EndTime: "10:00", ReservedBy: "John"}

	// ใส่เป็น function ของ mockUseCase ที่จะถูกเรียก
	mockUseCase.On("GetBookingByID", 1).Return(mockBooking, nil)

	r.GET("/bookings/:id", bookingCtrl.GetByID)
	req, _ := http.NewRequest(http.MethodGet, "/bookings/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertExpectations(t)
}

// 3. Test: Create Booking (POST)
func TestCreateBooking(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockUseCase := new(MockBookingUseCase)
	bookingCtrl := controller.NewBookingController(mockUseCase)

	newBooking := domain.Booking{RoomName: "Room C", StartTime: "13:00", EndTime: "14:00", ReservedBy: "Bob"}

	mockUseCase.On("CreateBooking", newBooking).Return(nil)

	r.POST("/bookings", bookingCtrl.Create)

	jsonValue, _ := json.Marshal(newBooking)
	req, _ := http.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// // print body
	// resp := w.Result()
	// body, _ := io.ReadAll(resp.Body)
	// log.Printf("Response: %s", body)

	assert.Equal(t, http.StatusCreated, w.Code)
	// assert.Equal(t, w.Body.String(), gin.H{"message": "Booking created successfully"})
	mockUseCase.AssertExpectations(t)
}

func TestUpdateBooking(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.New()

	mockUseCase := new(MockBookingUseCase)
	bookingCtrl := controller.NewBookingController(mockUseCase)

	updateBooking := domain.Booking{ID: 1, RoomName: "Room C", StartTime: "13:00", EndTime: "14:00", ReservedBy: "Bob"}

	mockUseCase.On("UpdateBooking", updateBooking).Return(nil)

	r.PUT("/bookings/:id", bookingCtrl.Update)

	jsonValue, _ := json.Marshal(updateBooking)
	req, _ := http.NewRequest(http.MethodPut, "/bookings/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertExpectations(t)
}

// 5. Test: Delete Booking (DELETE)
func TestDeleteBooking(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockUseCase := new(MockBookingUseCase)
	bookingCtrl := controller.NewBookingController(mockUseCase)

	mockUseCase.On("DeleteBooking", 1).Return(nil)

	r.DELETE("/bookings/:id", bookingCtrl.Delete)
	req, _ := http.NewRequest(http.MethodDelete, "/bookings/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertExpectations(t)
}

// 6. Test: Get All Bookings Error
func TestGetAllBookings_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockUseCase := new(MockBookingUseCase)
	bookingCtrl := controller.NewBookingController(mockUseCase)

	// Mock ให้ UseCase คืนค่า error
	mockUseCase.On("GetAll").Return(nil, errors.New("something went wrong"))

	r.GET("/bookings", bookingCtrl.GetAll)
	req, _ := http.NewRequest(http.MethodGet, "/bookings", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// ตรวจสอบว่า status code ที่ได้คือ 500 Internal Server Error
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// // ตรวจสอบว่า response มีข้อความ error ที่ถูกต้อง
	// assert.Contains(t, w.Body.String(), "something went wrong")

	// ตรวจสอบว่า response มีข้อความ error ที่ถูกต้องใน JSON response
	var responseBody map[string]string
	json.Unmarshal(w.Body.Bytes(), &responseBody)

	assert.Equal(t, "something went wrong", responseBody["message"])

	// ตรวจสอบว่า mock ถูกเรียกตามที่กำหนด
	mockUseCase.AssertExpectations(t)
}
