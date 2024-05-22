package entities

import "time"

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Gender string  `json:"gender"`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}