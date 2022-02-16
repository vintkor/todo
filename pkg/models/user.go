package models

type User struct {
	Id       uint32 `json:"_"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
