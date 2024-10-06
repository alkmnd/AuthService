package models

type User struct {
	Id    string `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
}
