package config

import (
	"dating/service/auth"
	"dating/service/premium_user"
	"dating/service/profile"
	"dating/service/swipe"
)

type Services struct {
	AuthService        auth.IAuthService
	ProfileService     profile.IProfileService
	SwipeService       swipe.ISwipeService
	PremiumUserService premium_user.IPremiumUserService
}

func InitServices(repository Repository) *Services {
	authService := &auth.AuthService{
		AuthRepository:  repository.AuthRepo,
		PremiumUserRepo: repository.PremiumUserRepo,
	}
	profileService := &profile.ProfileService{
		ProfileRepository: repository.ProfileRepo,
		AuthRepository:    repository.AuthRepo,
	}
	swipeService := &swipe.SwipeService{
		AuthRepository:  repository.AuthRepo,
		SwipeRepository: repository.SwipeRepo,
	}
	premiumUserService := &premium_user.PremiumUserService{
		PremiumUserRepository: repository.PremiumUserRepo,
	}

	return &Services{
		AuthService:        auth.InitAuthServices(authService),
		ProfileService:     profile.InitProfileServices(profileService),
		SwipeService:       swipe.InitSwipeServices(swipeService),
		PremiumUserService: premium_user.InitPremiumUserServices(premiumUserService),
	}
}
