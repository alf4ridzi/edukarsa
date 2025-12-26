package postgresql

import (
	"edukarsa-backend/internal/config"
	"edukarsa-backend/internal/domain/models"
	"fmt"
	"slices"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Migration = []any{
	&models.Role{},
	&models.User{},
	&models.Class{},
	&models.ClassUser{},
}

// var Migration = []any{
// 	&models.User{},
// 	&models.Class{},
// 	&models.Role{},
// }

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		config.AppConfig.DBHost, config.AppConfig.DBUser, config.AppConfig.DBPassword, config.AppConfig.DBName, config.AppConfig.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, err
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(Migration...)
}

func DropTable(db *gorm.DB) error {
	migration := Migration
	slices.Reverse(migration)
	return db.Migrator().DropTable(migration...)
}
