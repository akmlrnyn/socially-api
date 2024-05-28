package services

import (
	"gin_social/dto"
	"gin_social/entities"
	"gin_social/errorhandler"
	"gin_social/repository"
)

type PostService interface{
	Create(post *dto.PostRequest) error
}

type postService struct{
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService{
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *dto.PostRequest) error{	
	post := entities.Post{
		UserID: req.UserId,
		Tweet: req.Tweet,
	}

	if req.Picture != nil{
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil{
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}