package service

import (
	"github.com/irawankilmer/lms_backend/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(user *models.User) error {
	return s.DB.Create(user).Error
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user in the database
func (s *UserService) UpdateUser(user *models.User) error {
	return s.DB.Save(user).Error
}

// DeleteUser deletes a user from the database by ID
func (s *UserService) DeleteUser(id uint) error {
	if err := s.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetAllUsers retrieves all users from the database
func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
