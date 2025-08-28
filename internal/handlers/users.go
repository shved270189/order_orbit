package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"order_orbit/internal/storage"
)

type UserHandler struct {
	storage *storage.User
}

func NewUserHandler(storage *storage.User) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func (h *UserHandler) Index(c *gin.Context) {
	result := h.storage.All()

	c.JSON(http.StatusOK, result)
}

func (h *UserHandler) Show(c *gin.Context) {
	result := h.storage.Find(c.Param("id"))

	c.JSON(http.StatusOK, result)
}

func (h *UserHandler) Create(c *gin.Context) {
	var params struct {
		Login    string `json:"login" binding:"required"`
		FullName string `json:"fullName" binding:"required"`
	}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := h.storage.Create(
		map[string]any{
			"Login":    params.Login,
			"FullName": params.FullName,
		},
	)

	c.JSON(http.StatusCreated, result)
}

func (h *UserHandler) Destroy(c *gin.Context) {
	h.storage.Delete(c.Param("id"))
	c.Status(http.StatusNoContent)
}

func (h *UserHandler) Update(c *gin.Context) {
	var params struct {
		FullName string `json:"fullName" binding:"required"`
	}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(params)

	result := h.storage.Update(c.Param("id"), map[string]any{"FullName": params.FullName})

	c.JSON(http.StatusOK, result)
}
