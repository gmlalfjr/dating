package auth

import (
	"dating/service/auth"
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Register(c *gin.Context)
}

type AuthController struct {
	authService auth.IAuthService
}

func InitAuthController(auth auth.IAuthService) IAuthController {
	return &AuthController{
		authService: auth,
	}
}
