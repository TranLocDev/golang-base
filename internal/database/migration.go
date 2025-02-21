package database

import (
	"my-go-project/internal/models"

	"gorm.io/gorm"
)

// Migrate thực hiện auto-migration cho các model
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Countries{},
		&models.Students{},
	)
}
