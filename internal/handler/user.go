package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/irawankilmer/lms_backend/internal/models"
	"github.com/irawankilmer/lms_backend/internal/service"
	"github.com/irawankilmer/lms_backend/internal/utils"
	"strconv"
)

// Inisialisasi validator
var validate = validator.New()

// Inisialisasi services
var userService *service.UserService
var authService *service.AuthService

// SetServices digunakan untuk menginisialisasi services di vercel.go
func SetServices(u *service.UserService, a *service.AuthService) {
	userService = u
	authService = a
}

// Struct untuk input login
type LoginInput struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// @Summary Login user
// @Description Login user with username or email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param loginInput body LoginInput true "Login Input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSON(c, 400, "Invalid input", nil, err.Error())
		return
	}

	token, err := authService.Login(input.Identifier, input.Password)
	if err != nil {
		utils.JSON(c, 401, "Authentication failed", nil, err.Error())
		return
	}

	utils.JSON(c, 200, "Login successful", gin.H{"token": token}, "")
}

// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSON(c, 400, "Invalid input", nil, err.Error())
		return
	}

	if err := validate.Struct(&user); err != nil {
		utils.JSON(c, 400, "Validation failed", nil, err.Error())
		return
	}

	if err := userService.CreateUser(&user); err != nil {
		utils.JSON(c, 500, "Failed to create user", nil, err.Error())
		return
	}

	utils.JSON(c, 201, "User created successfully", user, "")
}

// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil {
		utils.JSON(c, 500, "Failed to retrieve users", nil, err.Error())
		return
	}

	utils.JSON(c, 200, "Users retrieved successfully", users, "")
}

// @Summary Get user by ID
// @Description Retrieve a user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.JSON(c, 400, "Invalid user ID", nil, err.Error())
		return
	}

	user, err := userService.GetUserByID(uint(id))
	if err != nil {
		utils.JSON(c, 404, "User not found", nil, err.Error())
		return
	}

	utils.JSON(c, 200, "User retrieved successfully", user, "")
}

// @Summary Update a user
// @Description Update a user's information by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User Data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.JSON(c, 400, "Invalid user ID", nil, err.Error())
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSON(c, 400, "Invalid input", nil, err.Error())
		return
	}

	if err := validate.Struct(&user); err != nil {
		utils.JSON(c, 400, "Validation failed", nil, err.Error())
		return
	}

	user.ID = uint(id)
	if err := userService.UpdateUser(&user); err != nil {
		utils.JSON(c, 500, "Failed to update user", nil, err.Error())
		return
	}

	utils.JSON(c, 200, "User updated successfully", user, "")
}

// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.JSON(c, 400, "Invalid user ID", nil, err.Error())
		return
	}

	if err := userService.DeleteUser(uint(id)); err != nil {
		utils.JSON(c, 500, "Failed to delete user", nil, err.Error())
		return
	}

	utils.JSON(c, 204, "User deleted successfully", nil, "")
}
