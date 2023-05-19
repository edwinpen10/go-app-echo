package models

type User struct {
	ID        int    `json:"id"`
	User_uuid string `json:"user_uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
