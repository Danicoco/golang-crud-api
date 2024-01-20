package database

import (
	"crud-with-postgres/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConn() (*gorm.DB, error) {
	dsn := configs.LoadEnvs().DB_URL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db

	// db.AutoMigrate(&models.Book{})
	return db, err
}
