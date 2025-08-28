package storage

import (
	"log"

	"gorm.io/gorm"

	"order_orbit/internal/entities"
)

type User struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *User {
	return &User{db: db}
}

func (s *User) All() []entities.User {
	var users []entities.User

	result := s.db.Find(&users)

	if result.Error != nil {
		log.Panic(result.Error)
	}
	return users
}

func (s *User) Find(id string) entities.User {
	var user entities.User
	result := s.db.First(&user, id)

	if result.Error != nil {
		log.Panic(result.Error)
	}
	return user
}

func (s *User) Create(attrs map[string]any) entities.User {
	user := entities.User{
		Login:    attrs["Login"].(string),
		FullName: attrs["FullName"].(string),
	}
	result := s.db.Create(&user)
	if result.Error != nil {
		log.Panic(result.Error)
	}
	s.db.First(&user, user.ID)
	return user
}

func (s *User) Delete(id string) {
	var user entities.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		log.Panic(result.Error)
	}
	s.db.Delete(&user)
}

func (s *User) Update(id string, attrs map[string]any) entities.User {
	var user entities.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		log.Panic(result.Error)
	}

	user.FullName = attrs["FullName"].(string)

	s.db.Save(&user)
	return user
}
