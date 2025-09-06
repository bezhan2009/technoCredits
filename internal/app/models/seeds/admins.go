package seeds

import (
	"errors"
	"gorm.io/gorm"
	"technoCredits/internal/app/models"
	"technoCredits/pkg/logger"
)

func SeedAdmins(db *gorm.DB) error {
	userAdmins := []models.User{
		{Username: "tamaev", Password: "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3", Email: "karimovbezan0@gmail.com", RoleID: 1},
	}

	for _, userAdmin := range userAdmins {
		var existingUserAdmin models.User
		if err := db.First(&existingUserAdmin, "username = ?", userAdmin.Username).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&userAdmin)
			} else {
				logger.Error.Printf("[seeds.SeedAdmins] Error seeding userAdmins: %v", err)

				return err
			}
		}
	}

	return nil
}
