package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"order_orbit/internal/entities"
)

type Connection struct {
	db *gorm.DB
}

func New(dbUrl string) *Connection {
	connection, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to database")
	return &Connection{db: connection}
}

func (c *Connection) DB() *gorm.DB {
	return c.db
}

func (c *Connection) Close() {
	if c.db == nil {
		return
	}
	sqlDB, err := c.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
	log.Print("Disconnected from database")
}

func (c *Connection) RunMigrations() {
	log.Print("Migrating database")
	c.db.AutoMigrate(&entities.User{})
	log.Print("Migration completed")
}
