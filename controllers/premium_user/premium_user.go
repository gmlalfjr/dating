package premium_user

import (
	"dating/domains/models"
	"dating/response"
	"dating/utils"
	"github.com/gin-gonic/gin"
)

func (p *PremiumUserController) CreatePremiumUser(c *gin.Context) {
	var (
		request *models.BaseRequest = utils.MapRequest(c, &models.BaseRequest{}, []string{})
	)
	err := p.premiumUserService.CreatePremiumUserService(c, request)
	if err != nil {
		response := response.Response{}
		response.Error(c, err)
		return
	}
	response := response.Response{}
	response.Success(c)
}
