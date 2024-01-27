package premium_user

import (
	"dating/domains/models"
	"dating/repository/sql/premium_user"
	"github.com/gin-gonic/gin"
)

type IPremiumUserService interface {
	CreatePremiumUserService(c *gin.Context, req *models.BaseRequest) error
}

type PremiumUserService struct {
	PremiumUserRepository premium_user.IPremiumUserRepository
}

func InitPremiumUserServices(premiumUserService *PremiumUserService) IPremiumUserService {
	return premiumUserService
}
