package premium_user

import (
	"dating/domains/entities"
	"dating/domains/models"
	"github.com/gin-gonic/gin"
	"time"
)

func (p *PremiumUserRepository) Create(c *gin.Context, userId int, expiredAt time.Time) error {
	response := &entities.PremiumUser{
		UserID:    userId,
		ExpiredAt: expiredAt,
	}
	if err := p.db.Create(response).Error; err != nil {
		return err
	}

	return nil
}
func (p *PremiumUserRepository) FindByUserId(c *gin.Context, userId int) (*entities.PremiumUser, error) {
	result := &entities.PremiumUser{}

	if err := p.db.Debug().Model(&entities.PremiumUser{}).Where("user_id = ? AND DATE(expired_at) > CURRENT_DATE", userId).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PremiumUserRepository) FindByEmail(c *gin.Context, email string) (*models.LoginUserPremium, error) {
	response := &models.LoginUserPremium{}

	if err := p.db.Where("email = ?", email).Model(&entities.User{}).
		Select("users.*, pu.id, pu.expired_at").
		Joins("LEFT JOIN premium_users pu ON users.id = pu.user_id").
		Scan(response).Error; err != nil {
		return nil, err
	}

	return response, nil
}
