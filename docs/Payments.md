# Payments API Spec

### Get Payments

Endpoint: GET /api/v1/payments

Response Body (Success):

```
{
	"success": true,
	"message": "all payment by category",
	"data": [
		{
			"id": "1",
			"category_name": "Vritual Account",
			"payments": [
				{
					"id": "1",
					"payment_name": "BCA VA",
					"logo": "padi-umkmk/bca-logo.png",
					"note": "*min 20rb",
					"created_at": "2024-09-01 21:00:00",
					"updated_at": "2024-09-01 21:00:00",
				},
				{
					"id": "2",
					"payment_name": "BNI VA",
					"logo": "padi-umkmk/bni-logo.png",
					"note": "*min 40rb",
					"created_at": "2024-09-01 21:00:00",
					"updated_at": "2024-09-01 21:00:00",
				},
			]
		},
		{
			"id": "2",
			"category_name": "Dompet Digital",
			"payments": [
				{
					"id": "4",
					"payment_name": "Dana",
					"logo": "padi-umkmk/dana-logo.png",
					"note": NULL,
					"created_at": "2024-09-01 21:00:00",
					"updated_at": "2024-09-01 21:00:00",
				},
				{
					"id": "7",
					"payment_name": "Gopay",
					"logo": "padi-umkmk/gopay-logo.png",
					"note": NULL,
					"created_at": "2024-09-01 21:00:00",
					"updated_at": "2024-09-01 21:00:00",
				},
			]
		},
		{
			"id": "3",
			"category_name": "Kartu Kredit",
			"payments": [
				{
					"id": "9",
					"payment_name": "Mandiri CC",
					"logo": "padi-umkmk/mandiri-logo.png",
					"note": "*min 100rb",
					"created_at": "2024-09-01 21:00:00",
					"updated_at": "2024-09-01 21:00:00",
				},
				{
					"id": "20",
					"payment_name": "BCA CC",
					"logo": "padi-umkmk/bca-logo.png",
					"note": NULL,
					"created_at": "2024-09-01 21:00:00",
					"updated_at": "2024-09-01 21:00:00",
				},
			]
		},
		{
			"id": "4",
			"category_name": "Qris",
			"payments": [
				{
					"id": "24",
					"payment_name": "Qris",
					"logo": "padi-umkmk/qris-logo.png",
					"note": NULL,
					"created_at": "2024-09-01 21:00:00",
					"updated_at": "2024-09-01 21:00:00",
				},
			]
		}
	]
}
```

Response Body (Failed):

```
{
	"status": "error",
	"message": "failed to get payments"
}
```

### Get Payments Recommendation / Suggestion

Endpoint: GET /api/v1/payments/suggest

Header :
	Authorization: Bearer access_token

Response Body (Success):

```
{
	"success": true,
	"message": "payment suggestion",
	"data": {
		"id": "20",
		"payment_name": "BCA CC",
		"logo": "padi-umkmk/bca-logo.png",
		"note": NULL,
		"created_at": "2024-09-01 21:00:00",
		"updated_at": "2024-09-01 21:00:00",
	},
}
```

Response Body (Failed):

```
{
	"status": "error",
	"message": "failed to get payment suggestion"
}
```

Response Body (No Transaction Found):

```
{
	"success": "info",
	"message": "No frequent payment found"
}
```
