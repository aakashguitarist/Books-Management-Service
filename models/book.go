package models

import "gorm.io/gorm"

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title  string `json:"title" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
	Year   int    `json:"year"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
