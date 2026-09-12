package nomen

import (
	"github.com/amsatrio/fiber-crud-demo/app/middleware"
	"github.com/gofiber/fiber/v3"
)

func Router(api fiber.Router) {
	nomen := api.Group("/nomen")
	nomen.Get("", NomenPage)
	nomen.Get("/:id", GetByID)
	nomen.Post("", middleware.AuthenticationMiddleware(), Create)
	nomen.Put("/:id", middleware.AuthenticationMiddleware(), Update)
	nomen.Delete("/:id", middleware.AuthenticationMiddleware(), Delete)
}
