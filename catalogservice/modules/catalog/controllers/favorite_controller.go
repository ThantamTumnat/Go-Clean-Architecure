package controllers

import (
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"

	"github.com/gofiber/fiber/v2"
)

type favoriteController struct {
	accRepo     entities.AccountRepository
	catalogRepo entities.CatalogRepository
	favSrv      entities.FavoriteService
}

func NewFavoriteHandler(r fiber.Router, accRepo entities.AccountRepository, catalogRepo entities.CatalogRepository, favSrv entities.FavoriteService) {
	// return favoriteController{accRepo, catalogRepo, favSrv}
	controllers := &favoriteController{
		accRepo, catalogRepo, favSrv,
	}
	r.Post("/favorite", controllers.CreatedFavorite)
	r.Get("/getFavorites", controllers.GetFavorite)

}

func (h favoriteController) GetFavorite(c *fiber.Ctx) error {
	command := entities.UserGetFavorite{}
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
	check, err := h.accRepo.CheckAccount(command.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "Not found",
			"status_code": fiber.StatusNotFound,
		})
	}
	_ = check

	// readed, err := obj.hisSrv.GetHistory(command.UserID)
	favorite, err := h.favSrv.GetFavorite(command.UserID)
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
		"favorites":   favorite,
	})
}

func (h favoriteController) CreatedFavorite(c *fiber.Ctx) error {
	command := entities.UserFavorite{}
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

	// created, err := h.catalogSrv.GetCatalog(command.UserID, command.ID)
	created, err := h.favSrv.CreatedFavorite(command.UserID, command.ID)
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
		"user_id":     command.UserID,
		"favorite":    created,
	})
}
