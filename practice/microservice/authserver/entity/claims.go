package entity

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	UserRole string `json:"role"`
	jwt.StandardClaims
}
