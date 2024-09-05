package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tantowish/padi-payment-be/controllers"
)

type PaymentRouteController struct {
	paymentController controllers.PaymentController
}

func NewRoutePaymentController(paymentController controllers.PaymentController) PaymentRouteController {
	return PaymentRouteController{paymentController}
}

func (pc *PaymentRouteController) PaymentRoute(rg *gin.RouterGroup) {

	router := rg.Group("payments")
	router.GET("/", pc.paymentController.GetList)
	router.GET("/suggest", pc.paymentController.GetSuggestion)
}
