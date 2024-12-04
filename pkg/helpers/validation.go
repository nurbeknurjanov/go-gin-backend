package helpers

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	ErrExistUserEmail = errors.New("User with this email already exists")
)

func RequiredIf(condition bool) validation.RuleFunc {
	return func(value interface{}) error {
		if condition {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}

func NumberRule() validation.RuleFunc {
	return func(value any) error {
		switch value.(type) {
		case int:
			return nil
		case *int:
			return nil
		case float64:
			return nil
		default:
			return errors.New("Must be a number")
		}
	}
}
