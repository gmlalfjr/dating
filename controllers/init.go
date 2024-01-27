package controllers

import (
	"dating/app/config"
	"dating/controllers/auth"
	"dating/controllers/profile"
	"dating/middleware"
)

type Controller struct {
	AuthController    auth.IAuthController
	AuthMiddleware    middleware.IAuthMiddleware
	ProfileController profile.IProfileController
}

func InitController(services config.Services) *Controller {
	return &Controller{
		AuthMiddleware:    middleware.InitAuthMiddleware(),
		AuthController:    auth.InitAuthController(services.AuthService),
		ProfileController: profile.InitProfileController(services.ProfileService),
	}
}
