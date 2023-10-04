package models

import (
	"github.com/go-playground/validator/v10"

	"api-go-lang-transations/lib/enums"
	"api-go-lang-transations/types"
	"time"
)

var validate = validator.New()

func ValidateStruct[T any](payload T) []*types.ErrorResponse {
	var errors []*types.ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element types.ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// Account Transaction info
// @Description Transaction information
type Transaction struct {
	ID          uint         `gorm:"primaryKey;autoIncrement"`
	Title       string       `gorm:"size:255;not null"`
	Description string       `gorm:"type:text;not null"`
	Status      enums.Status `gorm:"type:enum('processing', 'processed', 'created');not null"`
	Amount      float64      `gorm:"type:decimal(10,2);not null;check:Amount>=0"`
	Date        time.Time    `gorm:"not null"`
	FromUser    string       `gorm:"size:100;not null"`
	ToUser      string       `gorm:"size:100;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
} //@name Transaction

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
