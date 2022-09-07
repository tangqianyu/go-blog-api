package service

import (
	"blog/api/repository"
	"blog/models"
)

//BookService struct
type BookService struct {
	repository repository.BookRepository
}

//NewBookService : returns the BookService struct instance
func NewBookService(r repository.BookRepository) BookService {
	return BookService{
		repository: r,
	}
}

//Save -> calls book repository save method
func (b BookService) Save(book models.Book) error {
	return b.repository.Save(book)
}

//FindAll -> calls book repo find all method
func (b BookService) FindAll(book models.Book) (*[]models.Book, int64, error) {
	return b.repository.FindAll(book)
}

// Update -> calls bookrepo update method
func (b BookService) Update(book models.Book) error {
	return b.repository.Update(book)
}


// Find -> calls book repo find method
func (b BookService) Find(book models.Book) (models.Book, error) {
	return b.repository.Find(book)
}
