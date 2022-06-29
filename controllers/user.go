package controllers

import (
	"strconv"

	"github.com/rajajamal/golang-framework/models"
	"github.com/rajajamal/golang-framework/services"
	"github.com/rajajamal/golang-framework/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Service services.User
}

func (c User) Create(ctx *fiber.Ctx) error {
	form := models.CreateUser{}
	ctx.BodyParser(&form)
	messages, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": messages,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	model, status := c.Service.Create(form)
	if status != nil {
		ctx.JSON(map[string]string{
			"message": status.Message,
		})

		return ctx.SendStatus(status.Code)
	}

	ctx.JSON(model)

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c User) Update(ctx *fiber.Ctx) error {
	form := models.UpdateUser{}
	ctx.BodyParser(&form)
	messages, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": messages,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	form.ID, _ = strconv.Atoi(ctx.Params("id"))
	model, status := c.Service.Update(form)
	if status != nil {
		ctx.JSON(map[string]string{
			"message": status.Message,
		})

		return ctx.SendStatus(status.Code)
	}

	return ctx.JSON(model)
}

func (c User) UpdatePassword(ctx *fiber.Ctx) error {
	form := models.UpdatePassword{}
	ctx.BodyParser(&form)
	messages, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": messages,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	form.ID, _ = strconv.Atoi(ctx.Params("id"))
	model, status := c.Service.UpdatePassword(form)
	if status != nil {
		ctx.JSON(map[string]string{
			"message": status.Message,
		})

		return ctx.SendStatus(status.Code)
	}

	return ctx.JSON(model)
}

func (c User) Login(ctx *fiber.Ctx) error {
	form := models.Login{}
	ctx.BodyParser(&form)
	messages, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": messages,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	result, status := c.Service.ValidateLogin(form)
	if status != nil {
		ctx.JSON(map[string]string{
			"message": status.Message,
		})

		return ctx.SendStatus(status.Code)
	}

	claims := jwt.MapClaims{}
	claims["username"] = result.Username
	ctx.JSON(map[string]string{
		"token": utils.CreateJwtToken(claims),
	})

	return ctx.SendStatus(fiber.StatusOK)
}

func (c User) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	err := c.Service.Delete(id)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Message,
		})

		return ctx.SendStatus(err.Code)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c User) Get(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	user, err := c.Service.Get(id)
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Message,
		})

		return ctx.SendStatus(err.Code)
	}

	return ctx.JSON(user)
}

func (c User) GetPaginated(ctx *fiber.Ctx) error {
	users, err := c.Service.GetPaginated(*utils.NewPaginator(ctx))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "unable to get all users",
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(users)
}
