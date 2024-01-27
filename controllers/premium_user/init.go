package premium_user

import (
	"dating/service/premium_user"
	"github.com/gin-gonic/gin"
)

type IPremiumUserController interface {
	CreatePremiumUser(c *gin.Context)
}

type PremiumUserController struct {
	premiumUserService premium_user.IPremiumUserService
}

func InitPremiumUserController(premiumUserService premium_user.IPremiumUserService) IPremiumUserController {
	return &PremiumUserController{
		premiumUserService: premiumUserService,
	}
}
