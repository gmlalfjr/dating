package profile

import (
	"dating/domains/models"
	"dating/repository/sql/auth"
	"dating/repository/sql/profile"
	"github.com/gin-gonic/gin"
)

type IProfileService interface {
	GetListProfile(c *gin.Context, userReq *models.UserRequest, queryParam map[string]string) ([]*models.ListProfileResponse, int64, int, error)
}

type ProfileService struct {
	ProfileRepository profile.IProfileRepository
	AuthRepository    auth.IAuthRepository
}

func InitProfileServices(profileService *ProfileService) IProfileService {
	return profileService
}
