package auth

import "dating/repository/sql/auth"

type IAuthService interface {
}

type AuthService struct {
	AuthRepository auth.IAuthRepository
}

func InitAuthServices(authService *AuthService) IAuthService {
	return authService
}
