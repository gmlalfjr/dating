package auth

import (
	"dating/domains/models"
	"dating/response"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (a *AuthService) Register(c *gin.Context, request *models.RegisterRequest) (*models.RegisterResponse, error) {
	pass, err := a.hashingPassword(c, request.Password, request.PasswordRepeat)
	if err != nil {
		return nil, err
	}

	res, err := a.AuthRepository.FindOneUserByEmail(c, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if res != nil {
		return nil, response.NewError(errors.New("username already exist"), 400, "username already exist")
	}
	request.Password = pass
	res, err = a.AuthRepository.CreateUser(c, request)
	if err != nil {
		return nil, err
	}
	return &models.RegisterResponse{Email: res.Email}, nil

}

func (a *AuthService) hashingPassword(c *gin.Context, passwordOne, passwordTwo string) (string, error) {

	if passwordOne != passwordTwo {
		return "", response.NewError(errors.New("password not match"), 400, "password not match")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOne), 14)
	if err != nil {
		return "", response.NewError(err, 400, "failed register")

	}
	return string(hash), nil
}
