package database

import (
	"log"

	entities "order_orbit/internal/entities"
)

func Users() []entities.User {
	var users []entities.User

	result := Conn.Find(&users)

	if result.Error != nil {
		log.Panic(result.Error)
	}
	return users
}

func UserById(id string) entities.User {
	var user entities.User
	result := Conn.First(&user, id)

	if result.Error != nil {
		log.Panic(result.Error)
	}
	return user
}

func CreateUser(attr map[string]string) entities.User {
	user := entities.User{
		Login:    attr["Login"],
		FullName: attr["FullName"],
	}

	result := Conn.Create(&user)
	if result.Error != nil {
		log.Panic(result.Error)
	}
	Conn.First(&user, user.ID)
	return user
}

func DeleteUser(id string) {
	var user entities.User
	result := Conn.First(&user, id)
	if result.Error != nil {
		log.Panic(result.Error)
	}
	Conn.Delete(&user)
}

func UpdateUser(id string, attr map[string]string) entities.User {
	var user entities.User
	result := Conn.First(&user, id)
	if result.Error != nil {
		log.Panic(result.Error)
	}

	if newName, ok := attr["FullName"]; ok {
		user.FullName = newName
	}

	Conn.Save(&user)
	return user
}
