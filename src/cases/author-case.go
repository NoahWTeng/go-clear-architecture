package cases

import (
	"github.com/NoahWTeng/go-clear-architecture/common"
	"log"
)

type AuthorInteracts struct {
	AuthorRepository common.AuthorRepository
}

func NewAuthorInteractor(repository common.AuthorRepository) AuthorInteracts {
	return AuthorInteracts{repository}
}

func (interactor *AuthorInteracts) CreateAuthor(author common.Author) error {
	err := interactor.AuthorRepository.SaveAuthor(author)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *AuthorInteracts) FindAll() ([]*common.Author, error) {
	results, err := interactor.AuthorRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return results, nil
}
