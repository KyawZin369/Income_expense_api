package models

import "time"

type Currency struct {
	ID        uint   `gorm:"primaryKey"`
	Code      string `gorm:"uniqueIndex;size:3;not null"` // ISO 4217 code (USD, EUR, etc.)
	Name      string `gorm:"size:100;not null"`           // Full name (US Dollar, Euro, etc.)
	Symbol    string `gorm:"size:10;not null"`            // Symbol ($, €, etc.)
	IsActive  bool   `gorm:"default:true"`                // Whether currency is active
	CreatedAt time.Time
	UpdatedAt time.Time
}
