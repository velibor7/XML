package handlers

import (
	"github.com/XML/agent/database"
	"github.com/gofiber/fiber/v2"
)

func GetAdminCompany(c *fiber.Ctx) error {
	if !IsAdmin(c) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	id := c.Params("id")
	comp := database.GetAdminCompany(id)
	return c.JSON(comp)
}

func GetAdminCompanies(c *fiber.Ctx) error {
	if !IsAdmin(c) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	comps := database.GetAdminCompanies()
	return c.JSON(comps)

}

func GetUsers(c *fiber.Ctx) error {
	if !IsAdmin(c) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	users := database.GetUsers()
	return c.JSON(users)
}

func ActivateCompany(c *fiber.Ctx) error {
	if !IsAdmin(c) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	id := c.Params("id")
	return database.ActivateCompany(id)
}
