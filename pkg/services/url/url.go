package urlservice

import (
	"context"

	"github.com/subhroacharjee/urlshortner/config"
	"github.com/subhroacharjee/urlshortner/models/urls"
	"github.com/subhroacharjee/urlshortner/utils/common"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UrlService struct {
	dig.In

	Logger *zap.Logger
	Db     *gorm.DB
	Config config.Config
}

func (u *UrlService) HashUrl(ctx context.Context, url string) (*hashResponse, error) {
	defer func() {
		u.Logger.Info("HashUrl has stopped")
	}()

	db := u.Db
	var count int64

	if err := db.Transaction(func(tx *gorm.DB) error {
		u.Logger.Info("Running transaction to get the current count")
		return tx.Model(&urls.Urls{}).Count(&count).Error
	}); err != nil {
		u.Logger.Error("Tnx failed", zap.Error(err))
		return nil, err
	}

	hash := common.HashInt64(count)[:u.Config.GetHashBreak()]
	newUrl := urls.Urls{Url: url, Hash: hash}
	result := db.Create(&newUrl)
	if err := result.Error; err != nil {
		u.Logger.Error("Creation failed", zap.Error(err))
		return nil, err
	}

	u.Logger.Info("New Url hash has been created", zap.Any("newUrl", newUrl))
	return &hashResponse{
		Hash: hash,
		Url:  url,
	}, nil
}

func (u *UrlService) GetUrl(ctx context.Context, hash string) (*getResponse, error) {
	db := u.Db
	logger := u.Logger
	defer func() {
		logger.Info("GetUrl has stopped")
	}()
	logger.Info("Get Url started")

	url := urls.Urls{Hash: hash}
	if err := db.Model(&url).First(&url).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Info("Hash was not found", zap.String("hash", hash))
			return nil, nil
		}
		logger.Error("Not able to find the data", zap.Error(err))
		return nil, err
	}

	return &getResponse{
		Url: url.Url,
	}, nil
}
