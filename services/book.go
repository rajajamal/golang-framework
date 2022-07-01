package services

import (
	"errors"

	"github.com/rajajamal/golang-framework/models"
	"github.com/rajajamal/golang-framework/repositories"
)

type Book struct {
	Repository repositories.BookRepository
}

func (s *Book) Save(model *models.Book) error {
	err := s.Repository.Save(model)
	if err != nil {
		return errors.New("Title already exists")
	}

	if err != nil {
		return errors.New("Internal error")
	}

	return err
}
