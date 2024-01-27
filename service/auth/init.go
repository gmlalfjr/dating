package auth

import (
	"dating/domains/models"
	"dating/repository/sql/auth"
	"github.com/gin-gonic/gin"
)

type IAuthService interface {
	Register(c *gin.Context, request *models.RegisterRequest) (*models.RegisterResponse, error)
}

type AuthService struct {
	AuthRepository auth.IAuthRepository
}

func InitAuthServices(authService *AuthService) IAuthService {
	return authService
}
