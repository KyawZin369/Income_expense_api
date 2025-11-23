package models

import "time"

// Transfer model for money transfers between accounts
type Transfer struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null;index"`
	User          User      `gorm:"foreignKey:UserID"`
	FromAccountID uint      `gorm:"not null"`
	FromAccount   Account   `gorm:"foreignKey:FromAccountID"`
	ToAccountID   uint      `gorm:"not null"`
	ToAccount     Account   `gorm:"foreignKey:ToAccountID"`
	Amount        float64   `gorm:"type:decimal(15,2);not null"` // Transfer amount
	CurrencyID    uint      `gorm:"not null"`
	Currency      Currency  `gorm:"foreignKey:CurrencyID"`
	Description   string    `gorm:"size:255"`
	Date          time.Time `gorm:"not null;index"`
	Status        string    `gorm:"size:20;default:'completed'"`
	Fee           float64   `gorm:"type:decimal(15,2);default:0"` // Transfer fee
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
