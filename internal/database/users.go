package database

import (
	"log"

	entities "order_orbit/internal/entities"
)

func (s *service) Users() []entities.User {
	var users []entities.User

	result := s.db.Find(&users)

	if result.Error != nil {
		log.Panic(result.Error)
	}
	return users
}

func (s *service) UserById(id string) entities.User {
	var user entities.User
	result := s.db.First(&user, id)

	if result.Error != nil {
		log.Panic(result.Error)
	}
	return user
}

func (s *service) CreateUser(attr map[string]string) entities.User {
	user := entities.User{
		Login:    attr["Login"],
		FullName: attr["FullName"],
	}

	result := s.db.Create(&user)
	if result.Error != nil {
		log.Panic(result.Error)
	}
	s.db.First(&user, user.ID)
	return user
}

func (s *service) DeleteUser(id string) {
	var user entities.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		log.Panic(result.Error)
	}
	s.db.Delete(&user)
}

func (s *service) UpdateUser(id string, attr map[string]string) entities.User {
	var user entities.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		log.Panic(result.Error)
	}

	if newName, ok := attr["FullName"]; ok {
		user.FullName = newName
	}

	s.db.Save(&user)
	return user
}
