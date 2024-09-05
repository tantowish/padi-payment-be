package models

type PaymentCategory struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	CategoryName string `gorm:"not null" json:"category_name"`
	
	Payments     []Payment `gorm:"foreignKey:CategoryID" json:"payments,omitempty"`
}
