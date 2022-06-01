package handlers

import (
	"github.com/XML/agent/database"
	"github.com/XML/agent/model"
	"github.com/gofiber/fiber/v2"
)

func GetCompanies(c *fiber.Ctx) error {
	comps := database.GetCompanies()
	return c.JSON(comps)
}

func GetCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	comp := database.GetCompany(id)
	return c.JSON(comp)
}

func CreateCompanies(c *fiber.Ctx) error {
	is_user, username := IsUser(c)
	if is_user {
		comp := new(model.Company)
		if err := c.BodyParser(comp); err != nil {
			return c.SendStatus(503)
		}

		comp.UserName = username
		//redirect to home page
		return database.CreateCompanies(comp)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}
