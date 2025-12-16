package seeders

import "edukarsa-backend/internal/domain/models"

var roles = []models.Role{
	{
		Name: "teacher",
	},
	{
		Name: "student",
	},
}

func (s Seed) RoleSeed() error {
	for _, role := range roles {
		err := s.DB.Create(&role).Error
		if err != nil {
			return err
		}
	}

	return nil
}
