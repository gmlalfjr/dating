package premium_user

import (
	"dating/domains/entities"
	"dating/domains/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type IPremiumUserRepository interface {
	Create(c *gin.Context, userId int, expiredAt time.Time) error
	FindByUserId(c *gin.Context, userId int) (*entities.PremiumUser, error)
	FindByEmail(c *gin.Context, email string) (*models.LoginUserPremium, error)
}

type PremiumUserRepository struct {
	db *gorm.DB
}

func InitPremiumUserRepository(db *gorm.DB) IPremiumUserRepository {
	return &PremiumUserRepository{
		db: db,
	}
}
