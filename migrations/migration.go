package migrations

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Author string `gorm:"not null"`
	Year   int    `gorm:"not null"`
}

func getDBName() string {
	return os.Getenv("DB_NAME")
}

func getAdminDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return db
}

func EnsureDatabaseExists() {
	adminDB := getAdminDB()
	dbName := getDBName()

	// Check if DB exists
	var count int64
	adminDB.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count)
	if count == 0 {
		log.Printf("Database %s does not exist. Creating...", dbName)
		adminDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	}
}

func RunMigrations() {
	EnsureDatabaseExists()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		getDBName(),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run Migrations
	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrations completed successfully!")
}
