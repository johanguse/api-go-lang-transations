package controllers

import (
	"strconv"
	"time"

	database "api-go-lang-transations/database/schema"
	"api-go-lang-transations/initializers"
	"api-go-lang-transations/lib/enums"
	"api-go-lang-transations/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTransaction(c *fiber.Ctx) error {
	var payload *models.CreateTransactionSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newTransaction := database.Transaction{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      enums.Status(payload.Status),
		Amount:      payload.Amount,
		Date:        now,
		FromUser:    payload.FromUser,
		ToUser:      payload.ToUser,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := initializers.DB.Create(&newTransaction)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"transaction": newTransaction}})
}

func UpdateTransaction(c *fiber.Ctx) error {
	idStr := c.Params("transactionsId")

	var payload models.UpdateTransactionSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Error parsing body: " + err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var transaction database.Transaction
	result := initializers.DB.First(&transaction, "id = ?", idStr)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
		}
		// If some other error occurred while fetching from db
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Error fetching transaction:" + result.Error.Error()})
	}

	transaction.Title = payload.Title
	transaction.Description = payload.Description
	transaction.Status = enums.Status(payload.Status)
	transaction.Amount = payload.Amount
	transaction.FromUser = payload.FromUser
	transaction.ToUser = payload.ToUser

	if err := initializers.DB.Save(&transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to update transaction: " + err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": transaction})
}

func FindTransactionById(c *fiber.Ctx) error {
	idStr := c.Params("transactionsId")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var transaction database.Transaction
	transaction.ID = uint(id)
	initializers.DB.First(&transaction)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": transaction})
}

func ListTransactions(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page value",
		})
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid limit value",
		})
	}

	offset := (intPage - 1) * intLimit

	var transactions []database.Transaction
	var count int64

	initializers.DB.Offset(offset).Limit(intLimit).Find(&transactions)
	initializers.DB.Model(&database.Transaction{}).Count(&count)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   transactions,
		"offset": offset,
		"total":  int(count),
		"count":  len(transactions),
	})
}

func DeleteTransaction(c *fiber.Ctx) error {
	idStr := c.Params("transactionsId")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var transaction database.Transaction
	transaction.ID = uint(id)
	result := initializers.DB.Delete(&transaction)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete transaction",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Transaction successfully deleted",
	})
}
