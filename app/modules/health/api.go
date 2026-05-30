package health

import (
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/util"
	"github.com/gofiber/fiber/v3"
)

func Status(c fiber.Ctx) error {
	util.Log("INFO", "health", "api", "Status()")

	res := &response.Response{}
	res.Ok(c.Path(), "ok")

	return c.Status(res.Status).JSON(res)
}
