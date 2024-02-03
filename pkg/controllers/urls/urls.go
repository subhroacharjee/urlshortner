package urls

import (
	"github.com/gin-gonic/gin"
	urlservice "github.com/subhroacharjee/urlshortner/pkg/services/url"
	"github.com/subhroacharjee/urlshortner/utils/common"
	"github.com/subhroacharjee/urlshortner/utils/container"
)

type UrlController struct {
	UrlService urlservice.UrlService
}

type hashUrlRequestBody struct {
	Url string `json:"url" binding:"required"`
}

func NewUrlController() (*UrlController, error) {
	var urlService urlservice.UrlService
	if err := container.Container.Invoke(func(service urlservice.UrlService) {
		urlService = service
	}); err != nil {
		return nil, err
	}

	return &UrlController{
		UrlService: urlService,
	}, nil
}

func (u *UrlController) HashUrl(c *gin.Context) {
	var body hashUrlRequestBody
	if err := c.BindJSON(&body); err != nil {
		common.BadRequest(c, err)
		return
	}

	data, err := u.UrlService.HashUrl(c, body.Url)
	if err != nil {
		common.BadGateway(c, err)
		return
	}

	common.Created(c, data, "hashing done")
}

func (u *UrlController) GetUrl(c *gin.Context) {
	hash := c.Param("hash")
	redirect := c.Query("redirect")
	data, err := u.UrlService.GetUrl(c, hash)
	if err != nil {
		common.BadGateway(c, err)
	}
	if data == nil {
		common.NotFound(c)
		return
	}
	if redirect == "true" {
		c.Redirect(301, data.Url)
	}
	common.Success(c, data, "url fetched successfully")
	return
}

func InitUrlController(r *gin.RouterGroup) error {
	urlController, err := NewUrlController()
	if err != nil {
		return err
	}

	r.GET("/:hash", urlController.GetUrl)
	r.POST("/hash", urlController.HashUrl)
	return nil
}
