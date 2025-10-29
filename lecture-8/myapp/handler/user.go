package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ITMO-students/lecture-8/myapp/service"
)

type UserHandler struct {
	service service.UserService
}

func New(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
