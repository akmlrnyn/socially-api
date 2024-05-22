package services

import (
	"gin_social/dto"
	"gin_social/entities"
	"gin_social/errorhandler"
	"gin_social/helpers"
	"gin_social/repository"
)

type AuthService interface{
	Register(req *dto.RegisterRequest) error
}

type authService struct{
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "Email already exist"}
	}

	if req.Password != req.PasswordConfirm {
		return &errorhandler.BadRequestError{Message: "Password does not match"}
	}

	passwordHash, err := helpers.HashPassword(req.Password)

	if  err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := entities.User{
		Name: req.Name,
		Email: req.Email,
		Password: passwordHash,
		Gender: req.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
