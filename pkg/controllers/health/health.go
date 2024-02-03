package health

import (
	"github.com/gin-gonic/gin"
	healthservice "github.com/subhroacharjee/urlshortner/pkg/services/health"
	"github.com/subhroacharjee/urlshortner/utils/common"
	"github.com/subhroacharjee/urlshortner/utils/container"
)

type HealthController struct {
	HealthService healthservice.HealthService
}

func NewHealthController() (*HealthController, error) {
	var healthService healthservice.HealthService
	if err := container.Container.Invoke(func(hs healthservice.HealthService) {
		healthService = hs
	}); err != nil {
		return nil, err
	}
	return &HealthController{
		HealthService: healthService,
	}, nil
}

func (h *HealthController) Check(c *gin.Context) {
	if err := h.HealthService.Check(c); err != nil {
		common.BadGateway(c, err)
		return
	}
	common.Success(c, nil, "Health check successful")
}

func InitHealthController(r *gin.RouterGroup) error {
	healthController, err := NewHealthController()
	if err != nil {
		return err
	}
	r.GET("/health", healthController.Check)
	return nil
}
