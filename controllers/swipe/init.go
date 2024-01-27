package swipe

import (
	"dating/service/swipe"
	"github.com/gin-gonic/gin"
)

type ISwipeController interface {
	SwipeLeft(c *gin.Context)
	SwipeRight(c *gin.Context)
}

type SwipeController struct {
	swipeService swipe.ISwipeService
}

func InitSwipeController(auth swipe.ISwipeService) ISwipeController {
	return &SwipeController{
		swipeService: auth,
	}
}
