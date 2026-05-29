package v1_auth

import (
	"api/internal/framework"
)

type AuthController struct {
	*framework.BaseController
}

func New(options *framework.BaseController) *AuthController {
	return &AuthController{
		BaseController: options,
	}
}