package service

import (
	"fmt"
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
	"technoCredits/pkg/errs"
	"technoCredits/pkg/logger"
	"technoCredits/pkg/utils"
)

func CreateUser(user models.User) (uint, error) {
	usernameExists, emailExists, err := repository.UserExists(user.Username, user.Email)
	if err != nil {
		return 0, fmt.Errorf("failed to check existing user: %w", err)
	}

	if user.Password == "" || user.Username == "" {
		return 0, errs.ErrInvalidData
	}

	if usernameExists {
		logger.Error.Printf("[service.CreateUser] user with username %s already exists", user.Username)

		return 0, errs.ErrUsernameUniquenessFailed
	}

	if emailExists {
		logger.Error.Printf("user with email %s already exists", user.Email)

		return 0, errs.ErrEmailUniquenessFailed
	}

	user.Password = utils.GenerateHash(user.Password)

	var userDB models.User

	if userDB, err = repository.CreateUser(user); err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return userDB.ID, nil
}
