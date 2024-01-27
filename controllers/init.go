package controllers

import (
	"dating/app/config"
	"dating/controllers/auth"
	"dating/middleware"
)

type Controller struct {
	AuthController auth.IAuthController
	AuthMiddleware middleware.IAuthMiddleware
}

func InitController(services config.Services) *Controller {
	return &Controller{
		AuthController: auth.InitAuthController(services.AuthService),
		AuthMiddleware: middleware.InitAuthMiddleware(),
	}
}
