package main

import (
	"fmt"
	"log"

	"github.com/tantowish/padi-payment-be/initializers"
	"github.com/tantowish/padi-payment-be/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_status') THEN
				CREATE TYPE transaction_status AS ENUM ('PAID', 'CANCEL', 'PENDING');
			END IF;
		END
		$$;
	`)
	if err := initializers.DB.AutoMigrate(&models.User{}, &models.PaymentCategory{}, &models.Payment{}, &models.Transaction{}); err != nil {
		log.Fatal("Error during migration: ", err)
	}	
	fmt.Println("👍 Migration complete")
}
