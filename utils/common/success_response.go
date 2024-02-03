package common

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data any, message string) {
	response(c, data, message, 200)
}

func Created(c *gin.Context, data any, message string) {
	response(c, data, message, 201)
}
