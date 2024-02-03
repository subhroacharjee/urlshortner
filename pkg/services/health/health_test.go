package healthservice

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subhroacharjee/urlshortner/config"
	"github.com/subhroacharjee/urlshortner/utils/db"
	"go.uber.org/zap"
)

type mockConfig struct{}

// GetDSN implements config.Config.
func (mockConfig) GetDSN() string {
	panic("unimplemented")
}

// GetEnv implements config.Config.
func (mockConfig) GetEnv() config.Env {
	return config.Test
}

// GetPort implements config.Config.
func (mockConfig) GetPort() uint {
	panic("unimplemented")
}

func TestHealthService(t *testing.T) {
	ctx := context.Background()
	config := mockConfig{}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}

	db, err := db.NewDb(config)
	if err != nil {
		t.Fatal(err)
	}

	healthService := HealthService{
		Logger: logger,
		DB:     db,
	}

	t.Run("Check", func(t *testing.T) {
		err := healthService.Check(ctx)
		assert.Nil(t, err)
	})
}
