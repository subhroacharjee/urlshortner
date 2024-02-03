package healthservice

import (
	"context"

	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HealthService struct {
	dig.In

	Logger *zap.Logger
	DB     *gorm.DB
}

func (h *HealthService) Check(ctx context.Context) error {
	defer func() {
		h.Logger.Info("HealthService.Check has ended")
	}()
	h.Logger.Info("HealthService.Check has started")
	db, err := h.DB.DB()
	if err != nil {
		h.Logger.Error("Check has failed with ", zap.Error(err))
		return err
	}

	if err = db.Ping(); err != nil {
		h.Logger.Error("Check has failed with ", zap.Error(err))
		return err
	}

	h.Logger.Info("DB works")
	return nil
}
