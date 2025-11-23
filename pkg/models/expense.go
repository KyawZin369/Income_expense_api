package models

import "time"

// Expense model extending Transaction
type Expense struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null;index"`
	User        User      `gorm:"foreignKey:UserID"`
	AccountID   uint      `gorm:"not null"`
	Account     Account   `gorm:"foreignKey:AccountID"`
	CategoryID  *uint     `gorm:"index"`
	Category    *Category `gorm:"foreignKey:CategoryID"`
	Amount      float64   `gorm:"type:decimal(15,2);not null"` // Expense amount (positive)
	CurrencyID  uint      `gorm:"not null"`
	Currency    Currency  `gorm:"foreignKey:CurrencyID"`
	Description string    `gorm:"size:255"`
	Date        time.Time `gorm:"not null;index"`
	Status      string    `gorm:"size:20;default:'completed'"`
	Tags        string    `gorm:"size:500"` // Comma-separated tags
	Location    string    `gorm:"size:255"` // Location where expense occurred
	ReceiptURL  string    `gorm:"size:500"` // URL to receipt image
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
