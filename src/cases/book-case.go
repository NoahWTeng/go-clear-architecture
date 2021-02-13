package cases

import (
	"log"

	"github.com/NoahWTeng/go-clear-architecture/common"
)

type BookInteractor struct {
	BookRepository common.BookRepository
}

func NewBookInteractor(repository common.BookRepository) BookInteractor {
	return BookInteractor{repository}
}

func (interactor *BookInteractor) CreateBook(book common.Book) error {
	err := interactor.BookRepository.SaveBook(book)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *BookInteractor) FindAll() ([]*common.Book, error) {
	results, err := interactor.BookRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return results, nil
}
