package services

import (
	"books-management/cache"
	"books-management/kafka"
	"books-management/models"
	"books-management/repository"
	"encoding/json"
	"fmt"
)

func GetAllBooks() ([]models.Book, error) {
	cachedBooks, err := cache.Get("all_books")
	if err == nil {
		var books []models.Book
		json.Unmarshal([]byte(cachedBooks), &books)
		return books, nil
	}

	books, err := repository.GetAllBooks()
	if err == nil {
		booksJSON, _ := json.Marshal(books)
		cache.Set("all_books", string(booksJSON))
	}
	return books, err
}

func GetBookByID(id int) (models.Book, error) {
	cacheKey := fmt.Sprintf("book_%d", id)
	cachedBook, err := cache.Get(cacheKey)
	if err == nil {
		var book models.Book
		json.Unmarshal([]byte(cachedBook), &book)
		return book, nil
	}

	book, err := repository.GetBookByID(id)
	if err == nil {
		bookJSON, _ := json.Marshal(book)
		cache.Set(cacheKey, string(bookJSON))
	}
	return book, err
}

func BookExistsByID(id int) (bool, error) {
	cacheKey := fmt.Sprintf("book_%d", id)
	cachedBook, err := cache.Get(cacheKey)
	if err == nil {
		var book models.Book
		json.Unmarshal([]byte(cachedBook), &book)
		return false, nil
	}

	found, err := repository.BookExistsByID(id)
	if err == nil {
		bookJSON, _ := json.Marshal(found)
		cache.Set(cacheKey, string(bookJSON))
	}
	return found, err
}

func BookExistByTitle(title string) (bool, error) {
	cacheKey := fmt.Sprintf("book_%s", title)
	cachedBook, err := cache.Get(cacheKey)
	if err == nil {
		var book models.Book
		json.Unmarshal([]byte(cachedBook), &book)
		return false, nil
	}

	found, err := repository.BookExistByTitle(title)
	if err == nil {
		bookJSON, _ := json.Marshal(found)
		cache.Set(cacheKey, string(bookJSON))
	}
	return found, err
}

func CreateBook(book *models.Book) error {
	err := repository.CreateBook(book)
	if err == nil {
		cache.Delete("all_books")
		kafka.Publish("book_events", "Book Created: "+book.Title)
	}
	return err
}

func UpdateBook(id int, book *models.Book) error {
	err := repository.UpdateBook(id, book)
	if err == nil {
		cache.Delete("all_books")
		cache.Delete(fmt.Sprintf("book_%d", id))
		kafka.Publish("book_events", "Book Updated: "+book.Title)
	}
	return err
}

func DeleteBook(id int) error {
	err := repository.DeleteBook(id)
	if err == nil {
		cache.Delete("all_books")
		cache.Delete(fmt.Sprintf("book_%d", id))
		kafka.Publish("book_events", fmt.Sprintf("Book Deleted: %d", id))
	}
	return err
}
