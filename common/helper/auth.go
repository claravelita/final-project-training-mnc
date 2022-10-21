package helper

import (
	"os"

	"github.com/claravelita/final-project-training-mnc/common"
	"github.com/claravelita/final-project-training-mnc/dtos"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthHelper interface {
	ParseJWT(accessToken string) (claims dtos.TokenClaims, err error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	CreateJWTTokenLogin(users, email, username string) (token string, err error)
}

type authHelper struct {
	JwtKey string
}

func NewAuthHelper(jwtKey string) (AuthHelper, error) {
	return &authHelper{JwtKey: jwtKey}, nil
}

func (a authHelper) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), common.DefaultCost)
	return string(bytes), err
}

func (a authHelper) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a authHelper) CreateJWTTokenLogin(users, email, username string) (token string, err error) {
	tokenClaims := jwt.MapClaims{}
	tokenClaims["user_id"] = users
	tokenClaims["email"] = email
	tokenClaims["username"] = username

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	token, err = claim.SignedString([]byte(os.Getenv("SECRET_AUTH")))
	if err != nil {
		return "", err
	}
	return
}

func (a authHelper) ParseJWT(accessToken string) (claims dtos.TokenClaims, err error) {
	_, err = jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_AUTH")), nil
	})
	return
}
