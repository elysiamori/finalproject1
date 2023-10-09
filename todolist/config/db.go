package config

import (
	"github.com/elysiamori/finalproject1/kelompok6/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConn
func DBConn() *gorm.DB {
	conn := "postgres://postgres:postgres@localhost:5432/todos?sslmode=disable"

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		return nil
	}

	db.AutoMigrate(&models.Todos{})

	return db
}
