package repository

import (
	"technoCredits/internal/app/models"
	"technoCredits/pkg/db"
	"technoCredits/pkg/errs"
	"technoCredits/pkg/logger"
)

func GetAllCardsUser(month, year, userID, afterID int, search string, groupIDFilter int, userIDFilter int) (cards []models.CardsExpense, err error) {
	dbConn := db.GetDBConn().Model(&models.CardsExpense{}).
		Preload("CardsExpensePayers").
		Preload("CardsExpenseUsers").
		Preload("CardsExpensePayers.User").
		Preload("CardsExpenseUsers.User").
		Preload("Group").
		Preload("Group.Owner").
		Joins("JOIN group_members ON group_members.group_id = cards_expenses.group_id").
		Where("group_members.user_id = ?", userID).
		Where("EXTRACT(MONTH FROM cards_expenses.created_at) = ? AND EXTRACT(YEAR FROM cards_expenses.created_at) = ?", month, year).
		Where("cards_expenses.id > ?", afterID)

	if groupIDFilter > 0 {
		dbConn = dbConn.Where("cards_expenses.group_id = ?", groupIDFilter)
	}

	if userIDFilter > 0 {
		dbConn = dbConn.Joins(`
			LEFT JOIN cards_expense_payers ON cards_expense_payers.cards_expense_id = cards_expenses.id
		`).Joins(`
			LEFT JOIN cards_expense_users ON cards_expense_users.cards_expense_id = cards_expenses.id
		`).Where(`
			cards_expense_payers.user_id = ? OR cards_expense_users.user_id = ?
		`, userIDFilter, userIDFilter)
	}

	if search != "" {
		likeStr := "%" + search + "%"
		dbConn = dbConn.Joins(`
			LEFT JOIN cards_expense_payers ON cards_expense_payers.cards_expense_id = cards_expenses.id
		`).Joins(`
			LEFT JOIN users AS payer_users ON payer_users.id = cards_expense_payers.user_id
		`).Joins(`
			LEFT JOIN cards_expense_users ON cards_expense_users.cards_expense_id = cards_expenses.id
		`).Joins(`
			LEFT JOIN users AS expense_users ON expense_users.id = cards_expense_users.user_id
		`).Where(`
			cards_expenses.description ILIKE ? OR
			payer_users.username ILIKE ? OR
			expense_users.username ILIKE ?
		`, likeStr, likeStr, likeStr)
	}

	if err = dbConn.Order("cards_expenses.id ASC").Find(&cards).Error; err != nil {
		logger.Error.Printf("[repository.GetAllCardsUser] Error while getting all cards users: %v", err)
		return nil, TranslateGormError(err)
	}

	return cards, nil
}

func GetCardExpenseByID(userID, cardExpenseID uint) (card models.CardsExpense, err error) {
	if err = db.GetDBConn().
		Where("cards_expenses.id = ?", cardExpenseID).
		First(&card).Error; err != nil {
		logger.Error.Printf("[repository.GetCardExpenseByID] Error while getting card by id %v: %v", cardExpenseID, err)
		return models.CardsExpense{}, TranslateGormError(err)
	}

	return card, nil
}

func GetCardExpenseByIDOnly(cardExpenseID uint) (card models.CardsExpense, err error) {
	if err = db.GetDBConn().Where("id = ?", cardExpenseID).First(&card).Error; err != nil {
		logger.Error.Printf("[repository.GetCardExpenseByIDOnly] Error while getting card by id %v: %v", cardExpenseID, err)

		return models.CardsExpense{}, TranslateGormError(err)
	}

	return card, nil
}

func CheckCardAmountLimit(userID uint, cardExpenseID uint, paying float64) (models.CardsExpense, error) {
	card, err := GetCardExpenseByID(userID, cardExpenseID)
	if err != nil {
		return models.CardsExpense{}, TranslateGormError(err)
	}

	var payers []models.CardsExpensePayer
	if err = db.GetDBConn().Model(&models.CardsExpensePayer{}).Where("id = ?", cardExpenseID).First(&payers).Error; err != nil {
		logger.Error.Printf("[CheckCardAmountLimit] Error while getting card by id %v: %v", cardExpenseID, err)
		return models.CardsExpense{}, err
	}

	resAmount := 0.0
	for _, payer := range payers {
		resAmount += payer.PaidAmount
	}

	if card.TotalAmount < resAmount+paying {
		return models.CardsExpense{}, errs.ErrInsufficientFunds
	}

	return card, nil
}

func CreateCardExpense(expense models.CardsExpense) (err error) {
	_, err = GetAllUserGroupsByIDOnly(expense.GroupID)
	if err != nil {
		return errs.ErrGroupNotFound
	}

	if err = db.GetDBConn().Create(&expense).Error; err != nil {
		logger.Error.Printf("[repository.CreateCardExpense] Error while creating card expense: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func CreateCardExpensePayer(expense models.CardsExpensePayer) (err error) {
	_, err = CheckCardAmountLimit(expense.UserID, expense.CardsExpenseID, expense.PaidAmount)
	if err != nil {
		return TranslateGormError(err)
	}

	if err = db.GetDBConn().Create(&expense).Error; err != nil {
		logger.Error.Printf("[repository.CreateCardExpensePayer] Error while creating card expense: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func CreateCardExpenseUser(user models.CardsExpenseUser) (err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateCardExpenseUser] Error while creating card expense user: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateCardExpense(expense models.CardsExpense) (err error) {
	if err = db.GetDBConn().Where("id = ?", expense.ID).Updates(&expense).Error; err != nil {
		logger.Error.Printf("[repository.UpdateCardExpense] Error while updating card expense: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateCardExpensePayer(expense models.CardsExpensePayer) (err error) {
	if err = db.GetDBConn().Where("id = ?", expense.ID).Updates(&expense).Error; err != nil {
		logger.Error.Printf("[repository.UpdateCardExpensePayer] Error while updating card expense user: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateCardExpenseUser(user models.CardsExpenseUser) (err error) {
	if err = db.GetDBConn().Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		logger.Error.Printf("[repository.UpdateCardExpenseUser] Error while updating card expense user: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteCardExpense(id uint) (err error) {
	if err = db.GetDBConn().Where("id = ?", id).Delete(&models.CardsExpense{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteCardExpense] Error while deleting card expense: %v", err)
		return TranslateGormError(err)
	}
	return nil
}

func DeleteCardExpensePayer(id uint) (err error) {
	if err = db.GetDBConn().Where("id = ?", id).Delete(&models.CardsExpensePayer{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteCardExpensePayer] Error while deleting card expense payer: %v", err)
		return TranslateGormError(err)
	}
	return nil
}

func DeleteCardExpenseUser(id uint) (err error) {
	if err = db.GetDBConn().Where("id = ?", id).Delete(&models.CardsExpenseUser{}).Error; err != nil {
		logger.Error.Printf("[repository.DeleteCardExpenseUser] Error while deleting card expense user: %v", err)
		return TranslateGormError(err)
	}
	return nil
}
