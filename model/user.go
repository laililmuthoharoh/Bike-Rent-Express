package model

import "github.com/dgrijalva/jwt-go"

type (
	JWTClaim struct {
		jwt.StandardClaims
		Username string `json:"username"`
		Roles    string `json:"role"`
	}

	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	User struct {
		Username string
		Password string
		Roles    string
	}
)
