package handler_v1

import (
	"api/internal/framework"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	*framework.BaseController
}

func New(options *framework.BaseController) *IndexController {
	return &IndexController{
		BaseController: options,
	}
}

func (this *IndexController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"mogged": true,
	})
}