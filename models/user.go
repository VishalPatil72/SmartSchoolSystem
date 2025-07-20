package models

type UserLogin struct {
	UserId   uint   `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // e.g., principal, teacher, parent
}
