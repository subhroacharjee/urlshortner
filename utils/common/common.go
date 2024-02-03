package common

import "github.com/gin-gonic/gin"

func response(c *gin.Context, data any, message string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"data":    data,
		"message": message,
	})
}
