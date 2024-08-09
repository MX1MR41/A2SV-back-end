package models

// User struct defines the User data model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"passwword"`
	Role     string `json:"role"`
}
