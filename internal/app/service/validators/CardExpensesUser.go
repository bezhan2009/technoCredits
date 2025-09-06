package validators

import (
	"technoCredits/internal/app/models"
	"technoCredits/pkg/errs"
)

func ValidateCardExpenseUserExpenseID(c *models.CardsExpenseUser) error {
	if c.CardsExpenseID == 0 {
		return errs.ErrCardExpenseIDNotFound
	}
	return nil
}

func ValidateCardExpenseUserUserID(c *models.CardsExpenseUser) error {
	if c.UserID == 0 {
		return errs.ErrCardExpenseUserIDNotFound
	}
	return nil
}
