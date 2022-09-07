package repository

import (
	"blog/infrastructure"
	"blog/models"
)

//BookRepository -> BookRepository
type BookRepository struct {
	db infrastructure.Database
}

//NewBookRepository fetching database
func NewBookRepository(db infrastructure.Database) BookRepository {
	return BookRepository{
		db: db,
	}
}

//Save -> Method for saving book to database
func (b BookRepository) Save(book models.Book) error {
	return b.db.DB.Create(&book).Error
}

//FindAll -> Method for fetching all books from database
func (b BookRepository) FindAll(book models.Book) (*[]models.Book, int64, error) {
	var books []models.Book
	var totalRows int64 = 0

	queryBuider := b.db.DB.Order("created_at desc").Model(&models.Book{})

	err := queryBuider.
		Where(book).
		Find(&books).
		Count(&totalRows).Error
	return &books, totalRows, err
}

//Update -> Method for updating Book
func (b BookRepository) Update(book models.Book) error {
	return b.db.DB.Save(&book).Error
}

//Find -> Method for fetching book by id
func (b BookRepository) Find(book models.Book) (models.Book, error) {
	var books models.Book
	err := b.db.DB.
		Debug().
		Model(&models.Book{}).
		Where(&book).
		Take(&books).Error
	return books, err
}
