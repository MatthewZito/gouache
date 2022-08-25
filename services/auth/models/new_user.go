package models

// NewUserModel represents the necessary input data a client must provide to create a new `User`.
type NewUserModel struct {
	Username string
	Password string
}
