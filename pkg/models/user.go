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
	ID       *int    `json:"id" swaggerignore:"true"`
	Email    *string `json:"email"`
	Password *string `json:"password,omitempty" swaggerignore:"true"` //Password string     `json:"-"`
	Name     *string `json:"name"`
	Age      *int    `json:"age"`
	// ENUM 1=Male, 2=Female
	Sex *SexType `json:"sex"`
	// ENUM "1"=Enabled, "0"=Disabled
	Status    *StatusType `json:"status"` //Status StatusType `json:"status" binding:"required"`
	CreatedAt *string     `json:"createdAt" swaggerignore:"true"`
	UpdatedAt *string     `json:"updatedAt" swaggerignore:"true"`
}

type UserWithPassword struct {
	User
	Password *string `json:"password"`
}

/*func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   *u.ID,
		Name: *u.Name,
	})
}*/

/*func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	var temp struct {
		Alias
		End string `json:"end"` // Custom processing field
	}

	// Unmarshal into the temporary struct
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	//some logic
	return nil
}*/

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
