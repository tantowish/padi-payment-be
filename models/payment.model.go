package models

import "time"

type Payment struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	CategoryID  uint      `gorm:"not null" json:"category_id,omitempty"`
	PaymentName string    `gorm:"not null" json:"payment_name"`
	Logo        string    `gorm:"not null" json:"logo"`
	Note        string    `gorm:"default:null" json:"note"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at,omitempty"`

	Transactions []Transaction   `gorm:"foreignKey:PaymentID" json:"transactions,omitempty"`
}
