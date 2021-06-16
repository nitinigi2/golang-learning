package controller

import (
	"fmt"
	"net/http"

	"github.com/bookApiDocker/authserver/entity"
	"github.com/bookApiDocker/authserver/service"
)

func Login(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err == nil {
		if ok, _ := service.IsValidToken(cookie.Value); ok {
			fmt.Println("User already authenticated")
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	var user entity.User
	RequestBodyToObject(w, r, &user)

	token, expirationTime, err := service.IssueToken(&user, w)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

}
