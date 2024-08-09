package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func JSON(c *gin.Context, status int, message string, data interface{}, err string) {
	c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	})
}
