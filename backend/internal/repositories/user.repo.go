package repositories

import "gorm.io/gorm"

type UserRepo interface {
}

type userRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepoImpl{DB: db}
}
