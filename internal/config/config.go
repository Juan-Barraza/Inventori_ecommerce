package config

import (
	"inventory/pkg"
	"log"
	"github.com/joho/godotenv"
)

func SetConfig() (*pkg.Database, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load a file env: %v", err)
	}

	db, err := CreateTables()
	if err != nil {
		return nil, err
	}

	return db, nil
}