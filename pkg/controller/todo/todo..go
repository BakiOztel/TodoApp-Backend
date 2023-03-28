package controller

import (
	"github.com/BakiOztel/practice-1/pkg/model"

	"github.com/gofiber/fiber/v2"
)

type CreateTodoRequestBody struct {
	TodoText string `json:"todotext"`
	UserID   uint64 `json:"userid"`
}

type DeleteTodoRequestBody struct {
	TodoId uint64 `json:"todoid"`
}

func CreateTodo(c *fiber.Ctx) error {
	body := CreateTodoRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	todo := model.TodoModel{
		TodoText: body.TodoText,
		UserID:   body.UserID,
	}
	if err := model.CreateTodo(&todo); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else {
		return c.Status(fiber.StatusOK).JSON(&todo)
	}
}

func GetUserTodo(c *fiber.Ctx) error {
	if param := c.Params("userid"); param != "" {
		todoList, err := model.GetUserTodo(param)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(todoList)
	} else {
		return fiber.NewError(fiber.StatusBadRequest)
	}

}
func GetAllTodo(c *fiber.Ctx) error {
	todoList, err := model.GetAllTodo()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	} else {
		return c.Status(fiber.StatusOK).JSON(todoList)
	}
}

func DeleteTodo(c *fiber.Ctx) error {
	body := DeleteTodoRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	if check, err := model.DeleteTodo(body.TodoId); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	} else if check != 1 {
		return fiber.NewError(fiber.StatusBadRequest)
	} else {
		return c.SendStatus(fiber.StatusOK)
	}

}
