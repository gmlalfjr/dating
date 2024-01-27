package profile

import (
	"dating/service/profile"
	"github.com/gin-gonic/gin"
)

type IProfileController interface {
	ListProfile(c *gin.Context)
}

type ProfileController struct {
	profileService profile.IProfileService
}

func InitProfileController(profileService profile.IProfileService) IProfileController {
	return &ProfileController{
		profileService: profileService,
	}
}
