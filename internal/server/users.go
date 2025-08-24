package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) usersHandler(c *gin.Context) {
	result := s.db.AllUsers()

	c.JSON(http.StatusOK, result)
}

func (s *Server) userHandler(c *gin.Context) {
	result := s.db.GetUserById(c.Param("id"))

	c.JSON(http.StatusOK, result)
}

func (s *Server) createUserHandler(c *gin.Context) {
	userAttr := map[string]string{"Login": c.PostForm("login"), "FullName": c.PostForm("full_name")}
	log.Print(userAttr)
	result := s.db.CreateUser(userAttr)

	c.JSON(http.StatusCreated, result)
}

func (s *Server) deleteUserHandler(c *gin.Context) {
	s.db.DeleteUser(c.Param("id"))
	c.Status(http.StatusNoContent)
}

func (s *Server) updateUserHandler(c *gin.Context) {
	userAttr := map[string]string{"FullName": c.PostForm("full_name")}
	log.Print(userAttr)
	result := s.db.UpdateUser(c.Param("id"), userAttr)

	c.JSON(http.StatusOK, result)
}
