package repository

import (
	"gin_social/entities"

	"gorm.io/gorm"
)

type PostRepository interface{
	Create(post *entities.Post) error
}

type postRepository struct{
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository{
	return &postRepository{
		db: db,
	}
}


func (r *postRepository) Create(post *entities.Post) error{
	err := r.db.Create(&post).Error

	return err
}

