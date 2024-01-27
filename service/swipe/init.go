package swipe

import (
	"dating/domains/models"
	"dating/repository/sql/auth"
	"dating/repository/sql/swipe"
	"github.com/gin-gonic/gin"
)

type ISwipeService interface {
	SwipeLeftService(c *gin.Context, userReq *models.BaseRequest) error
	SwipeRightService(c *gin.Context, userReq *models.BaseRequest) (*models.SwipeRightResponse, error)
}

type SwipeService struct {
	AuthRepository  auth.IAuthRepository
	SwipeRepository swipe.ISwipeRepository
}

func InitSwipeServices(swipeService *SwipeService) ISwipeService {
	return swipeService
}
