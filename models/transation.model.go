package models

import (
	"github.com/go-playground/validator/v10"

	"api-go-lang-transations/lib/enums"
)

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type CreateTransactionSchema struct {
	Title       string       `json:"title" validate:"required"`
	Description string       `json:"description" validate:"required"`
	Status      enums.Status `json:"status,omitempty" validate:"omitempty,oneof=processing processed created"`
	Amount      float64      `json:"amount" validate:"required,gte=0"`
	FromUser    string       `json:"fromuser" validate:"required"`
	ToUser      string       `json:"touser" validate:"required"`
} //@name CreateTransaction

type UpdateTransactionSchema struct {
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Status      enums.Status `json:"status,omitempty" validate:"omitempty,oneof=processing processed created"`
	Amount      float64      `json:"amount,omitempty" validate:"omitempty,gte=0"`
	Date        string       `json:"date,omitempty"`
	FromUser    string       `json:"fromuser,omitempty"`
	ToUser      string       `json:"touser,omitempty"`
} //@name UpdateTransaction
