package entity

import "fmt"

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Roles    []string
}

func (user User) String() string {
	return fmt.Sprintf("{Username:%s, Password:%s}", user.UserName, user.Password)
}