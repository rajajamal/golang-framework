package services

import (
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	mocks "github.com/rajajamal/golang-framework/mocks/repositories"
	"github.com/rajajamal/golang-framework/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Book_Create_Error(t *testing.T) {
	form := models.CreateBook{
		Title:  "Bumi Manusia",
		Author: "Pram",
	}

	repository := mocks.BookRepository{}
	repository.On("Save", mock.Anything).Return(errors.New("")).Once()

	service := Book{Repository: &repository}
	_, err := service.Create(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusInternalServerError)
}
