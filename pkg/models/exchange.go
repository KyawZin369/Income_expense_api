package models

import "time"

// Exchange model for currency conversions
type Exchange struct {
	ID             uint      `gorm:"primaryKey"`
	UserID         uint      `gorm:"not null;index"`
	User           User      `gorm:"foreignKey:UserID"`
	FromAccountID  uint      `gorm:"not null"`
	FromAccount    Account   `gorm:"foreignKey:FromAccountID"`
	ToAccountID    uint      `gorm:"not null"`
	ToAccount      Account   `gorm:"foreignKey:ToAccountID"`
	FromAmount     float64   `gorm:"type:decimal(15,2);not null"` // Amount in source currency
	ToAmount       float64   `gorm:"type:decimal(15,2);not null"` // Amount in target currency
	FromCurrencyID uint      `gorm:"not null"`
	FromCurrency   Currency  `gorm:"foreignKey:FromCurrencyID"`
	ToCurrencyID   uint      `gorm:"not null"`
	ToCurrency     Currency  `gorm:"foreignKey:ToCurrencyID"`
	ExchangeRate   float64   `gorm:"type:decimal(15,6);not null"` // Exchange rate used
	Description    string    `gorm:"size:255"`
	Date           time.Time `gorm:"not null;index"`
	Status         string    `gorm:"size:20;default:'completed'"`
	Fee            float64   `gorm:"type:decimal(15,2);default:0"` // Exchange fee
	FeeCurrencyID  *uint     `gorm:"index"`
	FeeCurrency    *Currency `gorm:"foreignKey:FeeCurrencyID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
