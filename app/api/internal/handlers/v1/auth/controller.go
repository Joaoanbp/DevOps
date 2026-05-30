package v1_auth

import (
	"api/internal/framework"
)

type AuthController struct {
	*framework.BaseController
}

func NewLogin(options *framework.BaseController) *AuthController {
	return &AuthController{
		BaseController: options,
	}
}

func NewSignup(options *framework.BaseController) *AuthController {
	return &AuthController{
		BaseController: options,
	}
}