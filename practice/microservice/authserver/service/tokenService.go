package service

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/bookApiDocker/authserver/entity"
	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(os.Getenv("JWT_SIGNING"))

func GenerateJWT(user *entity.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &entity.Claims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatal("something Went Wrong: %s", err.Error())
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func IsValidToken(token string) (bool, error) {
	claims := &entity.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error")
		}
		return mySigningKey, nil
	})

	return tkn.Valid, err
}
