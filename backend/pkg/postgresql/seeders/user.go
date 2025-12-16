package seeders

import (
	"edukarsa-backend/internal/domain/models"
	"edukarsa-backend/internal/utils"

	faker "github.com/go-faker/faker/v4"
)

func (s Seed) UserSeed() error {
	for i := 0; i < 100; i++ {
		hashPassword, err := utils.HashPasswordBcrypt("user12345")
		if err != nil {
			return err
		}

		user := models.User{
			RoleID:   1,
			Name:     faker.Name(),
			Email:    faker.Email(),
			Username: faker.Username(),
			Password: hashPassword,
		}

		err = s.DB.Create(&user).Error
		if err != nil {
			return err
		}

	}
	return nil
}
