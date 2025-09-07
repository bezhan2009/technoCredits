package service

import (
	"fmt"
	"math"
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

	cardUsers, err := repository.GetCardExpenseUsersByCardExpenseID(card.ID)
	if err != nil {
		return err
	}

	group, err := repository.GetAllUserGroupsByIDOnly(card.GroupID)
	if err != nil {
		return err
	}

	res := 0.0
	for _, cardUserEx := range cardUsers {
		res += cardUserEx.PaidAmount
	}

	fmt.Println(res + cardUser.PaidAmount)

	if res+cardUser.PaidAmount == card.TotalAmount || res+cardUser.PaidAmount < card.TotalAmount {
		cardPayer, err := repository.GetCardExpensePayerByUserIDAndCardID(card.ID, userID)
		if err != nil {
			return err
		}

		cardUser.ShareAmount = cardPayer.PaidAmount - cardUser.PaidAmount

		cardUser.PaidAt = time.Now()

		err = repository.CreateCardExpenseUser(cardUser)
		if err != nil {
			return err
		}

		err = repository.SettlementCreate(
			&models.Settlement{
				GroupID:    card.GroupID,
				FromUserID: userID,
				ToUserID:   group.OwnerID,
				Amount:     math.Abs(cardUser.PaidAmount),
				Currency:   "TJS",
				Note:       "",
			},
		)
		if err != nil {
			return err
		}

		queueName := fmt.Sprintf("user_%d_queue", group.OwnerID)

		err = brokers.SendMessageToQueue(queueName,
			fmt.Sprintf("Пользователь %s оплатил за карту %s: %v %s        %d",
				user.Username, card.Description, cardUser.PaidAmount, card.Currency, userID,
			),
		)
		if err != nil {
			return err
		}
	}

	if res+cardUser.PaidAmount > card.TotalAmount {
		dutyUsers := []models.CardsExpenseUser{}
		sumOfDuty := 0.0
		for _, cardUserEx := range cardUsers {
			if cardUserEx.ShareAmount < 0 {
				dutyUsers = append(dutyUsers, cardUserEx)
				sumOfDuty += cardUserEx.ShareAmount
			}
		}

		for _, dutyUser := range dutyUsers {
			queueName := fmt.Sprintf("user_%d_queue", dutyUser.UserID)

			cardUser.PaidAmount += dutyUser.ShareAmount

			err = brokers.SendMessageToQueue(queueName,
				fmt.Sprintf("Пользователь %s оплатил за долг карты %s: %v %s        %d",
					user.Username, card.Description, cardUser.PaidAmount, card.Currency, userID,
				),
			)
			if err != nil {
				return err
			}

			err = repository.SettlementCreate(
				&models.Settlement{
					GroupID:    group.ID,
					FromUserID: userID,
					ToUserID:   dutyUser.UserID,
					Amount:     math.Abs(cardUser.PaidAmount),
					Currency:   "TJS",
					Note:       fmt.Sprintf("Оплачен долг за карту %s", card.Description),
				},
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func UpdateCardExpense(expense models.CardsExpense, userID uint) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	card, err := repository.GetCardExpenseByIDOnly(expense.ID)
	if err != nil {
		return err
	}

	cardUsers, err := repository.GetCardExpenseUsersByCardExpenseID(card.ID)
	if err != nil {
		return err
	}

	group, err := repository.GetAllUserGroupsByIDOnly(card.GroupID)
	if err != nil {
		return err
	}

	err = repository.UpdateCardExpense(expense)
	if err != nil {
		return err
	}

	totalPaid := 0.0
	for _, cu := range cardUsers {
		totalPaid += cu.PaidAmount
	}

	if totalPaid == expense.TotalAmount || totalPaid < expense.TotalAmount {
		cardPayer, err := repository.GetCardExpensePayerByUserIDAndCardID(card.ID, userID)
		if err != nil {
			return err
		}

		share := cardPayer.PaidAmount - expense.TotalAmount

		err = repository.SettlementCreate(&models.Settlement{
			GroupID:    card.GroupID,
			FromUserID: userID,
			ToUserID:   group.OwnerID,
			Amount:     math.Abs(share),
			Currency:   expense.Currency,
			Note:       fmt.Sprintf("Обновлен расход по карте %s", expense.Description),
		})
		if err != nil {
			return err
		}

		queueName := fmt.Sprintf("user_%d_queue", group.OwnerID)
		err = brokers.SendMessageToQueue(queueName,
			fmt.Sprintf("Пользователь %s обновил карту %s: %v %s",
				user.Username, expense.Description, expense.TotalAmount, expense.Currency,
			),
		)
		if err != nil {
			return err
		}
	}

	if totalPaid > expense.TotalAmount {
		dutyUsers := []models.CardsExpenseUser{}
		for _, cu := range cardUsers {
			if cu.ShareAmount < 0 {
				dutyUsers = append(dutyUsers, cu)
			}
		}

		for _, dutyUser := range dutyUsers {
			queueName := fmt.Sprintf("user_%d_queue", dutyUser.UserID)
			err = brokers.SendMessageToQueue(queueName,
				fmt.Sprintf("Пользователь %s обновил долг по карте %s",
					user.Username, expense.Description,
				),
			)
			if err != nil {
				return err
			}

			err = repository.SettlementCreate(&models.Settlement{
				GroupID:    group.ID,
				FromUserID: userID,
				ToUserID:   dutyUser.UserID,
				Amount:     math.Abs(dutyUser.ShareAmount),
				Currency:   expense.Currency,
				Note:       fmt.Sprintf("Обновлен долг по карте %s", expense.Description),
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
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
