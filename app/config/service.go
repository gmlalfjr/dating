package config

import (
	"dating/service/auth"
	"dating/service/profile"
	"dating/service/swipe"
)

type Services struct {
	AuthService    auth.IAuthService
	ProfileService profile.IProfileService
	SwipeService   swipe.ISwipeService
}

func InitServices(repository Repository) *Services {
	authService := &auth.AuthService{
		AuthRepository: repository.AuthRepo,
	}
	profileService := &profile.ProfileService{
		ProfileRepository: repository.ProfileRepo,
		AuthRepository:    repository.AuthRepo,
	}
	swipeService := &swipe.SwipeService{
		AuthRepository:  repository.AuthRepo,
		SwipeRepository: repository.SwipeRepo,
	}

	return &Services{
		AuthService:    auth.InitAuthServices(authService),
		ProfileService: profile.InitProfileServices(profileService),
		SwipeService:   swipe.InitSwipeServices(swipeService),
	}
}
