package repositories

import (
	"github.com/rajajamal/golang-framework/models"

	"gorm.io/gorm"
)

type book struct {
	storage *gorm.DB
}

func NewBookRepository(storage *gorm.DB) BookRepository {
	storage.AutoMigrate(&models.Book{})

	return book{storage: storage}
}

func (r book) Save(user *models.Book) error {
	return r.storage.Save(user).Error
}
