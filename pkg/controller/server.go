package controller

import (
	"log"

	todoc "github.com/BakiOztel/practice-1/pkg/controller/todo"

	userc "github.com/BakiOztel/practice-1/pkg/controller/user"

	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

func initHandlers(f *fiber.App) {

	f.Get("/user/all", userc.GetAllUser)
	f.Post("/user/register", userc.Register)
	f.Post("user/login", userc.Login)
	f.Post("/todo/create", todoc.CreateTodo)
	f.Get("/todo/:userid", todoc.GetUserTodo)
	f.Delete("/todo/delete", todoc.DeleteTodo)
}

func ServerStart() {
	App := fiber.New()
	initHandlers(App)
	if err := App.Listen(":3000"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
