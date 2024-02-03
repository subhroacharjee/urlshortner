package main

import (
	"github.com/subhroacharjee/urlshortner/internal/server"
	"github.com/subhroacharjee/urlshortner/pkg/controllers"
	"github.com/subhroacharjee/urlshortner/utils/container"
)

func run() error {
	container.InitDependencies()
	return container.Container.Invoke(func(s *server.Server) error {
		controllers.InitRouter(s.G)
		return s.Run()
	})
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
