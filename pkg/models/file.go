package models

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"os"
	"strings"
)

type File struct {
	ID               *int    `json:"id,omitempty"`
	OriginalFileName *string `json:"originalFileName,omitempty"`
	Ext              *string `json:"ext,omitempty"`
	UUID             *string `json:"uuid,omitempty"`

	Data      *map[string]string `json:"data,omitempty" form:"data"`
	ModelName *string            `json:"modelName,omitempty" form:"modelName"`
	ModelId   *int               `json:"modelId,omitempty" form:"modelId"`

	CreatedAt *string `json:"createdAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`

	Url *string `json:"url,omitempty"`
}

type FilePartial File

type FileFilter File

func (m *File) Validate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.ModelName, validation.Required),
		validation.Field(&m.ModelId, validation.Required, validation.By(helpers.NumberRule())),
	)
}

func (m *File) FileName() string {
	return strings.Join([]string{*m.UUID, *m.Ext}, ".")
}

func (m *File) GetUrl() string {
	return fmt.Sprintf("%s/%s/%s", os.Getenv("BACKEND_URL"), "public", m.FileName())
}

func (f *File) MarshalJSON() ([]byte, error) {
	clone := *f
	url := clone.GetUrl()
	clone.Url = &url
	json, err := json.Marshal(clone)
	return json, err
}
