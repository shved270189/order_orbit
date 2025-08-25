package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"

	entities "order_orbit/internal/entities"

	_ "github.com/joho/godotenv/autoload"
)

var (
	Conn *gorm.DB
)

func Connect() {
	if Conn != nil {
		return
	}

	var err error
	Conn, err = gorm.Open(sqlite.Open(os.Getenv("BLUEPRINT_DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to database")

	runMigrations()
}

func Close() {
	if Conn != nil {
		sqlDB, err := Conn.DB()
		if err != nil {
			log.Fatal(err)
		}
		sqlDB.Close()
		log.Print("Disconnected from database")
	}
}

func runMigrations() {
	log.Print("Migrating database")
	Conn.AutoMigrate(&entities.User{})
	log.Print("Migration completed")
}
