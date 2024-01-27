package auth

import (
	"dating/domains/models"
	"dating/response"
	"dating/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func (a *AuthController) Register(c *gin.Context) {
	var (
		request *models.BaseRequest = utils.MapRequest(c, &models.BaseRequest{BodyData: &models.RegisterRequest{}}, []string{})
	)
	data := request.BodyData.(*models.RegisterRequest)
	if len(data.Password) < 8 {
		err := response.NewError(errors.New("Password must 8 characters"), 400, "Password must 8 characters")
		response := response.Response{}
		response.Error(c, err)
		return
	}
	res, err := a.authService.Register(c, data)
	if err != nil {
		response := response.Response{}
		response.Error(c, err)
		return
	}
	response := response.Response{Data: res}
	response.Success(c)
}

func (a *AuthController) Login(c *gin.Context) {
	var (
		request *models.BaseRequest = utils.MapRequest(c, &models.BaseRequest{BodyData: &models.LoginRequest{}}, []string{})
	)
	data := request.BodyData.(*models.LoginRequest)
	res, err := a.authService.Login(c, data)
	if err != nil {
		response := response.Response{}
		response.Error(c, err)
		return
	}
	response := response.Response{Data: res}
	response.Success(c)
}
