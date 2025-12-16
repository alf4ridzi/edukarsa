package seeders

import "gorm.io/gorm"

type Seed struct {
	DB *gorm.DB
}

func (s Seed) Run() error {
	if err := s.RoleSeed(); err != nil {
		return err
	}

	if err := s.UserSeed(); err != nil {
		return err
	}

	return nil
}
