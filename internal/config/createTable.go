package config

import (
	"fmt"
	modelsgorm "inventory/internal/fiber/infrastructure/persistence/modelsGORM"
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
		&modelsgorm.User{},
		&modelsgorm.Provider{},
		&modelsgorm.Client{},
		&modelsgorm.Category{},
		&modelsgorm.Product{},
		&modelsgorm.PicturesProduct{},
		&modelsgorm.Order{},
		&modelsgorm.Transaction{},
	)

	if migratesErr != nil {
		return nil, migratesErr
	}

	log.Println("Conexión exitosa a la base de datos")

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	conexionLog, _ := fmt.Printf(
		"Me conecte a la base de datos PostgreSQL, en el servidor %s por el puerto %s ",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
	)

	log.Println(conexionLog)
	logger.Println(conexionLog)

	return db, nil
}
