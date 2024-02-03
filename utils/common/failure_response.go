package common

import "github.com/gin-gonic/gin"

func BadRequest(c *gin.Context, err error) {
	response(c, nil, err.Error(), 400)
}

func BadGateway(c *gin.Context, err error) {
	response(c, nil, err.Error(), 500)
}
