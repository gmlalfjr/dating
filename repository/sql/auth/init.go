package auth

import (
	"dating/domains/entities"
	"dating/domains/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(ctx *gin.Context, request *models.RegisterRequest) (*entities.User, error)
	FindOneUserByEmail(ctx *gin.Context, email string) (*entities.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func InitAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}
