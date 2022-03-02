package routes

import (
	"github.com/BucoTEC/golang-todo/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
    route.Get("", controllers.GetTodos)
}