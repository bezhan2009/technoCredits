package seeds

import (
	"errors"
	"gorm.io/gorm"
	"technoCredits/internal/app/models"
	"technoCredits/pkg/logger"
)

func SeedRoles(db *gorm.DB) error {
	roles := []models.Role{
		{ID: 1, Name: "Admin"},
		{ID: 2, Name: "User"},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := db.First(&existingRole, "name = ?", role.Name).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&role)
			} else {
				logger.Error.Printf("[seeds.SeedRoles] Error seeding roles: %v", err)

				return err
			}
		}
	}

	return nil
}
