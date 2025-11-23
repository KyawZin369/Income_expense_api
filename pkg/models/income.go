package models

import "time"

// Income model extending Transaction
type Income struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null;index"`
	User        User      `gorm:"foreignKey:UserID"`
	AccountID   uint      `gorm:"not null"`
	Account     Account   `gorm:"foreignKey:AccountID"`
	CategoryID  *uint     `gorm:"index"`
	Category    *Category `gorm:"foreignKey:CategoryID"`
	Amount      float64   `gorm:"type:decimal(15,2);not null"` // Income amount (positive)
	CurrencyID  uint      `gorm:"not null"`
	Currency    Currency  `gorm:"foreignKey:CurrencyID"`
	Description string    `gorm:"size:255"`
	Date        time.Time `gorm:"not null;index"`
	Status      string    `gorm:"size:20;default:'completed'"`
	Source      string    `gorm:"size:100"`      // Income source (employer, freelance, etc.)
	IsRecurring bool      `gorm:"default:false"` // Whether this is recurring income
	Frequency   string    `gorm:"size:20"`       // Frequency if recurring (monthly, weekly, etc.)
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
