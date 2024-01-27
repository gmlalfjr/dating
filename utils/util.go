package utils

import (
	"dating/domains/models"
	"encoding/json"
	"github.com/devfeel/mapper"

	"github.com/gin-gonic/gin"
	"time"
)

var TimeNow = time.Now()

func MapRequest(ctx *gin.Context, request *models.BaseRequest, keys []string) *models.BaseRequest {
	valMap := make(map[string]interface{})
	for i := range keys {
		valMap[keys[i]] = ctx.Param(keys[i])
	}
	valMap["token"] = ctx.GetHeader("Authorization")

	// get query param
	queryParams := make(map[string]string)
	for k, v := range ctx.Request.URL.Query() {
		if len(v) == 1 && len(v[0]) != 0 {
			queryParams[k] = v[0]
		}
	}

	request.QueryParam = queryParams

	mapper.MapperMap(valMap, request)
	ctx.ShouldBindJSON(request.BodyData)

	if ctx.Keys["USER"] != nil {
		user := models.UserRequest{}
		AutoMap(ctx.Keys["USER"], &user)
		request.User = user
	}

	return request
}

func AutoMap(from interface{}, to interface{}) error {
	jsonFrom, _ := json.Marshal(from)
	err := json.Unmarshal([]byte(jsonFrom), to)
	return err
}

func CheckMapString(list map[string]string, lookup string) bool {
	if _, ok := list[lookup]; ok {
		return true
	}
	return false
}
