package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
)

type ErrorResponse struct {
	Message      string            `json:"message"`
	Code         int               `json:"code,omitempty"`
	FieldsErrors map[string]string `json:"fieldsErrors,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func newErrorResponse(c *gin.Context, statusCode int, err error) {
	//json marshal error
	if errorValue, isUnmarshalTypeError := err.(*json.UnmarshalTypeError); isUnmarshalTypeError == true {
		newErrorResponseWithFieldsErrors(c, statusCode,
			errors.New("Validation error"),
			422,
			map[string]string{errorValue.Field: fmt.Sprintf("The field '%s' is wrong type", helpers.FirstToUpper(errorValue.Field))})
		return
	}

	//json missing fields error
	if errorValue, isJsonValidationErrors := err.(validator.ValidationErrors); isJsonValidationErrors == true {
		fieldsErrors := map[string]string{}
		for _, value := range errorValue {
			fieldsErrors[helpers.FirstToLower(value.Field())] = fmt.Sprintf("The field '%s' is missing", value.Field())
		}

		newErrorResponseWithFieldsErrors(c, statusCode,
			errors.New("Validation error"),
			422,
			fieldsErrors)
		return
	}

	//ozzo validation
	if errorValue, isValidationErrors := err.(validation.Errors); isValidationErrors == true {
		fieldsErrors := map[string]string{}
		for key, value := range errorValue {
			fieldsErrors[helpers.FirstToLower(key)] = helpers.FirstToUpper(value.Error())
		}

		newErrorResponseWithFieldsErrors(c, statusCode, errors.New("Validation error"), 422, fieldsErrors)
		return
	}

	//fmt.Println("->", reflect.TypeOf(err).String())
	//logrus.Error(err.Error())
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: err.Error()})
}

func newErrorResponseWithFieldsErrors(c *gin.Context, statusCode int, error error, errorCode int, fieldsErrors map[string]string) {
	//logrus.Error(error.Error())
	c.AbortWithStatusJSON(statusCode,
		ErrorResponse{
			Message:      error.Error(),
			Code:         errorCode,
			FieldsErrors: fieldsErrors,
		})
}
