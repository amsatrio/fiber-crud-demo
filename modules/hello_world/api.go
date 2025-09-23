package hello_world

import (
	"fiber-crud-demo/dto"
	"fiber-crud-demo/util"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	res := &dto.Response{}
	res.Ok(c.Path(), "hello world!")

	return c.Status(res.Status).JSON(res)
}

func HelloWorldPath(c *fiber.Ctx) error {
	res := &dto.Response{}
	res.Ok(c.Path(), c.Params("message"))

	return c.Status(res.Status).JSON(res)
}

func HelloWorldQuery(c *fiber.Ctx) error {
	res := &dto.Response{}
	res.Ok(c.Path(), c.Query("message"))

	return c.Status(res.Status).JSON(res)
}

type HelloWorldRequest struct {
	Message string `json:"message" validate:"required,min=5,max=20"`
}

var validate = validator.New()

func HelloWorldPayload(c *fiber.Ctx) error {
	payload := new(HelloWorldRequest)

	// parse payload
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	res := &dto.Response{}

	// validate payload
	if err := validate.Struct(payload); err != nil {
		out, _ := util.ValidateError(err)
		if out != nil {
			res.ErrMessagePayload(c.Path(), fiber.StatusBadRequest, "invalid payload", out)
			return c.Status(res.Status).JSON(res)
		}
	}

	res.Ok(c.Path(), payload)

	return c.Status(res.Status).JSON(res)
}

func HelloWorldError(c *fiber.Ctx) error {
	error_type := c.Params("type")
	res := &dto.Response{}
	res.Ok(c.Path(), nil)

	if error_type == "503" {
		return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
	}
	if error_type == "500" {
		return fiber.NewError(fiber.StatusInternalServerError, "On vacation!")
	}

	return c.Status(res.Status).JSON(res)
}
