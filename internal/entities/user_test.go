package entities

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserCreation(t *testing.T) {
	now := time.Now()
	user := User{
		ID:        1,
		CreatedAt: now,
		UpdatedAt: now,
		Login:     "testuser",
		FullName:  "Test User",
	}

	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, now, user.CreatedAt)
	assert.Equal(t, now, user.UpdatedAt)
	assert.Equal(t, "testuser", user.Login)
	assert.Equal(t, "Test User", user.FullName)
}

func TestUserJSONMarshaling(t *testing.T) {
	now := time.Now()
	user := User{
		ID:        1,
		CreatedAt: now,
		UpdatedAt: now,
		Login:     "testuser",
		FullName:  "Test User",
	}

	jsonData, err := json.Marshal(user)
	assert.NoError(t, err)

	var unmarshaled User
	err = json.Unmarshal(jsonData, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, user.ID, unmarshaled.ID)
	assert.Equal(t, user.CreatedAt.Unix(), unmarshaled.CreatedAt.Unix())
	assert.Equal(t, user.UpdatedAt.Unix(), unmarshaled.UpdatedAt.Unix())
	assert.Equal(t, user.Login, unmarshaled.Login)
	assert.Equal(t, user.FullName, unmarshaled.FullName)
}

func TestUserFields(t *testing.T) {
	user := User{}

	// Test default values
	assert.Equal(t, uint(0), user.ID)
	assert.Empty(t, user.Login)
	assert.Empty(t, user.FullName)

	// Test field setting
	user.ID = 42
	user.Login = "johndoe"
	user.FullName = "John Doe"

	assert.Equal(t, uint(42), user.ID)
	assert.Equal(t, "johndoe", user.Login)
	assert.Equal(t, "John Doe", user.FullName)
}