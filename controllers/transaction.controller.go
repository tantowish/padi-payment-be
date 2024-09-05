package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tantowish/padi-payment-be/models"
	"github.com/tantowish/padi-payment-be/utils"
	"gorm.io/gorm"
)

type TransactionController struct {
	DB *gorm.DB
}

func NewTransactionController(DB *gorm.DB) TransactionController {
	return TransactionController{DB}
}

func (pc *TransactionController) Create(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateTransactionRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	combinedPaymentNumber := fmt.Sprintf("%d%d", payload.PaymentID, utils.GenerateRandomNumber())

	now := time.Now()
	newTransaction := models.Transaction{
		UserID:      currentUser.ID,
		PaymentID:   payload.PaymentID,
		GrossAmount: payload.GrossAmount,
		NoPayment:   combinedPaymentNumber,
		Status:      models.PENDING,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := pc.DB.Create(&newTransaction)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	response := models.TransactionResponse{
		ID:          newTransaction.ID,
		UserID:      newTransaction.UserID,
		PaymentID:   newTransaction.PaymentID,
		GrossAmount: newTransaction.GrossAmount,
		NoPayment:   newTransaction.NoPayment,
		Status:      newTransaction.Status,
		CreatedAt:   newTransaction.CreatedAt,
		UpdatedAt:   newTransaction.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func (tc *TransactionController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	var transaction models.Transaction
	uuidID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid UUID format"})
		return
	}

	if err := tc.DB.Preload("Payment").Preload("User").First(&transaction, "id = ?", uuidID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
		return
	}

	response := models.TransactionDetailResponse{
		ID:          transaction.ID,
		UserID:      transaction.UserID,
		PaymentID:   transaction.PaymentID,
		GrossAmount: transaction.GrossAmount,
		NoPayment:   transaction.NoPayment,
		Status:      transaction.Status,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
		User: models.UserResponse{
			ID:        transaction.User.ID,
			Name:      transaction.User.Name,
			Email:     transaction.User.Email,
			CreatedAt: transaction.User.CreatedAt,
			UpdatedAt: transaction.User.UpdatedAt,
		},
		Payment: transaction.Payment,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func (tc *TransactionController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var transaction models.Transaction
	uuidID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid UUID format"})
		return
	}

	if err := tc.DB.First(&transaction, "id = ?", uuidID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
		return
	}

	transaction.Status = models.PAID
	transaction.UpdatedAt = time.Now()

	if err := tc.DB.Save(&transaction).Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	response := models.TransactionResponse{
		ID:          transaction.ID,
		UserID:      transaction.UserID,
		PaymentID:   transaction.PaymentID,
		GrossAmount: transaction.GrossAmount,
		NoPayment:   transaction.NoPayment,
		Status:      transaction.Status,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}
