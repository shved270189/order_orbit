package storage

import (
	"log"

	"order_orbit/internal/database"
	"order_orbit/internal/entities"
)

type usersStorage struct {
}

var (
	Users = usersStorage{}
)

func (s *usersStorage) All() []entities.User {
	var users []entities.User

	result := database.Conn.Find(&users)

	if result.Error != nil {
		log.Panic(result.Error)
	}
	return users
}

func (s *usersStorage) Find(id string) entities.User {
	var user entities.User
	result := database.Conn.First(&user, id)

	if result.Error != nil {
		log.Panic(result.Error)
	}
	return user
}

func (s *usersStorage) Create(attr map[string]string) entities.User {
	user := entities.User{
		Login:    attr["Login"],
		FullName: attr["FullName"],
	}

	result := database.Conn.Create(&user)
	if result.Error != nil {
		log.Panic(result.Error)
	}
	database.Conn.First(&user, user.ID)
	return user
}

func (s *usersStorage) Delete(id string) {
	var user entities.User
	result := database.Conn.First(&user, id)
	if result.Error != nil {
		log.Panic(result.Error)
	}
	database.Conn.Delete(&user)
}

func (s *usersStorage) Update(id string, attr map[string]string) entities.User {
	var user entities.User
	result := database.Conn.First(&user, id)
	if result.Error != nil {
		log.Panic(result.Error)
	}

	if newName, ok := attr["FullName"]; ok {
		user.FullName = newName
	}

	database.Conn.Save(&user)
	return user
}
