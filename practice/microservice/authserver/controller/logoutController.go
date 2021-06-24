package controller

import (
	"log"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token")

	if err != nil {
		log.Println("User is not logged in")
		w.WriteHeader(http.StatusOK)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
		MaxAge:  -1,
	})

	log.Println("User logged out successfully.....")
}
