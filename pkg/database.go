package pkg

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) Create(order *domain.Order) {
	panic("unimplemented")
}

func NewDatabase() (*Database, error) {
	conPostgres := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(conPostgres), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

func (d *Database) WithTransaction(fc func(tx *Database) error) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		return fc(&Database{DB: tx})
	})
}
