package db

import (
	"errors"
	"technoCredits/internal/app/models"
	"technoCredits/internal/app/models/seeds"
	"technoCredits/pkg/logger"
)

func Migrate() error {
	if dbConn == nil {
		logger.Error.Printf("[db.Migrate] Error because database connection is nil")

		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(
		models.User{},
		models.Role{},
		models.Group{},
		models.GroupMember{},
		models.CardsExpense{},
		models.CardsExpensePayer{},
		models.CardsExpenseUser{},
		models.Settlement{},
	)

	if err != nil {
		logger.Error.Printf("[db.Migrate] Error migrating tables: %v", err)

		return err
	}

	err = seeds.SeedRoles(dbConn)
	if err != nil {
		logger.Error.Printf("[db.Migrate] Error seeding roles: %v", err)

		return err
	}

	return nil
}
