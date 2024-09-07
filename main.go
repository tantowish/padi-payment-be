package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tantowish/padi-payment-be/controllers"
	"github.com/tantowish/padi-payment-be/initializers"
	"github.com/tantowish/padi-payment-be/routes"
)

var (
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController

	PaymentController      controllers.PaymentController
	PaymentRouteController routes.PaymentRouteController

	TransactionController      controllers.TransactionController
	TransactionRouteController routes.TransactionRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	PaymentController = controllers.NewPaymentController(initializers.DB)
	PaymentRouteController = routes.NewRoutePaymentController(PaymentController)

	TransactionController = controllers.NewTransactionController(initializers.DB)
	TransactionRouteController = routes.NewRouteTransactionController(TransactionController)
}


func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://localhost:3000", config.ClientOrigin}
	// corsConfig.AllowCredentials = true

	server := gin.Default()
	// Custom CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", config.ClientOrigin}, // Replace config.ClientOrigin with the actual allowed client origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},              // Specify allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},   // Specify allowed headers
		ExposeHeaders:    []string{"Content-Length"},                            // Specify which headers are exposed to the client
		AllowCredentials: true,                                                  // Allow credentials like cookies
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000" || origin == config.ClientOrigin
		},
	}

	// Apply CORS middleware
	server.Use(cors.New(corsConfig))

	router := server.Group("/api/v1")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)
	PaymentRouteController.PaymentRoute(router)
	TransactionRouteController.TransactionRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
