package t_token

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTTokenRepository(initializer.DB)
	service := NewTTokenService(repo)
	handler := NewTTokenHandler(service, validate)

	api.Post("/t-token", handler.TTokenCreate)
	api.Put("/t-token", handler.TTokenUpdate)
	api.Get("/t-token/:id", handler.TTokenIndex)
	api.Get("/t-token", handler.TTokenPage)
	api.Delete("/t-token/:id", handler.TTokenDelete)

	//api.Get("/generator/t-token/:size", handler.GenerateTToken)
}
