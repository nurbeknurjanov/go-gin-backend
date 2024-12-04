package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Product struct {
	ID          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedAt   *string `json:"createdAt,omitempty"`
	UpdatedAt   *string `json:"updatedAt,omitempty"`

	ImageID *int  `json:"imageId,omitempty"`
	Image   *File `json:"image,omitempty"`
}

type ProductPartial Product

type ProductFilter Product

var ProductRules = map[string][]validation.Rule{
	"Name":        []validation.Rule{validation.Required},
	"Description": []validation.Rule{},
}

func (m *Product) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Name, ProductRules["Name"]...),
		validation.Field(&m.Description, ProductRules["Description"]...),
	)
}

func (m *ProductPartial) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Name, ProductRules["Name"][1:]...),
		validation.Field(&m.Description, ProductRules["Description"]...),
	)
}
