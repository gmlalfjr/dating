package auth

import (
	"dating/domains/models"
	"dating/repository/sql/auth"
	"dating/repository/sql/premium_user"
	"github.com/gin-gonic/gin"
)

type IAuthService interface {
	Register(c *gin.Context, request *models.RegisterRequest) (*models.RegisterResponse, error)
	Login(c *gin.Context, request *models.LoginRequest) (*models.LoginResponse, error)
}

type AuthService struct {
	AuthRepository  auth.IAuthRepository
	PremiumUserRepo premium_user.IPremiumUserRepository
}

func InitAuthServices(authService *AuthService) IAuthService {
	return authService
}
