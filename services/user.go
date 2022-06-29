package services

import (
	"github.com/rajajamal/golang-framework/models"
	"github.com/rajajamal/golang-framework/repositories"
	"github.com/rajajamal/golang-framework/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type User struct {
	Repository repositories.UserRepository
}

func (s User) GetPaginated(paginator utils.Paginator) (utils.Paginator, *models.Error) {
	result, err := s.Repository.FindPaginated(paginator)
	if err != nil {
		return result, &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}

	return result, nil
}

func (s User) Get(id int) (models.User, *models.Error) {
	result, err := s.Repository.Find(id)
	if err != nil {
		return result, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	return result, nil
}

func (s User) ValidateLogin(login models.Login) (models.User, *models.Error) {
	result, err := s.Repository.FindByUsername(login.Username)
	if err != nil {
		return models.User{}, &models.Error{
			Message: "user not found or password not match",
			Code:    fiber.StatusBadRequest,
		}
	}

	if !utils.ValidatePassword(result.Password, login.Password) {
		return models.User{}, &models.Error{
			Message: "user not found or password not match",
			Code:    fiber.StatusBadRequest,
		}
	}

	return result, nil
}

func (s User) Create(user models.CreateUser) (models.User, *models.Error) {
	model := models.User{}
	copier.Copy(&model, &user)
	model.Password = utils.EncodePassword(model.Password)
	err := s.Repository.Save(&model)
	if err != nil {
		return model, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusInternalServerError,
		}
	}

	return model, nil
}

func (s User) Delete(id int) *models.Error {
	model, err := s.Repository.Find(id)
	if err != nil {
		return &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	err = s.Repository.Delete(&model)
	if err != nil {
		return &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}

	return nil
}

func (s User) Update(user models.UpdateUser) (models.User, *models.Error) {
	model, err := s.Repository.Find(user.ID)
	if err != nil {
		return model, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	copier.Copy(&model, &user)
	err = s.Repository.Save(&model)
	if err != nil {
		return model, &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}

	return model, nil
}

func (s User) UpdatePassword(input models.UpdatePassword) (models.User, *models.Error) {
	model, err := s.Repository.Find(input.ID)
	if err != nil {
		return model, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	if !utils.ValidatePassword(model.Password, input.OldPassword) {
		return model, &models.Error{
			Message: "old password not match",
			Code:    fiber.StatusBadRequest,
		}
	}

	model.Password = utils.EncodePassword(input.Password)
	err = s.Repository.Save(&model)
	if err != nil {
		return model, &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}

	return model, nil
}
