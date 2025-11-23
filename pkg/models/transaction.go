package models

import "time"

// Base transaction model
type Transaction struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null;index"`
	User        User      `gorm:"foreignKey:UserID"`
	AccountID   uint      `gorm:"not null"`
	Account     Account   `gorm:"foreignKey:AccountID"`
	CategoryID  *uint     `gorm:"index"` // Optional category
	Category    *Category `gorm:"foreignKey:CategoryID"`
	Amount      float64   `gorm:"type:decimal(15,2);not null"` // Transaction amount
	CurrencyID  uint      `gorm:"not null"`
	Currency    Currency  `gorm:"foreignKey:CurrencyID"`
	Description string    `gorm:"size:255"`                    // Transaction description
	Date        time.Time `gorm:"not null;index"`              // Transaction date
	Type        string    `gorm:"size:20;not null"`            // Transaction type (expense, income, transfer, exchange)
	Status      string    `gorm:"size:20;default:'completed'"` // Transaction status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
