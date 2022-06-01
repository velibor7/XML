package main

import (
	"log"

	"github.com/XML/agent/database"
	"github.com/XML/agent/handlers"
	"github.com/XML/agent/model"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	//dont need to have an acc
	users := app.Group("/users")
	companies := app.Group("/companies")

	users.Post("/register", handlers.RegisterUsers)
	users.Post("/login", handlers.LoginUsers)
	companies.Get("", handlers.GetCompanies)
	companies.Get("/:id", handlers.GetCompany)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	admin := app.Group("/admin")
	//need to have an acc
	companies.Post("", handlers.CreateCompanies)
	companies.Post("/:id/posts", handlers.CreatePost)

	//need to be admin
	admin.Get("/companies", handlers.GetAdminCompanies)
	admin.Get("/companies/:id", handlers.GetAdminCompany)
	admin.Post("/companies/:id/activate", handlers.ActivateCompany)
	admin.Get("/users", handlers.GetUsers)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("agent.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect")
	}
	database.DBConn.AutoMigrate(&model.User{}, &model.Company{}, &model.Post{})
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	log.Fatal(app.Listen(":8081"))
}
