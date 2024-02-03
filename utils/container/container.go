package container

import (
	"github.com/gin-gonic/gin"
	"github.com/subhroacharjee/urlshortner/config"
	"github.com/subhroacharjee/urlshortner/internal/server"
	"github.com/subhroacharjee/urlshortner/utils/db"
	"github.com/subhroacharjee/urlshortner/utils/logger"
	"go.uber.org/dig"
)

var Container = dig.New()

// add all the dependency here
func InitDependencies() {
	Container.Provide(config.NewConfig)
	Container.Provide(newGinEngine)
	Container.Provide(logger.NewLogger)
	Container.Provide(server.NewServer)

	// functions which do not provide any new interface but is equally important
	Container.Provide(db.NewDb)
	if err := Container.Invoke(db.MigrateModels); err != nil {
		panic(err)
	}
}

func newGinEngine() *gin.Engine {
	return gin.Default()
}
