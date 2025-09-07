package service

import (
	"fmt"
	"technoCredits/internal/app/models"
	"technoCredits/internal/repository"
	"technoCredits/pkg/brokers"
	"time"
)

func GetAllCardsUser(month, year, userID, afterID int, search string, groupIDFilter int, userIDFilter int) (cards []models.CardsExpense, err error) {
	return repository.GetAllCardsUser(month, year, userID, afterID, search, groupIDFilter, userIDFilter)
}

func GetCardExpenseByID(userID, cardExpenseID uint) (card models.CardsExpense, err error) {
	return repository.GetCardExpenseByID(userID, cardExpenseID)
}

func GetAllCardsUserByID(cardID uint) (cards []models.CardsExpense, err error) {
	return repository.GetAllCardsUserByID(cardID)
}

func CreateCardExpense(expense models.CardsExpense) (err error) {
	return repository.CreateCardExpense(expense)
}

func CreateCardExpensePayer(expense models.CardsExpensePayer) (err error) {
	return repository.CreateCardExpensePayer(expense)
}

func CreateCardExpenseUser(cardUser models.CardsExpenseUser, userID uint) (err error) {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	card, err := repository.GetCardExpenseByIDOnly(cardUser.CardsExpenseID)
	if err != nil {
		return err
	}

	group, err := repository.GetAllUserGroupsByIDOnly(card.GroupID)
	if err != nil {
		return err
	}

	queueName := fmt.Sprintf("user_%d_queue", group.OwnerID)

	err = brokers.SendMessageToQueue(queueName,
		fmt.Sprintf("Пользователь %s оплатил за карту %s: %v %s",
			user.Username, card.Description, cardUser.PaidAmount, card.Currency,
		),
	)
	if err != nil {
		return err
	}

	cardPayer, err := repository.GetCardExpensePayerByUserIDAndCardID(card.ID, userID)
	if err != nil {
		return err
	}

	cardUser.ShareAmount = cardPayer.PaidAmount - cardUser.PaidAmount

	cardUser.PaidAt = time.Now()
	//
	//cardUsers, err := repository.GetCardExpenseUsersByCardExpenseID(card.ID)
	//if err != nil {
	//	return err
	//}

	err = repository.CreateCardExpenseUser(cardUser)
	if err != nil {
		return err
	}

	err = repository.SettlementCreate(
		&models.Settlement{
			GroupID:    card.GroupID,
			FromUserID: userID,
			ToUserID:   group.OwnerID,
			Amount:     cardUser.PaidAmount,
			Currency:   "TJS",
			Note:       "",
		},
	)
	if err != nil {
		return err
	}

	return nil
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
