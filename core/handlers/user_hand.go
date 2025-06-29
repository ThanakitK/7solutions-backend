package handlers

import (
	"7solutions/backend/core/models"
	"7solutions/backend/core/services"

	"github.com/gofiber/fiber/v2"
)

type userHand struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHand {
	return userHand{
		userSrv: userSrv,
	}
}

func (h userHand) CreateUser(c *fiber.Ctx) error {
	body := models.SrvCreateUserModel{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	result := h.userSrv.CreateUser(body)
	return c.Status(result.Code).JSON(result)
}

func (h userHand) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result := h.userSrv.GetUserByID(id)
	return c.Status(result.Code).JSON(result)
}

func (h userHand) SignIn(c *fiber.Ctx) error {
	body := models.SrvSignInModel{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	result := h.userSrv.SignIn(body)
	return c.Status(result.Code).JSON(result)
}

func (h userHand) GetUsers(c *fiber.Ctx) error {
	result := h.userSrv.Gets()
	return c.Status(result.Code).JSON(result)
}

func (h userHand) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	body := models.SrvUpdateUserModel{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	result := h.userSrv.UpdateUser(id, body)
	return c.Status(result.Code).JSON(result)
}

func (h userHand) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	result := h.userSrv.DeleteUser(id)
	return c.Status(result.Code).JSON(result)
}
