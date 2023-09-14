package database

import (
	"api-go-lang-transations/lib/enums"
	"time"

	"gorm.io/gorm"
)

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
}

func CreateTransactionTable(DB *gorm.DB) {
	DB.Migrator().CreateTable(&Transaction{})
}
