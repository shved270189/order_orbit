package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"order_orbit/internal/database"
	"order_orbit/internal/handlers"
	"order_orbit/internal/storage"
)

func registerRoutes(dbConnection *database.Connection) http.Handler {
	r := gin.Default()

	r.Use(cors.Default())

	userHandler := handlers.NewUserHandler(storage.NewUserStorage(dbConnection.DB()))
	r.GET("/users/:id", userHandler.Show)
	r.DELETE("/users/:id", userHandler.Destroy)
	r.PUT("/users/:id", userHandler.Update)
	r.GET("/users", userHandler.Index)
	r.POST("/users", userHandler.Create)

	return r
}
