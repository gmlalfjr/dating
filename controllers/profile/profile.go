package profile

import (
	"dating/domains/models"
	"dating/response"
	"dating/utils"
	"github.com/gin-gonic/gin"
)

func (p *ProfileController) ListProfile(c *gin.Context) {
	var (
		request *models.BaseRequest = utils.MapRequest(c, &models.BaseRequest{}, []string{})
	)

	res, totalData, totalPage, err := p.profileService.GetListProfile(c, &request.User, request.QueryParam)
	if err != nil {
		response := response.Response{}
		response.Error(c, err)
		return
	}
	response := response.Response{
		Data:       res,
		TotalPages: totalPage,
		TotalData:  int(totalData),
	}
	response.Success(c)
}
