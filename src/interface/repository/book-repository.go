package repository

import (
	"github.com/NoahWTeng/go-clear-architecture/common"
)

type BookRepo struct {
	handler DBHandler
}

func NewBookRepo(handler DBHandler) BookRepo {
	return BookRepo{handler}
}

func (repo BookRepo) SaveBook(book common.Book) error {
	err := repo.handler.SaveBook(book)
	if err != nil {
		return err
	}
	return err
}

func (repo BookRepo) FindAll() ([]*common.Book, error) {
	results, err := repo.handler.FindAllBooks()
	if err != nil {
		return results, err
	}
	return results, err
}
