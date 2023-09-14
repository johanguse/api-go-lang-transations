package seeds

import (
	"time"

	"gorm.io/gorm"

	database "api-go-lang-transations/database/schema"
)

type TransactionSeed struct{}

func (t *TransactionSeed) Seed(db *gorm.DB) {
	transactions := []database.Transaction{
		{
			Title:       "Resgate",
			Description: "et labore proident aute nulla",
			Status:      "created",
			Amount:      2078.66,
			Date:        time.Date(2018, 12, 22, 0, 0, 0, 0, time.UTC),
			FromUser:    "Aposentadoria",
			ToUser:      "Conta Warren",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Depósito",
			Description: "excepteur veniam proident irure pariatur",
			Status:      "created",
			Amount:      148856.29,
			Date:        time.Date(2017, 7, 23, 0, 0, 0, 0, time.UTC),
			FromUser:    "Trade",
			ToUser:      "Conta Warren",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Movimentação interna",
			Description: "eu officia laborum labore aute",
			Status:      "processed",
			Amount:      25092.8,
			Date:        time.Date(2016, 8, 25, 0, 0, 0, 0, time.UTC),
			FromUser:    "Férias",
			ToUser:      "Trade",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Pagamento",
			Description: "adipisicing elit eiusmod laboris",
			Status:      "created",
			Amount:      1000.0,
			Date:        time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC),
			FromUser:    "Conta Warren",
			ToUser:      "Fornecedor",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Transferência",
			Description: "incididunt veniam laboris nostrud",
			Status:      "processing",
			Amount:      5000.0,
			Date:        time.Date(2021, 9, 30, 0, 0, 0, 0, time.UTC),
			FromUser:    "Conta Warren",
			ToUser:      "Conta Corrente",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Compra",
			Description: "exercitation culpa laboris est",
			Status:      "created",
			Amount:      150.0,
			Date:        time.Date(2021, 9, 29, 0, 0, 0, 0, time.UTC),
			FromUser:    "Cartão de Crédito",
			ToUser:      "Conta Warren",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Transferência",
			Description: "incididunt veniam laboris nostrud",
			Status:      "processing",
			Amount:      10000.0,
			Date:        time.Date(2021, 9, 30, 0, 0, 0, 0, time.UTC),
			FromUser:    "Conta Warren",
			ToUser:      "Conta Corrente",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Compra",
			Description: "exercitation culpa laboris est",
			Status:      "created",
			Amount:      250.0,
			Date:        time.Date(2021, 9, 29, 0, 0, 0, 0, time.UTC),
			FromUser:    "Cartão de Crédito",
			ToUser:      "Conta Warren",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, transaction := range transactions {
		db.Create(&transaction)
	}
}
