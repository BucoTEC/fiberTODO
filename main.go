package main

import (
	"github.com/BucoTEC/golang-todo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//handler for routing
func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "success":  true,
            "message": "You are at the endpoint 😉",
        })
    })

	api := app.Group("/api")
	routes.TodoRoute(api.Group("/todos"))
}

func main(){
	app:= fiber.New()
	app.Use(logger.New())
	
	setupRoutes(app)

	err := app.Listen(":5000")

	if err != nil {
		panic(err)
	}
}