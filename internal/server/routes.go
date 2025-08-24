package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/users/:id", s.userHandler)
	r.DELETE("/users/:id", s.deleteUserHandler)
	r.PUT("/users/:id", s.updateUserHandler)
	r.GET("/users", s.usersHandler)
	r.POST("/users", s.createUserHandler)

	return r
}
