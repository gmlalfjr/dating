package profile

import (
	"dating/domains/entities"
	"gorm.io/gorm"
)

type IProfileRepository interface {
	FindProfileExcludeUserId(userId int, limit, offset int) ([]*entities.User, int64, error)
}

type ProfileRepository struct {
	db *gorm.DB
}

func InitProfileRepository(db *gorm.DB) IProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}
