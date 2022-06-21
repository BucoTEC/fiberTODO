package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error  {
		return c.JSON("hello dera frined")
	})

	app.Get("/:name", func (c *fiber.Ctx) error  {
		name := c.Params("name")
		msg := fmt.Sprintf("hello dera frined %v", name)
		return c.JSON(msg)
	})

	 app.Listen(":5000") 


	
}
