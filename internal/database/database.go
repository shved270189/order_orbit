package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"

	entities "order_orbit/internal/entities"

	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	AllUsers() []entities.User
	GetUserById(string) entities.User
	CreateUser(map[string]string) entities.User
	DeleteUser(string)
	UpdateUser(string, map[string]string) entities.User
}

type service struct {
	db *gorm.DB
}

var (
	dburl      = os.Getenv("BLUEPRINT_DB_URL")
	dbInstance *service
)

func New() *service {
	if dbInstance != nil {
		return dbInstance
	}

	db, err := gorm.Open(sqlite.Open(dburl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entities.User{})

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}
