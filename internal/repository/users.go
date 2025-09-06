package repository

import (
	"errors"
	"gorm.io/gorm"
	"technoCredits/internal/app/models"
	"technoCredits/pkg/db"
	"technoCredits/pkg/errs"
	"technoCredits/pkg/logger"
	"time"
)

func GetUsersWithPagination(
	search string,
	afterCreatedAt *time.Time,
	afterID *uint,
	limit int,
) ([]models.User, error) {
	dbConn := db.GetDBConn()

	query := dbConn

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where(`
			users.username ILIKE ? OR 
			users.email ILIKE ? OR 
			users.first_name ILIKE ? OR 
			users.last_name ILIKE ? OR 
			roles.name ILIKE ?`,
			searchPattern, searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}

	// Cursor pagination: fetch users after given (created_at, id)
	if afterCreatedAt != nil && afterID != nil {
		query = query.Where(`
			(users.created_at > ?) OR 
			(users.created_at = ? AND users.id > ?)`,
			*afterCreatedAt, *afterCreatedAt, *afterID,
		)
	}

	// Apply ordering and limit
	var users []models.User
	err := query.
		Order("users.created_at ASC").
		Order("users.id ASC").
		Limit(limit).
		Find(&users).Error

	if err != nil {
		logger.Error.Printf("[repository.GetUsersWithPagination] error: %s", err.Error())
		return nil, TranslateGormError(err)
	}

	return users, nil
}

func GetAllUsers(search string) (users []models.User, err error) {
	query := db.GetDBConn()
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where(`
			users.username ILIKE ? OR 
			users.email ILIKE ? OR 
			users.first_name ILIKE ? OR 
			users.last_name ILIKE ? OR 
			roles.name ILIKE ?`,
			searchPattern, searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}

	err = query.Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, TranslateGormError(err)
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().
		Where("id = ?", id).
		First(&user).Error

	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return user, TranslateGormError(err)
	}
	return user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func UserExists(username, email string) (bool, bool, error) {
	users, err := GetAllUsers("")
	if err != nil {
		return false, false, err
	}

	var usernameExists, emailExists bool
	for _, user := range users {
		if user.Username == username {
			usernameExists = true
		}
		if user.Email == email {
			emailExists = true
		}
	}
	return usernameExists, emailExists, nil
}

func CreateUser(user models.User) (userDB models.User, err error) {
	tx := db.GetDBConn().Begin()

	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return userDB, TranslateGormError(err)
	}

	if err = tx.Commit().Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error committing user: %v\n", err)
		return userDB, TranslateGormError(err)
	}

	userDB = user
	return userDB, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, TranslateGormError(err)
	}

	return user, nil
}

func UpdateUser(user models.User) (err error) {
	if err = db.GetDBConn().Model(models.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		logger.Error.Printf("[repository.UpdateUserStruct] Error while updating user: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteUserByID(userID uint) (err error) {
	if err = db.GetDBConn().Model(&models.User{}).Delete(&models.User{}, userID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteUserByID] error deleting user: %v\n", err)
		return TranslateGormError(err)
	}

	return nil
}
