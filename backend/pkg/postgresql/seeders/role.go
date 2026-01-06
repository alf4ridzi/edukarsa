package seeders

import "edukarsa-backend/internal/domain/models"

var roles = []models.Role{
	{
		Name: "admin",
	},
	{
		Name: "teacher",
	},
	{
		Name: "student",
	},
}

func (s Seed) RoleSeed() error {
	for _, r := range roles {
		role := r
		err := s.DB.FirstOrCreate(&role).Error
		if err != nil {
			return err
		}
	}

	return nil
}
