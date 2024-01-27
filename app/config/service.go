package config

import (
	"dating/service/auth"
	"dating/service/profile"
)

type Services struct {
	AuthService    auth.IAuthService
	ProfileService profile.IProfileService
}

func InitServices(repository Repository) *Services {
	authService := &auth.AuthService{
		AuthRepository: repository.AuthRepo,
	}
	profileService := &profile.ProfileService{
		ProfileRepository: repository.ProfileRepo,
		AuthRepository:    repository.AuthRepo,
	}

	return &Services{
		AuthService:    auth.InitAuthServices(authService),
		ProfileService: profileService,
	}
}
