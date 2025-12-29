package services

import (
	"context"
	"edukarsa-backend/internal/domain"
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/repositories"
	"edukarsa-backend/internal/utils"
	"net/mail"

	"gorm.io/gorm"
)

type UserService interface {
	Register(ctx context.Context, reg *models.RegisterUser) error
	Login(ctx context.Context, reg *models.Login) (*models.User, error)
	FindByID(ctx context.Context, id uint64) (*models.User, error)
	UpdateUserData(ctx context.Context, id uint64, data models.UpdateUserData) error
}

type userServiceImpl struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) UpdateUserData(ctx context.Context, id uint64, data models.UpdateUserData) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if data.Name != nil {
		user.Name = *data.Name
	}

	if data.Email != nil {
		user.Email = *data.Email
	}

	if data.Username != nil {
		user.Username = *data.Username
	}

	if data.BirthDay != nil {
		user.BirthDay = data.BirthDay
	}

	return s.repo.Update(ctx, user)
}

func (s *userServiceImpl) FindByID(ctx context.Context, id uint64) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *userServiceImpl) Register(ctx context.Context, reg *models.RegisterUser) error {
	usernameExist, err := s.repo.ExistByUsername(ctx, reg.Username)
	if err != nil {
		return err
	}

	if usernameExist {
		return domain.ErrUsernameExist
	}

	emailExist, err := s.repo.ExistByEmail(ctx, reg.Email)
	if err != nil {
		return err
	}

	if emailExist {
		return domain.ErrEmailExist
	}

	user := models.User{
		Name:     reg.Name,
		Email:    reg.Email,
		Username: reg.Username,
		Password: reg.Password,
	}

	return s.repo.Create(ctx, &user)
}

func (s *userServiceImpl) Login(ctx context.Context, reg *models.Login) (*models.User, error) {
	var user *models.User
	var err error

	_, err = mail.ParseAddress(reg.Identifier)
	if err == nil {
		user, err = s.repo.FindByEmail(ctx, reg.Identifier)
	} else {
		user, err = s.repo.FindByUsername(ctx, reg.Identifier)
	}

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, gorm.ErrRecordNotFound
	}

	if !utils.ValidatePasswordBcrypt(reg.Password, user.Password) {
		return nil, domain.ErrWrongPassword
	}

	return user, nil
}
