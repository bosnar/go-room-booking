package main

import (
	"database/sql"
	"fmt"
	"log"
	"meeting-room-booking/config"
)

func main() {

	db := config.PgConnect()
	log.Println("Starting database migration...")

	// ตัวอย่างการสร้างตารางใหม่
	if err := migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully!")
}

func migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS bookings (
		id SERIAL PRIMARY KEY,
		room_name VARCHAR(255) NOT NULL,
		start_time VARCHAR(255) NOT NULL,
		end_time VARCHAR(255) NOT NULL,
		reserved_by VARCHAR(255) NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Rooms table created.")
	return nil
}
