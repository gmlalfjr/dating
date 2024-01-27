package auth

import "gorm.io/gorm"

type IAuthRepository interface {
}

type AuthRepository struct {
	db *gorm.DB
}

func InitAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}
