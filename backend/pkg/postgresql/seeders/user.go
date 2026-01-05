package seeders

import "edukarsa-backend/internal/domain/models"

// DO NOT USE IT FOR PRODUCTION!!!
var users = []models.User{
	{
		RoleID:   1,
		Name:     "admin testing",
		Email:    "admin@gmail.com",
		Username: "admin",
		Password: "admin123",
	},
	{
		RoleID:   2,
		Name:     "teacher testing",
		Email:    "teacher@gmail.com",
		Username: "teacher",
		Password: "teacher123",
	},
	{
		RoleID:   3,
		Name:     "user testing",
		Email:    "user@gmail.com",
		Username: "user",
		Password: "user123",
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

	for _, user := range users {
		err := s.DB.FirstOrCreate(&user).Error
		if err != nil {
			return err
		}
	}

	return nil
}
