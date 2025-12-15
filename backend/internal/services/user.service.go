package services

import "edukarsa-backend/internal/repositories"

type UserService interface {
}

type userServiceImpl struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userServiceImpl{repo: repo}
}
