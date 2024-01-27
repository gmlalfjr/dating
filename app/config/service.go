package config

import (
	"dating/service/auth"
)

type Services struct {
	AuthService auth.IAuthService
}

func InitServices(repository Repository) *Services {
	authService := &auth.AuthService{
		AuthRepository: repository,
	}

	return &Services{
		AuthService: auth.InitAuthServices(authService),
	}
}
