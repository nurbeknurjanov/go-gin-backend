package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
)

type SexType int8

const (
	_          SexType = iota
	SEX_MALE   SexType = iota
	SEX_FEMALE SexType = iota
)

type StatusType string

const (
	STATUS_ENABLED  StatusType = "1"
	STATUS_DISABLED StatusType = "0"
)

type User struct {
	ID       *int    `json:"id"`
	Email    *string `json:"email"`
	Password *string `json:"password,omitempty"`
	//Password string     `json:"-"`
	Name      *string     `json:"name"`
	Age       *int        `json:"age"`
	Sex       *SexType    `json:"sex"`
	Status    *StatusType `json:"status"`
	CreatedAt *string     `json:"createdAt"`
	UpdatedAt *string     `json:"updatedAt"`
	//Status StatusType `json:"status" binding:"required"`
}

type UserPartial User

type UserFilter struct {
	User
	Sex           *[]SexType    `json:"sex,omitempty"`
	Status        *[]StatusType `json:"status,omitempty"`
	CreatedAtFrom *string       `json:"createdAtFrom,omitempty"`
	CreatedAtTo   *string       `json:"createdAtTo,omitempty"`
	UpdatedAtFrom *string       `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo   *string       `json:"updatedAtTo,omitempty"`
}

var UserRules = map[string][]validation.Rule{
	"Name":     []validation.Rule{validation.Required},
	"Email":    []validation.Rule{validation.Required, is.Email},
	"Password": []validation.Rule{validation.By(helpers.RequiredIf(true)), validation.Length(6, 100)},
	"Age":      []validation.Rule{validation.Required, validation.By(helpers.NumberRule())},
	"Sex":      []validation.Rule{validation.Required, validation.In(SEX_MALE, SEX_FEMALE).Error("Must be a valid sex value")},
	"Status":   []validation.Rule{validation.Required, validation.In(STATUS_ENABLED, STATUS_DISABLED).Error("Must be a valid status value")},
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, UserRules["Name"]...),
		validation.Field(&u.Email, UserRules["Email"]...),
		validation.Field(&u.Password, UserRules["Password"]...),
		validation.Field(&u.Age, UserRules["Age"]...),
		validation.Field(&u.Sex, UserRules["Sex"]...),
		validation.Field(&u.Status, UserRules["Status"]...),
	)
}
func (u *UserPartial) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, UserRules["Name"][1:]...),
		validation.Field(&u.Email, UserRules["Email"][1:]...),
		validation.Field(&u.Age, UserRules["Age"][1:]...),
		validation.Field(&u.Sex, UserRules["Sex"][1:]...),
		validation.Field(&u.Status, UserRules["Status"][1:]...),
	)
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(password))
}
