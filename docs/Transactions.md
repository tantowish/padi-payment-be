# Transactions API Spec

### Create Transactions

Endpoint: POST /api/v1/transactions

Header :
	Authorization: Bearer access_token

Request Body:

```
{
	"payment_id": 1,
	"user_id": "0f8c7831-e2c3-441b-8dbb-9db3ebb5a69a",
	"gross_amount": 1200000
}
```

Response Body (Success):

```
{
	"success": true,
	"message": "transaction success",
	"data": {
		"id": "47e41fbf-5a4e-40b9-b789-bc3f02a181ab"
		"payment_id": 1,
		"user_id": "0f8c7831-e2c3-441b-8dbb-9db3ebb5a69a",
		"gross_amount": 1200000,
		"status": "PENDING",
		"created_at": "2024-09-01 21:00:00",
		"updated_at": "2024-09-01 21:00:00"
	}
}
```

Response Body (Failed):

```
{
	"status": "error",
	"message": "failed to make transaction"
}
```

### Complete Transactions

Endpoint: POST /api/v1/transactions/:transaction_id/complete

Header :
	Authorization: Bearer access_token

Response Body (Success):

```
{
	"success": true,
	"message": "transaction success",
	"data": {
		"id": "47e41fbf-5a4e-40b9-b789-bc3f02a181ab"
		"payment_id": 1,
		"user_id": "0f8c7831-e2c3-441b-8dbb-9db3ebb5a69a",
		"gross_amount": 1200000,
		"status": "PAID",
		"created_at": "2024-09-01 21:00:00",
		"updated_at": "2024-09-01 21:00:00"
	}
}
```

Response Body (Failed):

```
{
	"status": "error",
	"message": "failed to complete transaction"
}
```
