package controllers

import (
	"dating/app/config"
	"dating/controllers/auth"
)

type Controller struct {
	AuthController auth.IAuthController
}

func InitController(services config.Services) *Controller {
	return &Controller{
		AuthController: auth.InitAuthController(services.AuthService),
	}
}
