package storage

import (
	"testing"
	"strconv"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"order_orbit/internal/entities"
	"order_orbit/internal/testutil"
)

func createTestUsers(t *testing.T, db *gorm.DB) []entities.User {
	users := []entities.User{
		{Login: "user1", FullName: "User One"},
		{Login: "user2", FullName: "User Two"},
		{Login: "user3", FullName: "User Three"},
	}
	for i := range users {
		result := db.Create(&users[i])
		assert.NoError(t, result.Error)
	}

	return users
}

func TestNewUserStorage(t *testing.T) {
	userStorage := NewUserStorage(database)
	assert.NotNil(t, userStorage)
	assert.Equal(t, database, userStorage.db)
}

func TestUserStorage_All(t *testing.T) {
	testutil.WithTx(database, func(db *gorm.DB) {
		testUsers := createTestUsers(t, db)
		userStorage := NewUserStorage(db)
		users := userStorage.All()
		assert.Equal(t, len(testUsers), len(users))
		for i, user := range users {
			assert.Equal(t, testUsers[i].Login, user.Login)
			assert.Equal(t, testUsers[i].FullName, user.FullName)
		}
	})
}

func TestUserStorage_Find(t *testing.T) {
	testutil.WithTx(database, func(db *gorm.DB) {
		testUsers := createTestUsers(t, db)
		userStorage := NewUserStorage(db)
		user := userStorage.Find(strconv.FormatUint(uint64(testUsers[1].ID), 10))
		assert.Equal(t, testUsers[1].ID, user.ID)
		assert.Equal(t, testUsers[1].Login, user.Login)
		assert.Equal(t, testUsers[1].FullName, user.FullName)
	})
}

// func TestUserStorage_Create(t *testing.T) {
// 	db := setupTestDB(t)
// 	userStorage := NewUserStorage(db)

// 	attrs := map[string]any{
// 		"Login":    "newuser",
// 		"FullName": "New User",
// 	}

// 	user := userStorage.Create(attrs)

// 	assert.NotEmpty(t, user.ID)
// 	assert.Equal(t, "newuser", user.Login)
// 	assert.Equal(t, "New User", user.FullName)

// 	// Verify in the database
// 	var dbUser entities.User
// 	result := db.First(&dbUser, user.ID)
// 	assert.NoError(t, result.Error)
// 	assert.Equal(t, user.ID, dbUser.ID)
// }

// func TestUserStorage_Delete(t *testing.T) {
// 	db := setupTestDB(t)
// 	testUsers := createTestUsers(t, db)
// 	userStorage := NewUserStorage(db)

// 	userStorage.Delete(testUsers[0].ID)

// 	// Verify deletion
// 	var user entities.User
// 	result := db.First(&user, testUsers[0].ID)
// 	assert.Error(t, result.Error)
// 	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
// }

// func TestUserStorage_Update(t *testing.T) {
// 	db := setupTestDB(t)
// 	testUsers := createTestUsers(t, db)
// 	userStorage := NewUserStorage(db)

// 	attrs := map[string]any{
// 		"FullName": "Updated Name",
// 	}

// 	user := userStorage.Update(testUsers[0].ID, attrs)

// 	assert.Equal(t, testUsers[0].ID, user.ID)
// 	assert.Equal(t, testUsers[0].Login, user.Login)
// 	assert.Equal(t, "Updated Name", user.FullName)

// 	// Verify in database
// 	var dbUser entities.User
// 	result := db.First(&dbUser, user.ID)
// 	assert.NoError(t, result.Error)
// 	assert.Equal(t, "Updated Name", dbUser.FullName)
// }
