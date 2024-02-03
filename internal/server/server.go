package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/subhroacharjee/urlshortner/config"
)

type Server struct {
	G      *gin.Engine
	Config config.Config
}

func NewServer(g *gin.Engine, conf config.Config) *Server {
	return &Server{
		G:      g,
		Config: conf,
	}
}

func (s *Server) Run() error {
	port := fmt.Sprintf(":%d", s.Config.GetPort())
	return s.G.Run(port)
}
