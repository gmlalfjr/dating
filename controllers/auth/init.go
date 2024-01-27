package auth

import (
	"dating/service/auth"
)

type IAuthController interface {
}

type AuthController struct {
	authService auth.IAuthService
}

func InitAuthController(auth auth.IAuthService) IAuthController {
	return &AuthController{}
}
