package auth

import (
	"dating/constants"
	"dating/domains/models"
	"dating/response"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"

	"gorm.io/gorm"
)

func (a *AuthService) Register(c *gin.Context, request *models.RegisterRequest) (*models.RegisterResponse, error) {
	pass, err := a.hashingPassword(c, request.Password, request.PasswordRepeat)
	if err != nil {
		return nil, err
	}

	res, err := a.AuthRepository.FindOneUserByEmail(c, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if res != nil {
		return nil, response.NewError(errors.New("username already exist"), 400, "username already exist")
	}
	request.Password = pass
	res, err = a.AuthRepository.CreateUser(c, request)
	if err != nil {
		return nil, err
	}
	return &models.RegisterResponse{Email: res.Email}, nil

}

func (a *AuthService) Login(c *gin.Context, request *models.LoginRequest) (*models.LoginResponse, error) {
	var isPremium = false
	res, err := a.PremiumUserRepo.FindByEmail(c, request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewError(err, 404, "Wrong email or password")
		}
		return nil, err
	}
	err = a.checkPasswordHash(c, res.Password, request.Password)
	if err != nil {
		return nil, err
	}

	genTokenData := jwt.MapClaims{
		"id":         res.ID,
		"email":      res.Email,
		"is_premium": false,
		"exp":        time.Now().Add(time.Minute * 1200).Unix(),
	}
	if !res.ExpiredAt.IsZero() && res.ExpiredAt.After(time.Now()) {
		isPremium = true
		genTokenData["is_premium"] = isPremium
		genTokenData["premium_expired_at"] = res.ExpiredAt
	}
	token, err := a.generateToken(genTokenData)
	return &models.LoginResponse{
		Email:        res.Email,
		Token:        token.Token,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (a *AuthService) hashingPassword(c *gin.Context, passwordOne, passwordTwo string) (string, error) {

	if passwordOne != passwordTwo {
		return "", response.NewError(errors.New("password not match"), 400, "password not match")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOne), 14)
	if err != nil {
		return "", response.NewError(err, 400, "failed register")

	}
	return string(hash), nil
}

func (a *AuthService) generateToken(data jwt.MapClaims) (*models.TokenLogin, error) {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	data["exp"] = time.Now().Add(time.Minute * 2400).Unix()
	generateRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, err := generateToken.SignedString([]byte(constants.JWTToken))
	refreshTokenString, errRefresh := generateRefreshToken.SignedString([]byte(constants.RefreshJWTToken))
	if err != nil || errRefresh != nil {
		return nil, response.NewError(errors.New("Failed Generated Token"), 400, "Failed Generated Token")
	}
	return &models.TokenLogin{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (a *AuthService) checkPasswordHash(c *gin.Context, hash string, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return response.NewError(err, 400, "Wrong email or password")
	}
	return nil
}
