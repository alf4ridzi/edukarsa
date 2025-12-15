package services

import (
	"context"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/utils"
)

type UserService interface {
	Register(ctx context.Context, reg *models.RegisterUser) error
}

type userServiceImpl struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) Register(ctx context.Context, reg *models.RegisterUser) error {

	usernameExist, err := s.repo.ExistByUsername(ctx, reg.Username)
	if err != nil {
		return err
	}

	if usernameExist {
		return models.ErrUsernameExist
	}

	emailExist, err := s.repo.ExistByEmail(ctx, reg.Email)
	if err != nil {
		return err
	}

	if emailExist {
		return models.ErrEmailExist
	}

	user := models.User{
		Name:     reg.Name,
		Email:    reg.Email,
		Username: reg.Username,
	}

	hashPass, err := utils.HashPasswordBcrypt(reg.Password)
	if err != nil {
		return err
	}

	user.Password = hashPass

	return s.repo.Create(ctx, user)
}
