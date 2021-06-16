package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func IsAuthorized(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("token")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	postData := strings.NewReader("")

	fmt.Println("Request: ", "http://"+os.Getenv("AUTH_SERVER_HOST")+":"+os.Getenv("AUTH_SERVER_IP")+"/login")
	req, err := http.NewRequest("POST", "http://"+os.Getenv("AUTH_SERVER_HOST")+":"+os.Getenv("AUTH_SERVER_IP")+"/login", postData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	req.AddCookie(cookie)

	resp, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	fmt.Println(resp)

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}
