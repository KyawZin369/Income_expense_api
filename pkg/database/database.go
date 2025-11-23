package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/example/graphql-mysql-api/pkg/models"
)

func Connect(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Currency{},
		&models.Account{},
		&models.Category{},
		&models.Transaction{},
		&models.Expense{},
		&models.Income{},
		&models.Exchange{},
		&models.Transfer{},
	)
}
