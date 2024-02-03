package logger

import (
	"github.com/subhroacharjee/urlshortner/config"
	"go.uber.org/zap"
)

func NewLogger(conf config.Config) (*zap.Logger, error) {
	if conf.GetEnv() == config.Prod {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}
