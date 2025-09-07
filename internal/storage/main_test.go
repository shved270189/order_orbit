package storage

import (
	"os"
	"testing"

	"gorm.io/gorm"

	"order_orbit/internal/testutil"
)

var (
	database *gorm.DB
)

func TestMain(m *testing.M) {
	var code int
	testutil.WithDatabaseConnection(func(db *gorm.DB) {
		database = db
		code = m.Run()
	})

	os.Exit(code)
}
