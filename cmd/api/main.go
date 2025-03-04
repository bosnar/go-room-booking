package main

import (
	"fmt"
	"meeting-room-booking/config"
	"meeting-room-booking/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Start the server
	fmt.Println("Starting api server...")

	db := config.PgConnect()

	if db == nil {
		panic("Cannot connect to the database")
	}

	// init router http
	r := gin.New()

	router.Router(r, db)
	r.Run(":4000")
}
