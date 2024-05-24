package dto

import (
	"mime/multipart"
	"time"
)


type PostResponse struct{
	ID int `json:"id"`
	UserId int `json:"user_id"`
	User User `gorm:foreignKey:"UserId" json:"User"`
	Tweet string `json:"tweet"`
	PictureUrl string `json:"picture_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRequest struct{
	UserId int `form:"user_id"`
	Tweet string `form:"tweet"`
	Picture *multipart.FileHeader `form:"picture"`
}

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}