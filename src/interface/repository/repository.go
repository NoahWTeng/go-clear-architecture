package repository

import "github.com/NoahWTeng/go-clear-architecture/common"

type DBHandler interface {
	FindAllBooks() ([]*common.Book, error)
	SaveBook(book common.Book) error
	FindAllAuthor() ([]*common.Author, error)
	SaveAuthor(author common.Author) error
}
