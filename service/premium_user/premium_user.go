package premium_user

import (
	"dating/domains/models"
	"dating/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func (p *PremiumUserService) CreatePremiumUserService(c *gin.Context, req *models.BaseRequest) error {
	res, err := p.PremiumUserRepository.FindByUserId(c, req.User.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if res != nil {
		return response.NewError(errors.New("user already premium"), 400, "user already premium")
	}
	err = p.PremiumUserRepository.Create(c, req.User.ID, time.Now().Add(time.Hour*168))
	if err != nil {
		return err
	}
	return nil
}
