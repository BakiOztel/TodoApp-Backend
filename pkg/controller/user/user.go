package controller

import (
	"github.com/BakiOztel/practice-1/pkg/middleware"
	"github.com/BakiOztel/practice-1/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type RegisterUserRequestBody struct {
	UserName string `json:"userName"`
	Pw       string `json:"password"`
}

// I haven't added too many genres because I'm in the learning stage.
// Normally I need to ask for more information in the registration section
type LoginUserRequestBody struct {
	UserName string `json:"userName"`
	Pw       string `json:"password"`
}

var Store session.Store

func GetAllUser(c *fiber.Ctx) error {

	userList, err := model.GetAllUser()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	} else {
		return c.Status(fiber.StatusOK).JSON(userList)
	}
}

func Register(c *fiber.Ctx) error {
	body := RegisterUserRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var user model.UserSt

	user.UserName = body.UserName
	hash, err := middleware.HashPassword(user.Pw)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	user.Pw = hash
	if err := model.CreateUser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else {
		return c.Status(fiber.StatusOK).JSON(&user)
	}
}

func Login(c *fiber.Ctx) error {
	body := LoginUserRequestBody{}
	sess, _ := Store.Get(c)

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var user model.UserSt

	user.UserName = body.UserName
	user.Pw = body.Pw
	userg, err := model.Login(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if check := middleware.CheckPasswordHash(user.Pw, userg.Pw); check == nil {
		sess.Set("token", "token")
		return c.Status(fiber.StatusOK).JSON(&userg)
	} else {

		return fiber.NewError(fiber.StatusBadRequest, check.Error())
	}

}
