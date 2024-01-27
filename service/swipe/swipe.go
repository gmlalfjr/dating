package swipe

import (
	"dating/domains/models"
	"dating/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func (s *SwipeService) SwipeLeftService(c *gin.Context, userReq *models.BaseRequest) error {
	//var dateExpired time.Time
	if userReq.User.ID == userReq.ID {
		return response.NewError(errors.New("cannot swipe yourself"), 400, "cannot swipe yourself")
	}
	err := s.checkUserPremium(c, &userReq.User)
	if err != nil {
		return err
	}

	err = s.SwipeRepository.CreateSwipe(c, userReq.User.ID, userReq.ID, "pass")
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
	err := s.checkUserPremium(c, &userReq.User)
	if err != nil {
		return nil, err
	}

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

func (s *SwipeService) checkUserPremium(c *gin.Context, user *models.UserRequest) error {
	if user.IsPremium {
		layout := "2006-01-02T15:04:05.99999-07:00"
		dateExpired, err := time.Parse(layout, user.PremiumExpiredAt)
		if err != nil {
			return fmt.Errorf("failed to parse premium expiration date: %v", err)
		}

		if dateExpired.Before(time.Now()) {
			return s.checkSwipeLimit(c, user.ID)
		}
		return nil
	}

	return s.checkSwipeLimit(c, user.ID)
}

func (s *SwipeService) checkSwipeLimit(c *gin.Context, userID int) error {
	totalSwipe, err := s.SwipeRepository.CountUserSwipe(c, userID)
	if err != nil {
		return fmt.Errorf("failed to count user swipes: %v", err)
	}

	if totalSwipe >= 10 {
		return response.NewError(errors.New("already reach maximum swipe"), 400, "already reach maximum swipe")
	}

	return nil
}
