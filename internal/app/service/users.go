package service

import (
	"fmt"
	"strings"
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
	"technoCredits/pkg/errs"
	"technoCredits/pkg/logger"
	"technoCredits/pkg/utils"
)

func GetUserByID(userID uint) (user models.User, err error) {
	user, err = repository.GetUserByID(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

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

func UpdateUserPassword(userID uint, oldPassword, newPassword string) error {
	newPassword = strings.TrimSpace(newPassword)
	newPassword = utils.GenerateHash(newPassword)

	oldPassword = strings.TrimSpace(oldPassword)
	oldPassword = utils.GenerateHash(oldPassword)

	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Password != oldPassword {
		return errs.ErrPermissionDenied
	}

	user.Password = newPassword
	err = repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user models.User) (err error) {
	userDB, err := GetUserByID(user.ID)
	if err == nil {
		user.Password = userDB.Password

		if user.Username == userDB.Username {
			user.Username = ""
		}

		if user.Email == userDB.Email {
			user.Email = ""
		}
	}

	err = repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
