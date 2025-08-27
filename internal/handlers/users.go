package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"order_orbit/internal/storage"
)

type usersHandler struct {
}

type userCreateParams struct {
	Login    string `form:"login" binding:"required"`
	FullName string `form:"fullName" binding:"required"`
}

var (
	Users = usersHandler{}
)

func (h *usersHandler) Index(c *gin.Context) {
	result := storage.Users.All()

	c.JSON(http.StatusOK, result)
}

func (h *usersHandler) Show(c *gin.Context) {
	result := storage.Users.Find(c.Param("id"))

	c.JSON(http.StatusOK, result)
}

func (h *usersHandler) Create(c *gin.Context) {
	var params userCreateParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(params)
	result := storage.Users.Create(map[string]string{"Login": params.Login, "FullName": params.FullName})

	c.JSON(http.StatusCreated, result)
}

func (h *usersHandler) Destroy(c *gin.Context) {
	storage.Users.Delete(c.Param("id"))
	c.Status(http.StatusNoContent)
}

func (h *usersHandler) Update(c *gin.Context) {
	userAttr := map[string]string{"FullName": c.PostForm("full_name")}
	log.Print(userAttr)
	result := storage.Users.Update(c.Param("id"), userAttr)

	c.JSON(http.StatusOK, result)
}
