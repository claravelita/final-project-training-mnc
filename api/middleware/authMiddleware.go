package middleware

import (
	"errors"
	"os"

	"github.com/claravelita/final-project-training-mnc/common/helper"
	"github.com/labstack/echo/v4"
)

func CheckJWTAndCheckAuthUser(token string, c echo.Context) (valid bool, err error) {
	valid = true
	authLibrary, errInit := helper.NewAuthHelper(os.Getenv("SECRET_AUTH"))
	if errInit != nil {
		return false, errInit
	}

	tokenClaims, errParse := authLibrary.ParseJWT(token)
	if errParse != nil {
		return false, errParse
	}
	if tokenClaims.UserID == "" {
		return false, errors.New("invalid user")
	}

	c.Set("user_id", tokenClaims.UserID)
	c.Set("username", tokenClaims.Username)
	c.Set("email", tokenClaims.Email)
	return
}
