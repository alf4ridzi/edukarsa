package seeders

import (
	"edukarsa-backend/internal/domain/models"
)

type UserSeedData struct {
	RoleName string
	models.User
}

// DO NOT USE IT FOR PRODUCTION!!!
var users = []UserSeedData{
	{
		RoleName: "admin",
		User: models.User{
			Name:     "admin testing",
			Email:    "admin@gmail.com",
			Username: "admin",
			Password: "admin123",
		},
	},
	{
		RoleName: "teacher",
		User: models.User{
			Name:     "teacher testing",
			Email:    "teacher@gmail.com",
			Username: "teacher",
			Password: "teacher123",
		},
	},
	{
		RoleName: "student",
		User: models.User{
			Name:     "student testing",
			Email:    "student@gmail.com",
			Username: "student",
			Password: "student123",
		},
	},
}

func (s Seed) UserSeed() error {
	// for range 100 {
	// 	user := models.User{
	// 		RoleID:   1,
	// 		Name:     faker.Name(),
	// 		Email:    faker.Email(),
	// 		Username: faker.Username(),
	// 		Password: "user12345",
	// 	}

	// 	err := s.DB.Create(&user).Error
	// 	if err != nil {
	// 		return err
	// 	}

	for _, u := range users {
		var role models.Role
		err := s.DB.Where("name = ?", u.RoleName).First(&role).Error
		if err != nil {
			return err
		}

		user := u.User

		user.RoleID = role.ID

		err = s.DB.Where("username = ? AND email = ?", user.Username, user.Email).FirstOrCreate(&user).Error
		if err != nil {
			return err
		}
	}

	return nil
}
