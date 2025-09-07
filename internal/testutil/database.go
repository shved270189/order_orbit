package testutil

import (
	"log"

	"gorm.io/gorm"

	"order_orbit/internal/config"
	"order_orbit/internal/database"
)

func WithDatabaseConnection(fn func(db *gorm.DB)) {
	config := config.Load("test")
	dbConnection := database.New(config.ProjectRoot + config.DatabaseURL)
	defer dbConnection.Close()
	dbConnection.RunMigrations()

	fn(dbConnection.DB())
}

func WithTx(db *gorm.DB, fn func(tx *gorm.DB)) {
    tx := db.Begin()
    if tx.Error != nil {
        log.Fatalf("failed to begin tx: %v", tx.Error)
    }
    defer tx.Rollback() // cleanup after test

    fn(tx) // run the test logic
}
