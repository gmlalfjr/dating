package swipe

import (
	"dating/domains/models"
	"dating/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *SwipeService) SwipeLeftService(c *gin.Context, userReq *models.BaseRequest) error {
	if userReq.User.ID == userReq.ID {
		return response.NewError(errors.New("cannot swipe yourself"), 400, "cannot swipe yourself")
	}
	err := s.SwipeRepository.CreateSwipe(c, userReq.User.ID, userReq.ID, "pass")
	if err != nil {
		return err
	}
	return nil

}

func (s *SwipeService) SwipeRightService(c *gin.Context, userReq *models.BaseRequest) (*models.SwipeRightResponse, error) {
	var resSwipe = &models.SwipeRightResponse{}
	if userReq.User.ID == userReq.ID {
		return nil, response.NewError(errors.New("cannot swipe yourself"), 400, "cannot swipe yourself")
	}
	// TODO: find first already like or not
	res, err := s.SwipeRepository.FindOneByUserIdAndSwipedUserId(c, userReq.User.ID, userReq.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = s.SwipeRepository.CreateSwipe(c, userReq.User.ID, userReq.ID, "like")
	if err != nil {
		return nil, err
	}
	if res != nil {
		resSwipe.IsMatch = true
	}

	return resSwipe, nil
}
