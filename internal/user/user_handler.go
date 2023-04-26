package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repository UserRepository
}

func NewUserHandler(repository UserRepository) *UserHandler {
	return &UserHandler{
		Repository: repository,
	}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/users", h.getUsers)
	router.GET("/users/:id", h.getUser)
	router.POST("/users", h.createUser)
	router.PUT("/users/:id", h.updateUser)
	router.DELETE("/users/:id", h.deleteUser)
}

func (h *UserHandler) getUsers(c *gin.Context) {
	users, err := h.Repository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.Repository.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}
	user, err := h.Repository.Create(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}
	user, err := h.Repository.Update(id, updatedUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Repository.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
