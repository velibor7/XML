package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func IsAdmin(c *fiber.Ctx) bool {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["username"].(string) == "admin" {
		return true
	}
	return false
}

func IsUser(c *fiber.Ctx) (bool, string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["userType"] == nil {
		return false, ""
	}
	return true, claims["username"].(string)
}

func IsOwner(c *fiber.Ctx, username string) bool {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["username"].(string) == username {
		return true
	}
	return false
}
