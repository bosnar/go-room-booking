package router

import (
	"database/sql"
	"meeting-room-booking/controller"
	"meeting-room-booking/repository"
	"meeting-room-booking/usecase"

	"github.com/gin-gonic/gin"
)

func BookingRouter(router *gin.RouterGroup, db *sql.DB) {
	repo := repository.NewBookingRepository(db)
	usecase := usecase.NewBookingUseCase(repo)
	controller := controller.NewBookingController(usecase)

	{
		router.GET("/bookings", controller.GetAll)
		router.GET("/booking/:id", controller.GetByID)
		router.POST("/booking", controller.Create)
		router.PUT("/booking/:id", controller.Update)
		router.DELETE("/booking/:id", controller.Delete)
	}
}
