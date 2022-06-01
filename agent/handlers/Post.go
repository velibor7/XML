package handlers

import (
	"github.com/XML/agent/database"
	"github.com/XML/agent/model"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	is_user, username := IsUser(c)
	if is_user && !IsOwner(c, username) {
		post := new(model.Post)
		if err := c.BodyParser(post); err != nil {
			return c.SendStatus(503)
		}
		//redirect to home page
		return database.CreatePost(post)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}
