package handlers

import (
	"time"

	"github.com/XML/agent/database"
	"github.com/XML/agent/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RegisterUsers(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(503)
	}
	//redirect to home page
	return database.RegisterUsers(user)

}

func LoginUsers(c *fiber.Ctx) error {
	// user := c.FormValue("username")
	// pass := c.FormValue("password")
	login := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&login); err != nil {
		return err
	}
	oldUser := database.GetUser(login.Username)

	if oldUser.Password != login.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	claims := jwt.MapClaims{
		"username": oldUser.Username,
		"userType": oldUser.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}
