# Transactions API Spec

### Get Transactions

Endpoint: POST /api/v1/transactions/:transaction_id

Header :
	Authorization: Bearer access_token

Response Body (Success):

```
{
	"status": "success"
    	"data": {
       		"id": "71454552-bdb7-4ce1-a3d8-5f31ceb22aed",
        	"user_id": "853eba85-960a-4785-a5ae-fb48e115999f",
        	"payment_id": 1,
        	"gross_amount": 10000,
        	"no_payment": "19865250074",
        	"status": "PENDING",
        	"created_at": "2024-09-06T13:23:44.578201Z",
        	"updated_at": "2024-09-06T13:23:44.578201Z",
        	"user": {
            		"id": "853eba85-960a-4785-a5ae-fb48e115999f",
            		"name": "user",
            		"email": "user@gmail.com",
            		"created_at": "2024-09-06T13:14:19.357018Z",
            		"updated_at": "2024-09-06T13:14:19.357018Z"
        	},
        	"payment": {
            		"id": 1,
            		"category_id": 1,
            		"payment_name": "Mandiri",
            		"logo": "mandiri.jpg",
            		"note": "",
            		"created_at": "2024-09-06T00:00:00Z",
            		"updated_at": "2024-09-06T00:00:00Z"
        }
    },
}
```

Response Body (Failed):

```
{
	"status": "error",
	"message": "failed to make transaction"
}
```

Response Body (404 Not Found):

```
{
	"status": "error",
	"message": "record not found"
}
```


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
	"status": success,
	"data": {
		"id": "47e41fbf-5a4e-40b9-b789-bc3f02a181ab"
		"payment_id": 1,
		"user_id": "0f8c7831-e2c3-441b-8dbb-9db3ebb5a69a",
		"gross_amount": 1200000,
		"status": "PENDING",
		"no_payment": "12241421321",
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

Endpoint: PUT /api/v1/transactions/:transaction_id/complete

Header :
	Authorization: Bearer access_token

Response Body (Success):

```
{
	"status": "success",
	"data": {
		"id": "47e41fbf-5a4e-40b9-b789-bc3f02a181ab"
		"payment_id": 1,
		"user_id": "0f8c7831-e2c3-441b-8dbb-9db3ebb5a69a",
		"gross_amount": 1200000,
		"status": "PAID",
		"no_payment": "12241421321",
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
