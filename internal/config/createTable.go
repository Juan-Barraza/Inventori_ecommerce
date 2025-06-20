package config

import (
	"fmt"
	domain "inventory/internal/fiber/domain/entities"
	"inventory/pkg"
	"log"
	"os"
)

func CreateTables() (*pkg.Database, error) {
	db, err := pkg.NewDatabase()
	if err != nil {
		return nil, err
	}

	migratesErr := db.DB.AutoMigrate(
		&domain.User{},
		&domain.Provider{},
		&domain.Client{},
		&domain.Category{},
		&domain.Product{},
		&domain.PicturesProduct{},
		&domain.Order{},
		&domain.Transaction{},
	)

	if migratesErr != nil {
		log.Fatalf("Error al crear las tablas: %v", migratesErr)
	} else {
		log.Println("Tablas creadas exitosamente")
	}

	// log.Println("Conexión exitosa a la base de datos")

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	conexionLog, _ := fmt.Printf(
		"Me conecte a la base de datos PostgreSQL, en el servidor %s por el puerto %s ",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	logger.Println(conexionLog)

	return db, nil
}
