package controllers

import (
	"fmt"     // new
	"strconv" // new

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
    Id        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

var todos = []*Todo{
    {
        Id:        1,
        Title:     "Walk the dog 🦮",
        Completed: false,
    },
    {
        Id:        2,
        Title:     "Walk the cat 🐈",
        Completed: false,
    },
}

// get all todos
func GetTodos(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "success": true,
        "data": fiber.Map{
            "todos": todos,
        },
    })
}

// add todo
func CreateTodo(c *fiber.Ctx) error {
    type Request struct {
        Title string `json:"title"`
    }

    var body Request

    err := c.BodyParser(&body)

    // if error
    if err != nil {
        fmt.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "success":  false,
            "message": "Cannot parse JSON",
        })
    }

    // create a todo variable
    todo := &Todo{
        Id:        len(todos) + 1,
        Title:     body.Title,
        Completed: false,
    }

    // append in todos
    todos = append(todos, todo)

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "success": true,
        "data": fiber.Map{
            "todo": todo,
        },
    })
}

//get singel todo
func GetTodo(c *fiber.Ctx) error {
    // get parameter value
    paramId := c.Params("id")

    // convert parameter value string to int
    id, err := strconv.Atoi(paramId)

    // if error in parsing string to int
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "success":  false,
            "message": "Cannot parse Id",
        })
    }

    // find todo and return
    for _, todo := range todos {
        if todo.Id == id {
            return c.Status(fiber.StatusOK).JSON(fiber.Map{
                "success": true,
                "data": fiber.Map{
                    "todo": todo,
                },
            })
        }
    }

    // if todo not available
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
        "success": false,
        "message": "Todo not found",
    })
}
