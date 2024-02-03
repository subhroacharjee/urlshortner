package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/subhroacharjee/urlshortner/pkg/controllers/health"
	"github.com/subhroacharjee/urlshortner/pkg/controllers/urls"
)

func InitRouter(g *gin.Engine) {
	router := g.Group("/api")
	// define routes and controller.
	health.InitHealthController(router)
	urls.InitUrlController(router)
}
