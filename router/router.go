package router

import (
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, db *sql.DB) {

	V1 := r.Group("v1")
	{
		V1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	BookingRouter(V1, db)
}
