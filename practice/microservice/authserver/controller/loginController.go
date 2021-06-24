package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bookApiDocker/authserver/entity"
	"github.com/bookApiDocker/authserver/service"
)

func Login(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err == nil {
		if ok, userRole, _ := service.IsValidToken(cookie.Value); ok {
			fmt.Println("User already authenticated")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(userRole))
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

	log.Println("User logged in successfully....")
}
