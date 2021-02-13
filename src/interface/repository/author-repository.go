package repository

import (
	"github.com/NoahWTeng/go-clear-architecture/common"
)

type AuthorRepo struct {
	handler DBHandler
}

func (repo AuthorRepo) FindAll() ([]*common.Author, error) {
	panic("implement me")
}

func NewAuthorRepo(handler DBHandler) AuthorRepo {
	return AuthorRepo{handler}
}

func (repo AuthorRepo) SaveAuthor(author common.Author) error {
	err := repo.handler.SaveAuthor(author)
	if err != nil {
		return err
	}
	return err
}

func (repo AuthorRepo) FindAllAuthor(author common.Author) ([]*common.Author, error) {
	resp, err := repo.handler.FindAllAuthor()

	if err != nil {
		return resp, err
	}

	return resp, err

}
