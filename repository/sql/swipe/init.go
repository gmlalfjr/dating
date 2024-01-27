package swipe

import (
	"dating/domains/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ISwipeRepository interface {
	CreateSwipe(c *gin.Context, userId int, swipedUserId int, action string) error
	FindOneByUserIdAndSwipedUserId(c *gin.Context, userId int, swipedUserId int) (*entities.Swipe, error)
}

type SwipeRepository struct {
	db *gorm.DB
}

func InitSwipeRepository(db *gorm.DB) ISwipeRepository {
	return &SwipeRepository{
		db: db,
	}
}
