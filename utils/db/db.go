package db

import (
	"github.com/subhroacharjee/urlshortner/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDb(conf config.Config) (*gorm.DB, error) {
	if conf.GetEnv() == config.Test {
		return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}

	return gorm.Open(postgres.Open(conf.GetDSN()), &gorm.Config{})
}
