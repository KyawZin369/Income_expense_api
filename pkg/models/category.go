package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	User      User      `gorm:"foreignKey:UserID"`
	Name      string    `gorm:"size:100;not null"`
	Type      string    `gorm:"size:20;not null"`
	Color     *string   `gorm:"size:7"`  // Hex color code, nullable
	Icon      *string   `gorm:"size:50"` // Icon identifier, nullable
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
