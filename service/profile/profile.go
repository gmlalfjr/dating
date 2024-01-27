package profile

import (
	"dating/domains/models"
	"dating/response"
	"dating/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (p *ProfileService) GetListProfile(c *gin.Context, userReq *models.UserRequest, queryParam map[string]string) ([]*models.ListProfileResponse, int64, int, error) {
	var (
		limit  = 10
		offset = 0
		result = []*models.ListProfileResponse{}
	)
	if utils.CheckMapString(queryParam, "limit") {
		limNum := queryParam["limit"]
		num, err := strconv.Atoi(limNum)
		if err != nil {
			return nil, 0, 0, response.NewError(err, 400, err.Error())
		}
		limit = num
	}

	if utils.CheckMapString(queryParam, "offset") {
		offsetNum := queryParam["offset"]
		num, err := strconv.Atoi(offsetNum)
		if err != nil {
			return nil, 0, 0, response.NewError(err, 400, err.Error())
		}
		offset = num
	}

	prf, totalData, err := p.ProfileRepository.FindProfileExcludeUserId(userReq.ID, limit, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	for _, user := range prf {
		result = append(result, &models.ListProfileResponse{
			Id:       user.ID,
			FullName: user.FullName,
			Address:  user.Address,
			Age:      user.Age,
			Sex:      user.Sex,
			Email:    user.Email,
		})
	}

	totalPage := (int(totalData) - offset + limit - 1) / limit

	return result, totalData, totalPage, nil
}
