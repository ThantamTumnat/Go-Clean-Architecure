package controllers

import (
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"

	"github.com/gofiber/fiber/v2"
)

type catalogController struct {
	catalogSrv  entities.CatalogService
	accRepo     entities.AccountRepository
	catalogRepo entities.CatalogRepository
	// favSrv      services.FavoriteService
}

func NewCatalogController(r fiber.Router, catalogSrv entities.CatalogService, accRepo entities.AccountRepository, catalogRepo entities.CatalogRepository) {
	controllers := &catalogController{
		catalogSrv, accRepo, catalogRepo,
	}

	r.Get("/read", controllers.GetCatalog)
	r.Get("/getCatalogs", controllers.GetCatalogs)

}

func (h catalogController) GetCatalogs(c *fiber.Ctx) error {
	command := entities.UserGetCatalogs{}
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
	checkAccount, err := h.accRepo.CheckAccount(command.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "Not found",
			"status_code": fiber.StatusNotFound,
		})
	}
	_ = checkAccount
	catalogs, err := h.catalogSrv.GetCatalogs()
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
		"Message":     "Successfully, Your can get all of the joke stories",
		"user_id":     command.UserID,
		"catalogs":    catalogs,
	})
}

func (h catalogController) GetCatalog(c *fiber.Ctx) error {
	command := entities.UserRead{}
	err := c.BodyParser(&command)
	if err != nil {
		logs.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}
	if command.UserID == "" || command.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad request",
			"status_code": fiber.StatusBadRequest,
		})
	}

	checkAccount, err := h.accRepo.CheckAccount(command.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "Not found",
			"status_code": fiber.StatusNotFound,
		})
	}
	_ = checkAccount
	checkCatalog, err := h.catalogRepo.GetCatalogId(command.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "Not found",
			"status_code": fiber.StatusNotFound,
		})
	}
	_ = checkCatalog

	getCatalog, err := h.catalogSrv.GetCatalog(command.UserID, command.ID)
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
		"Message":     "Successfully, Your account was created",
		"user_id":     command.UserID,
		"histories":   getCatalog,
	})
}
