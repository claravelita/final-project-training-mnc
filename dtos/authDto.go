package dtos

import "github.com/golang-jwt/jwt"

type (
	authPackage struct {
		JwtKey string
	}

	TokenClaims struct {
		UserID   string `json:"user_id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		jwt.StandardClaims
	}

	TokenResponse struct {
		Token string `json:"token"`
	}
)
