package controllers

import (
	"dating/app/config"
	"dating/controllers/auth"
	"dating/controllers/premium_user"
	"dating/controllers/profile"
	"dating/controllers/swipe"
	"dating/middleware"
)

type Controller struct {
	AuthController        auth.IAuthController
	AuthMiddleware        middleware.IAuthMiddleware
	ProfileController     profile.IProfileController
	SwipeController       swipe.ISwipeController
	PremiumUserController premium_user.IPremiumUserController
}

func InitController(services config.Services) *Controller {
	return &Controller{
		AuthMiddleware:        middleware.InitAuthMiddleware(),
		AuthController:        auth.InitAuthController(services.AuthService),
		ProfileController:     profile.InitProfileController(services.ProfileService),
		SwipeController:       swipe.InitSwipeController(services.SwipeService),
		PremiumUserController: premium_user.InitPremiumUserController(services.PremiumUserService),
	}
}
