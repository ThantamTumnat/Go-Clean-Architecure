package controllers

import (
	"userservice/modules/entities"
	"userservice/modules/logs"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userSrv entities.UserService
	hisSrv  entities.HistoryService
	accRepo entities.Account_userRepository
}

func NewUserController(r fiber.Router, userSrv entities.UserService, hisSrv entities.HistoryService, accRepo entities.Account_userRepository) {
	// return userController{userSrv, hisSrv, accRepo}
	controllers := &userController{
		userSrv, hisSrv, accRepo,
	}
	r.Post("/created", controllers.UserCreated)
	r.Get("/readed", controllers.UserReaded)
	r.Put("/updated", controllers.UserUpdated)
	r.Delete("/deleted", controllers.UserDeleted)
}

func (obj userController) UserCreated(c *fiber.Ctx) error {
	command := entities.UserCreated{}
	err := c.BodyParser(&command)
	if err != nil {
		logs.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}
	if command.Username == "" || command.Password == "" || command.Address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}

	created_produce, err := obj.userSrv.UserCreated(&command)
	if err != nil {
		logs.Error(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      "Unprocessable Entity",
			"status_code": fiber.StatusUnprocessableEntity,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "Successfully, Your account was created",
		"user_id":     created_produce.UserID,
		"username":    created_produce.Username,
	})
}

func (obj userController) UserUpdated(c *fiber.Ctx) error {
	command := entities.UserUpdated{}
	err := c.BodyParser(&command)
	if err != nil {
		logs.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}
	if command.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}
	check, err := obj.accRepo.CheckAccount(command.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "Not found",
			"status_code": fiber.StatusNotFound,
		})
	}
	_ = check

	updated, err := obj.userSrv.UserUpdated(&command)
	if err != nil {
		logs.Error(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      "Unprocessable Entity",
			"status_code": fiber.StatusUnprocessableEntity,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "Successfully, Your account was updated",
		"user_id":     updated.UserID,
		"password":    updated.Password,
		"address":     updated.Address,
	})
}

func (obj userController) UserReaded(c *fiber.Ctx) error {

	command := entities.UserReaded{}
	err := c.BodyParser(&command)
	if err != nil {
		logs.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}
	if command.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}
	check, err := obj.accRepo.CheckAccount(command.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "Not found",
			"status_code": fiber.StatusNotFound,
		})
	}
	_ = check

	readed, err := obj.hisSrv.GetHistory(command.UserID)
	if err != nil {
		logs.Error(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      "Unprocessable Entity",
			"status_code": fiber.StatusUnprocessableEntity,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status":      "ok",
		"status_code": fiber.StatusOK,
		"History":     readed,
	})
}

func (obj userController) UserDeleted(c *fiber.Ctx) error {
	command := entities.UserDeleted{}
	err := c.BodyParser(&command)
	if err != nil {
		logs.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}
	if command.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}
	check, err := obj.accRepo.CheckAccount(command.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "Not found",
			"status_code": fiber.StatusNotFound,
		})
	}
	_ = check
	deleted, err := obj.userSrv.UserDeleted(&command)
	if err != nil {
		logs.Error(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      "Unprocessable Entity",
			"status_code": fiber.StatusUnprocessableEntity,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "Successfully, Your account was deleted",
		"user_id":     deleted.UserID,
	})
}
