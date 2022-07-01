package repositories

import (
	"github.com/rajajamal/golang-framework/models"
	"github.com/rajajamal/golang-framework/utils"
)

type UserRepository interface {
	Find(id int) (models.User, error)
	FindByUsername(username string) (models.User, error)
	FindPaginated(paginator utils.Paginator) (utils.Paginator, error)
	Save(user *models.User) error
	Delete(user *models.User) error
}

type BookRepository interface {
	Save(model *models.Book) error
}
