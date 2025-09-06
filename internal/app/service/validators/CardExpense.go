package validators

import (
	"technoCredits/internal/app/models"
	"technoCredits/pkg/errs"
)

func ValidateCardExpenseGroupID(c *models.CardsExpense) error {
	if c.GroupID == 0 {
		return errs.ErrCardExpenseGroupIDNotFound
	}
	return nil
}

func ValidateCardDescription(c *models.CardsExpense) error {
	if c.Description == "" {
		return errs.ErrCardExpenseDescriptionNotFound
	}
	return nil
}

func ValidateCardTotalAmount(c *models.CardsExpense) error {
	if c.TotalAmount == 0 {
		return errs.ErrCardExpenseTotalAmountNotFound
	}
	return nil
}
