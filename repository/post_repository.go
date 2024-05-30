package repository

import (
	"errors"
	"fmt"
	"gin_social/dto"
	"gin_social/entities"

	"gorm.io/gorm"
)

type PostRepository interface{
	Create(post *entities.Post) error
	CountAll(params *dto.FilterParam) (int64, error)
	FindAll(params *dto.FilterParam) (*[]dto.PostResponse, error)
	Detail(id *int) (dto.PostResponse, error)
	Update(id int, post *entities.Post) error
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

func (r *postRepository) CountAll(params *dto.FilterParam) (int64, error) {
	var count int64

	query := r.db.Model(&entities.Post{})
	err := query.Count(&count)

	if params.Search != "" {
		query.Where("tweet LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	if err != nil{
		return count, err.Error
	}

	return count, nil
}

func (r *postRepository) FindAll(params *dto.FilterParam) (*[]dto.PostResponse, error) {
	var postResponse []dto.PostResponse

	query := r.db.Model(&entities.Post{}).Select("id, user_id, tweet, picture_url, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s) as updated_at")


	if params.Search != ""{
		query.Where("tweet LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	err := query.Preload("User").Offset(params.Offset).Limit(params.Limit).Find(&postResponse).Error

	return &postResponse, err
}


func (r *postRepository) Detail(id *int) (dto.PostResponse, error) {
	var postResponse dto.PostResponse

	err := r.db.Model(&entities.Post{}).Select("id, user_id, tweet, picture_url, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s) as updated_at").Preload("User").First(&postResponse, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("tweet not found")
	}

	return postResponse, err
}

func (r *postRepository) Update(id int, post *entities.Post) error{
	err := r.db.Model(&post).Where("id = ?", id).Updates(&post).Error

	return err
}


