package models

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewUserTemplate struct {
	Username string
	Password string
}

type SessionResponse struct {
	Username string `json:"username"`
	Exp      int    `json:"exp"`
}
