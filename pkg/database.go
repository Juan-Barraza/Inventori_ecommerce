package pkg

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
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

	db, err := retryConnectDB(conPostgres)
	if err != nil {
		log.Fatal("error to connected to database %w", err)
	}
	log.Println("DB conectada correctamente: ")
	return &Database{DB: db}, nil
}

func (d *Database) WithTransaction(fc func(tx *Database) error) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		return fc(&Database{DB: tx})
	})
}

func retryConnectDB(conPostgres string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		// Intentar abrir la conexión a la base de datos
		db, err = gorm.Open(postgres.Open(conPostgres), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
		if err == nil {
			return db, nil
		}
		// Si la conexión falla, espera 2 segundos y vuelve a intentar
		fmt.Printf("Error al conectar a la base de datos, reintentando... (intento %d/%d)\n", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after %d retries: %w", maxRetries, err)
}
