package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"order_orbit/internal/database"
)

type user struct {
	Index gin.HandlerFunc
	Create gin.HandlerFunc
	Update gin.HandlerFunc
	Delete gin.HandlerFunc
	Show gin.HandlerFunc
}

var(
	Users = user{
		Index: usersHandler,
		Create: createUserHandler,
		Update: updateUserHandler,
		Delete: deleteUserHandler,
		Show: userHandler,
	}
)

func usersHandler(c *gin.Context) {
	result := database.Users()

	c.JSON(http.StatusOK, result)
}

func userHandler(c *gin.Context) {
	result := database.UserById(c.Param("id"))

	c.JSON(http.StatusOK, result)
}

func createUserHandler(c *gin.Context) {
	userAttr := map[string]string{"Login": c.PostForm("login"), "FullName": c.PostForm("full_name")}
	log.Print(userAttr)
	result := database.CreateUser(userAttr)

	c.JSON(http.StatusCreated, result)
}

func deleteUserHandler(c *gin.Context) {
	database.DeleteUser(c.Param("id"))
	c.Status(http.StatusNoContent)
}

func updateUserHandler(c *gin.Context) {
	userAttr := map[string]string{"FullName": c.PostForm("full_name")}
	log.Print(userAttr)
	result := database.UpdateUser(c.Param("id"), userAttr)

	c.JSON(http.StatusOK, result)
}
