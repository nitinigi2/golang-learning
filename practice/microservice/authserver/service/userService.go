package service

import (
	"net/http"
	"time"

	"github.com/bookApiDocker/authserver/entity"
	"github.com/bookApiDocker/authserver/repository"
)

func IssueToken(user *entity.User, w http.ResponseWriter) (string, time.Time, error) {
	// if user is valid i.e present in db
	user, err := repository.IsUserValid(user)
	if err != nil {
		return "", time.Now(), err
	}
	// issue a token
	return GenerateJWT(user)
}
