package service

import (
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
)

func GetAllCardsUser(month, year, userID, afterID int, search string) (cards []models.CardsExpense, err error) {
	return repository.GetAllCardsUser(month, year, userID, afterID, search)
}

func GetCardExpenseByID(userID, cardExpenseID uint) (card models.CardsExpense, err error) {
	return repository.GetCardExpenseByID(userID, cardExpenseID)
}

func CreateCardExpense(expense models.CardsExpense) (err error) {
	return repository.CreateCardExpense(expense)
}

func CreateCardExpensePayer(expense models.CardsExpensePayer) (err error) {
	return repository.CreateCardExpensePayer(expense)
}

func CreateCardExpenseUser(user models.CardsExpenseUser) (err error) {
	return repository.CreateCardExpenseUser(user)
}

func UpdateCardExpense(expense models.CardsExpense) (err error) {
	return repository.UpdateCardExpense(expense)
}

func UpdateCardExpensePayer(expense models.CardsExpensePayer) (err error) {
	return repository.UpdateCardExpensePayer(expense)
}

func UpdateCardExpenseUser(user models.CardsExpenseUser) (err error) {
	return repository.UpdateCardExpenseUser(user)
}

func DeleteCardExpense(id uint) (err error) {
	return repository.DeleteCardExpense(id)
}

func DeleteCardExpensePayer(id uint) (err error) {
	return repository.DeleteCardExpensePayer(id)
}

func DeleteCardExpenseUser(id uint) (err error) {
	return repository.DeleteCardExpenseUser(id)
}
