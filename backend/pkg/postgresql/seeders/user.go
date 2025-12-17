package seeders

import (
	"edukarsa-backend/internal/domain/models"

	faker "github.com/go-faker/faker/v4"
)

func (s Seed) UserSeed() error {
	for i := 0; i < 100; i++ {

		user := models.User{
			RoleID:   1,
			Name:     faker.Name(),
			Email:    faker.Email(),
			Username: faker.Username(),
			Password: "user12345",
		}

		err := s.DB.Create(&user).Error
		if err != nil {
			return err
		}

	}
	return nil
}
