package models

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserChangePassword struct {
	Password string `json:"password"`
}

func (u *User) ValidateNewPassword(p *UserChangePassword) error {
	return validation.ValidateStruct(p, validation.Field(&p.Password, UserRules["Password"]...))
	/*return validation.Validate(fieldsMap, validation.Map(
		validation.Key("password", UserRules["Password"]...),
	))*/
}

type ProfileChangePassword struct {
	UserChangePassword
	CurrentPassword string `json:"currentPassword"`
}

func (u *User) ValidateCurrentPassword(p *ProfileChangePassword) error {
	if err := validation.ValidateStruct(p,
		validation.Field(&p.CurrentPassword, UserRules["Password"]...),
		validation.Field(&p.Password, UserRules["Password"]...),
	); err != nil {
		return err
	}

	if err := u.ValidatePassword(p.CurrentPassword); err != nil {
		return validation.Errors{"currentPassword": errors.New("Old password is invalid")}
	}

	return nil
}
