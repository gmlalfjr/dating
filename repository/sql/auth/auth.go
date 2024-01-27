package auth

import (
	"dating/domains/entities"
	"dating/domains/models"
	"github.com/gin-gonic/gin"
)

func (a *AuthRepository) CreateUser(ctx *gin.Context, request *models.RegisterRequest) (*entities.User, error) {
	response := &entities.User{
		Email:    request.Email,
		Password: request.Password,
		FullName: request.FullName,
		Address:  request.Address,
		Age:      request.Age,
		Sex:      request.Sex,
	}
	if err := a.db.Create(response).Error; err != nil {
		return nil, err
	}

	return response, nil
}

func (a *AuthRepository) FindOneUserByEmail(ctx *gin.Context, email string) (*entities.User, error) {
	response := &entities.User{}
	if err := a.db.Where("email  = ?", email).First(response).Error; err != nil {
		return nil, err
	}

	return response, nil
}
