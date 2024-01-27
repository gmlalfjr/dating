package swipe

import (
	"dating/domains/entities"
	"github.com/gin-gonic/gin"
)

func (s *SwipeRepository) CreateSwipe(c *gin.Context, userId int, swipedUserId int, action string) error {
	response := &entities.Swipe{
		UserID:       userId,
		SwipedUserID: swipedUserId,
		Action:       action,
	}
	if err := s.db.Create(response).Error; err != nil {
		return err
	}

	return nil

}

func (s *SwipeRepository) FindOneByUserIdAndSwipedUserId(c *gin.Context, userId int, swipedUserId int) (*entities.Swipe, error) {
	response := &entities.Swipe{}
	if err := s.db.Where("user_id = ? and swiped_user_id =?", swipedUserId, userId).First(response).Error; err != nil {
		return nil, err
	}

	return response, nil
}

func (s *SwipeRepository) CountUserSwipe(c *gin.Context, userId int) (int64, error) {

	var totalSwipes int64
	if err := s.db.Debug().Model(&entities.Swipe{}).Where("user_id = ? AND DATE(created_at) = CURRENT_DATE", userId).Count(&totalSwipes).Error; err != nil {
		return 0, err
	}

	return totalSwipes, nil

}
