package repository

import (
	"books-management/config"
	"books-management/models"
)

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := config.DB.Find(&books)
	return books, result.Error
}

func GetBookByID(id int) (models.Book, error) {
	var book models.Book
	result := config.DB.First(&book, id)
	return book, result.Error
}

func CreateBook(book *models.Book) error {
	return config.DB.Create(book).Error
}

func UpdateBook(id int, book *models.Book) error {
	return config.DB.Model(&models.Book{}).Where("id = ?", id).Updates(book).Error
}

func DeleteBook(id int) error {
	return config.DB.Delete(&models.Book{}, id).Error
}

func BookExistsByID(id int) (bool, error) {
	var count int64
	err := config.DB.Model(&models.Book{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func BookExistByTitle(title string) (bool, error) {
	var count int64
	err := config.DB.Model(&models.Book{}).Where("LOWER(title) = LOWER(?)", title).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
