package controllers

import (
	"strconv"
	"time"

	"api-go-lang-transations/initializers"
	"api-go-lang-transations/lib/enums"
	model "api-go-lang-transations/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateTransaction creates a new transaction.
//
// @Summary Create a new transaction
// @Description Create a new transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Router /transactions [post]
// @Param transaction body model.CreateTransactionSchema true "Transaction"
// @Success 200 {object} model.JSONDataResult{data=[]model.Transaction}
// @Failure	500	{object} model.JSONCodeResult{}
// @Failure 502 {object} model.JSONCodeResult{}
func CreateTransaction(c *fiber.Ctx) error {
	var payload *model.CreateTransactionSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 500, "status": "fail", "message": err.Error()})
	}

	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	now := time.Now()
	newTransaction := model.Transaction{
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
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"code": 500, "status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"code": 200, "status": "success", "data": newTransaction, "message": "Transaction created successfully"})
}

// UpdateTransaction updates a transaction by ID.
//
// @Summary Update a transaction by ID
// @Description Update a transaction by its unique ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Router /transactions/{transactionsId} [put]
// @Param transactionsId path int true "Transaction ID"
// @Param transaction body model.UpdateTransactionSchema true "Transaction"
// @Success 200 {object} model.JSONDataResult{data=[]model.UpdateTransactionSchema}
// @Failure	500	{object} model.JSONCodeResult{}
// @Failure 502 {object} model.JSONCodeResult{}
// @Failure	404	{object} model.JSONCodeResult{}
func UpdateTransaction(c *fiber.Ctx) error {
	idStr := c.Params("transactionsId")

	var payload model.UpdateTransactionSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 500, "status": "fail", "message": "Error parsing body: " + err.Error()})
	}

	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var transaction model.Transaction
	result := initializers.DB.First(&transaction, "id = ?", idStr)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": 404, "status": "fail", "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": 500, "status": "fail", "message": "Error fetching transaction:" + result.Error.Error()})
	}

	transaction.Title = payload.Title
	transaction.Description = payload.Description
	transaction.Status = enums.Status(payload.Status)
	transaction.Amount = payload.Amount
	transaction.FromUser = payload.FromUser
	transaction.ToUser = payload.ToUser

	if err := initializers.DB.Save(&transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": 500, "status": "fail", "message": "Failed to update transaction: " + err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 200, "status": "success", "message": "Transaction updated successfully", "data": transaction})
}

// FindTransactionById finds a transaction by ID.
//
// @Summary Find a transaction by ID
// @Description Find a transaction by its unique ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Router /transactions/{transactionsId} [get]
// @Param transactionsId path int true "Transaction ID"
// @Success 200 {object} model.JSONDataResult{data=[]model.Transaction}
// @Failure	500	{object} model.JSONCodeResult{}
// @Failure	404	{object} model.JSONCodeResult{}
func FindTransactionById(c *fiber.Ctx) error {
	idStr := c.Params("transactionsId")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 500, "status": "fail", "message": "Invalid ID"})
	}

	var transaction model.Transaction
	result := initializers.DB.First(&transaction, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": 404, "status": "fail", "message": "Transaction not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 200, "status": "success", "message": "Transaction found", "data": transaction})
}

// ListTransactions lists transactions with pagination.
//
// @Summary List transactions with pagination
// @Description List transactions with pagination
// @Tags Transactions
// @Accept json
// @Produce json
// @Router /transactions [get]
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} model.JSONDataResultWithPagination{data=[]model.Transaction}
// @Failure	500	{object} model.JSONCodeResult{}
func ListTransactions(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 500, "status": "fail", "message": "Invalid page value"})
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 500, "status": "fail", "message": "Invalid limit value"})
	}

	offset := (intPage - 1) * intLimit

	var transactions []model.Transaction
	var count int64

	initializers.DB.Offset(offset).Limit(intLimit).Find(&transactions)
	initializers.DB.Model(&model.Transaction{}).Count(&count)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   200,
		"status": "success",
		"data":   transactions,
		"offset": offset,
		"total":  int(count),
		"count":  len(transactions),
	})
}

// DeleteTransaction deletes a transaction by ID godoc
//
// @Summary		Delete a transaction by ID
// @Description	Delete a transaction by its unique ID
// @Tags			Transactions
// @Accept			json
// @Produce		json
// @Router /transactions/:transactionsId [delete]
// @Param			id	path		int	true	"transaction ID"
// @Success 200 {object} model.JSONCodeResult{}
// @Failure	404	{object} model.JSONCodeResult{}
// @Failure 500 {object} model.JSONCodeResult{}
func DeleteTransaction(c *fiber.Ctx) error {
	idStr := c.Params("transactionsId")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    404,
			"status":  "fail",
			"message": "No note with that Id exists",
		})
	}

	var transaction model.Transaction
	transaction.ID = uint(id)
	result := initializers.DB.Delete(&transaction)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": 500, "status": "fail", "message": "Failed to delete transaction"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 200, "status": "success", "message": "Transaction successfully deleted"})
}
