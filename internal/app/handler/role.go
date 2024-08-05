package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/lms_backend/internal/config"
	"github.com/irawankilmer/lms_backend/internal/models"
	"net/http"
)

// CreateRole godoc
// @Summary Create a new role
// @Description Create a new role with the given details
// @Tags roles
// @Accept json
// @Produce json
// @Param role body models.Role true "Role to create"
// @Success 200 {object} models.Role
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles [post]
func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}

// GetRoles godoc
// @Summary Get all roles
// @Description Get a list of all roles
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {array} models.Role
// @Failure 500 {object} map[string]interface{}
// @Router /roles [get]
func GetRoles(c *gin.Context) {
	var roles []models.Role
	if err := config.DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// GetRole godoc
// @Summary Get a role by ID
// @Description Get a role by its ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} models.Role
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles/{id} [get]
func GetRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// UpdateRole godoc
// @Summary Update a role by ID
// @Description Update a role by its ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param role body models.Role true "Role to update"
// @Success 200 {object} models.Role
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles/{id} [put]
func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Role not found"})
		return
	}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}

// DeleteRole godoc
// @Summary Delete a role by ID
// @Description Delete a role by its ID
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 204
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles/{id} [delete]
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Role{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
