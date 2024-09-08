package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tantowish/padi-payment-be/models"
	"gorm.io/gorm"
)

type PaymentController struct {
	DB *gorm.DB
}

func NewPaymentController(DB *gorm.DB) PaymentController {
	return PaymentController{DB}
}

func (pc *PaymentController) GetList(ctx *gin.Context) {
	var paymentCategories []*models.PaymentCategory
	if err := pc.DB.Preload("Payments", func(db *gorm.DB) *gorm.DB {
		return db.Order("id ASC")
	}).Find(&paymentCategories).Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": paymentCategories})
}

func (pc *PaymentController) GetSuggestion(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	type FrequentPayment struct {
		PaymentID uint
	}

	var frequentPayment FrequentPayment

	query := `
		WITH recent_transactions AS (
			SELECT payment_id, created_at
			FROM transactions
			WHERE user_id = ?
			ORDER BY created_at DESC
			LIMIT 10
		),
		payment_counts AS (
			SELECT payment_id, COUNT(*) AS count, MAX(created_at) AS latest_created_at
			FROM recent_transactions
			GROUP BY payment_id
		)
		SELECT payment_id
		FROM payment_counts
		ORDER BY count DESC, latest_created_at DESC
		LIMIT 1;
	`

	if err := pc.DB.Raw(query, currentUser.ID).Scan(&frequentPayment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to retrieve frequent payment ID"})
		return
	}

	if frequentPayment.PaymentID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "info", "message": "No frequent payment found"})
		return
	}

	var payment models.Payment
	if err := pc.DB.First(&payment, frequentPayment.PaymentID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Payment not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": payment})
}
