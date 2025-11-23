package models

import "time"

type Account struct {
	ID         uint     `gorm:"primaryKey"`
	UserID     uint     `gorm:"not null;index"`
	User       User     `gorm:"foreignKey:UserID"`
	Name       string   `gorm:"size:100;not null"` // Account name (e.g., "Main Checking")
	Type       string   `gorm:"size:50;not null"`  // Account type (checking, savings, credit, etc.)
	CurrencyID uint     `gorm:"not null"`
	Currency   Currency `gorm:"foreignKey:CurrencyID"`
	Balance    float64  `gorm:"type:decimal(15,2);default:0"` // Current balance
	IsActive   bool     `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
