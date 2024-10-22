package controller

import (
	"meeting-room-booking/domain"
	"meeting-room-booking/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingUseCase usecase.BookingUseCase
}

func NewBookingController(bookingUseCase usecase.BookingUseCase) *BookingController {
	return &BookingController{bookingUseCase: bookingUseCase}
}

func (bc *BookingController) GetAll(c *gin.Context) {
	_, err := bc.bookingUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": nil,
	})
}

func (ctrl *BookingController) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	booking, err := ctrl.bookingUseCase.GetBookingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func (ctrl *BookingController) Create(c *gin.Context) {
	var booking domain.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.bookingUseCase.CreateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Booking created successfully"})
}

func (ctrl *BookingController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var booking domain.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booking.ID = id
	if err := ctrl.bookingUseCase.UpdateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking updated successfully"})
}

func (ctrl *BookingController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.bookingUseCase.DeleteBooking(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
