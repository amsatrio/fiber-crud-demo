package health

import (
	"fiber-crud-demo/dto"
	"fiber-crud-demo/util"

	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx) error {
	util.Log("INFO", "health", "api", "Status()")

	res := &dto.Response{}
	res.Ok(c.Path(), "ok")

	return c.Status(res.Status).JSON(res)
}
