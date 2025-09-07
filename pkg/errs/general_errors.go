package errs

import "errors"

// General Errors
var (
	ErrRecordNotFound        = errors.New("ErrRecordNotFound")
	ErrEmptyMessageText      = errors.New("ErrEmptyMessageText")
	ErrNoAIRecommends        = errors.New("ErrNoAIRecommends")
	ErrGroupNotFound         = errors.New("ErrGroupNotFound")
	ErrAlreadyExists         = errors.New("ErrAlreadyExists")
	ErrRoleIsRequired        = errors.New("ErrRoleIsRequired")
	ErrSomethingWentWrong    = errors.New("ErrSomethingWentWrong")
	ErrGeminiIsNotWorking    = errors.New("ErrGeminiIsNotWorking")
	ErrDeleteFailed          = errors.New("ErrDeleteFailed")
	ErrInvalidAfterID        = errors.New("ErrInvalidAfterID")
	ErrInvalidMonth          = errors.New("ErrInvalidMonth")
	ErrInvalidYear           = errors.New("ErrInvalidYear")
	ErrNoCourseFound         = errors.New("ErrNoCourseFound")
	ErrinsufficientFunds     = errors.New("ErrinsufficientFunds")
	ErrNoVacancyFound        = errors.New("ErrNoVacancyFound")
	ErrNoUsersStatisticFound = errors.New("ErrNoUsersStatisticFound")
)
