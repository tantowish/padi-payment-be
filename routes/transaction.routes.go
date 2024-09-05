package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tantowish/padi-payment-be/controllers"
	"github.com/tantowish/padi-payment-be/middleware"
)

type TransactionRouteController struct {
	transactionController controllers.TransactionController
}

func NewRouteTransactionController(transactionController controllers.TransactionController) TransactionRouteController {
	return TransactionRouteController{transactionController}
}

func (tc *TransactionRouteController) TransactionRoute(rg *gin.RouterGroup) {

	router := rg.Group("transactions")
	router.Use(middleware.DeserializeUser())
	router.POST("/", tc.transactionController.Create)
	router.PUT("/:id/complete", tc.transactionController.Update)
	router.GET("/:id", tc.transactionController.Get)
}
