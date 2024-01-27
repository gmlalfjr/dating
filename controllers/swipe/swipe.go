package swipe

import (
	"dating/domains/models"
	"dating/response"
	"dating/utils"
	"github.com/gin-gonic/gin"
)

func (s *SwipeController) SwipeLeft(c *gin.Context) {
	var (
		request *models.BaseRequest = utils.MapRequest(c, &models.BaseRequest{}, []string{"id"})
	)
	err := s.swipeService.SwipeLeftService(c, request)
	if err != nil {
		response := response.Response{}
		response.Error(c, err)
		return
	}
	response := response.Response{}
	response.Success(c)
}

func (s *SwipeController) SwipeRight(c *gin.Context) {
	var (
		request *models.BaseRequest = utils.MapRequest(c, &models.BaseRequest{}, []string{"id"})
	)
	res, err := s.swipeService.SwipeRightService(c, request)
	if err != nil {
		response := response.Response{}
		response.Error(c, err)
		return
	}
	response := response.Response{Data: res}
	response.Success(c)
}
