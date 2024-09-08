package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionStatus string

const (
	PAID    TransactionStatus = "PAID"
	CANCEL  TransactionStatus = "CANCEL"
	PENDING TransactionStatus = "PENDING"
)

type Transaction struct {
	ID          uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID      uuid.UUID         `gorm:"type:uuid;not null" json:"user_id,omitempty"`
	PaymentID   uint              `gorm:"not null" json:"payment_id,omitempty"`
	GrossAmount uint              `gorm:"not null" json:"gross_amount"`
	NoPayment   string            `gorm:"default: null" json:"no_payment"`
	Status      TransactionStatus `gorm:"type:transaction_status;default:'PENDING'" json:"status"`
	ExpireAt    time.Time         `gorm:"not null" json:"expire_at,omitempty"`
	CreatedAt   time.Time         `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time         `gorm:"not null" json:"updated_at,omitempty"`

	Payment Payment `gorm:"foreignKey:PaymentID"`
	User    User    `gorm:"foreignKey:UserID"`
}

type CreateTransactionRequest struct {
	PaymentID   uint `json:"payment_id"  binding:"required"`
	GrossAmount uint `json:"gross_amount" binding:"required"`
}

type TransactionResponse struct {
	ID          uuid.UUID         `json:"id"`
	UserID      uuid.UUID         `json:"user_id"`
	PaymentID   uint              `json:"payment_id"`
	GrossAmount uint              `json:"gross_amount"`
	NoPayment   string            `json:"no_payment"`
	Status      TransactionStatus `json:"status"`
	ExpireAt    time.Time         `json:"expire_at"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type TransactionDetailResponse struct {
	ID          uuid.UUID         `json:"id"`
	UserID      uuid.UUID         `json:"user_id"`
	PaymentID   uint              `json:"payment_id"`
	GrossAmount uint              `json:"gross_amount"`
	NoPayment   string            `json:"no_payment"`
	Status      TransactionStatus `json:"status"`
	ExpireAt    time.Time         `json:"expire_at"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	User        UserResponse      `json:"user"`
	Payment     Payment           `json:"payment"`
}
