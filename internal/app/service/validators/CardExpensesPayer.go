package validators

import (
	"technoCredits/internal/app/models"
	"technoCredits/pkg/errs"
)

func ValidateCardExpensePayerPaidAmount(c *models.CardsExpensePayer) error {
	if c.PaidAmount == 0 {
		return errs.ErrCardExpensePayerPaidAmountNotFound
	}
	return nil
}

func ValidateCardExpensePayerExpenseID(c *models.CardsExpensePayer) error {
	if c.CardsExpenseID == 0 {
		return errs.ErrCardExpenseIDNotFound
	}
	return nil
}

func ValidateCardExpensePayerUserID(c *models.CardsExpensePayer) error {
	if c.UserID == 0 {
		return errs.ErrCardExpenseUserIDNotFound
	}
	return nil
}
