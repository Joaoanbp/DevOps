package handler_v1

import (
	"api/internal/framework"
)

type IndexController struct {
	*framework.BaseController
}

func New(options *framework.BaseController) *IndexController {
	return &IndexController{
		BaseController: options,
	}
}

func (this *IndexController) Get() {
	// this.Ctx.WriteString("Hello, World!")
}