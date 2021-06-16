package controller

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
		MaxAge:  -1,
	})
}
