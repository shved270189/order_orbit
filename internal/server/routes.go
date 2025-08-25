package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"order_orbit/internal/handlers"
)

func RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/users/:id", handlers.Users.Show)
	r.DELETE("/users/:id", handlers.Users.Destroy)
	r.PUT("/users/:id", handlers.Users.Update)
	r.GET("/users", handlers.Users.Index)
	r.POST("/users", handlers.Users.Create)

	return r
}
