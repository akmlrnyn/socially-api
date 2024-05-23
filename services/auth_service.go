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
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
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

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)

	if err != nil{
		return nil, &errorhandler.ErrrorNotFound{Message: err.Error()}
	}

	if err := helpers.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.ErrrorNotFound{Message: "Password not on the records"}
	}

	token, err := helpers.GenerateToken(user)

	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID: user.ID,
		Name: user.Name,
		Token: token,
	}

	return &data, nil
}
