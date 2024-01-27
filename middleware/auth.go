package middleware

import (
	"dating/constants"
	"dating/response"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

type IAuthMiddleware interface {
	JWTVerifyToken(c *gin.Context)
}

type AuthMiddleware struct {
}

func InitAuthMiddleware() IAuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) JWTVerifyToken(c *gin.Context) {
	getHeader := c.GetHeader("Authorization")
	if len(getHeader) <= 0 {
		err := response.NewError(errors.New("not authorize"), 401, "not authorize")

		response := response.Response{}
		response.Error(c, err)
		c.Abort()
		return
	}
	if !strings.Contains(getHeader, "Bearer") {
		err := response.NewError(errors.New("bad request error"), 400, "bad request error")

		response := response.Response{}
		response.Error(c, err)
		c.Abort()
		return
	}
	tokenString := strings.Replace(getHeader, "Bearer ", "", -1)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTToken), nil
	})
	if err != nil {
		err := response.NewError(errors.New("Token not valid"), 400, "Token not valid")

		response := response.Response{}
		response.Error(c, err)
		c.Abort()
		return
	}

	c.Set("USER", claims)
	c.Next()
}
