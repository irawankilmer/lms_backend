package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/irawankilmer/lms_backend/internal/dto"
	"github.com/irawankilmer/lms_backend/internal/models"
	"github.com/irawankilmer/lms_backend/internal/service"
	"github.com/irawankilmer/lms_backend/internal/utils"
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

// @Summary Login user
// @Description Login user with username or email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param dto.loginInput body dto.LoginInput true "Login Input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	var input dto.LoginInput

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

// @Summary Request Password Reset
// @Description Mengirimkan email untuk reset password
// @Tags auth
// @Accept json
// @Produce json
// @Param email body dto.PasswordResetRequestInput true "Email untuk reset password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /password-reset/request [post]
func RequestPasswordReset(c *gin.Context) {
	var input dto.PasswordResetRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSON(c, 400, "Invalid input", nil, err.Error())
		return
	}

	token, err := authService.GeneratePasswordResetToken(input.Email)
	if err != nil {
		utils.JSON(c, 400, "Failed to generate reset token", nil, err.Error())
		return
	}

	err = authService.SendPasswordResetEmail(input.Email, token)
	if err != nil {
		utils.JSON(c, 500, "Failed to send reset email", nil, err.Error())
		return
	}

	utils.JSON(c, 200, "Password reset email sent", nil, "")
}

// @Summary Reset Password
// @Description Reset password menggunakan token
// @Tags auth
// @Accept json
// @Produce json
// @Param passwordReset body dto.PasswordResetInput true "Token dan Password baru"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /password-reset/reset [post]
func ResetPassword(c *gin.Context) {
	var input dto.PasswordResetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSON(c, 400, "Invalid input", nil, err.Error())
		return
	}

	err := authService.ResetPassword(input.Token, input.Password)
	if err != nil {
		utils.JSON(c, 400, "Failed to reset password", nil, err.Error())
		return
	}

	utils.JSON(c, 200, "Password reset successful", nil, "")
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
