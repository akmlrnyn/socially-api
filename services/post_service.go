package services

import (
	"gin_social/dto"
	"gin_social/entities"
	"gin_social/errorhandler"
	"gin_social/repository"
	"math"
)

type PostService interface{
	Create(post *dto.PostRequest) error
	FindAll(param *dto.FilterParam) (*[]dto.PostResponse, *dto.Paginate, error)
	Detail(id int) (*dto.PostResponse, error)
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

func (s *postService) FindAll(params *dto.FilterParam) (*[]dto.PostResponse, *dto.Paginate, error) {
	total, err := s.repository.CountAll(params)

	if err != nil{
		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	posts, err := s.repository.FindAll(params)

	if err != nil{
		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}


	paginate := &dto.Paginate{
		Total:      int(total),
		PerPage:    params.Limit,
		Page:       params.Page,
		TotalPage: int(math.Ceil(float64(total) / float64(params.Limit))),
	}

	return posts, paginate, nil
}

func (s *postService) Detail(id int) (*dto.PostResponse, error){
	post, err := s.repository.Detail(&id)

	if err != nil{
		return nil, &errorhandler.ErrrorNotFound{Message: "tweet not found"}
	}

	return &post, nil
}