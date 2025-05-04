package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Kushal-Dalasaniya/golang-backend/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	/* PostgreSQL connection string */
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	fmt.Println(dsn)

	/* Initialize database */
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db = _db
	fmt.Println("🚀 Database connected successfully!")

	/* Auto-migrate tables */
	db.AutoMigrate(&entity.User{})
}

func GetDB() *gorm.DB {
	return db
}