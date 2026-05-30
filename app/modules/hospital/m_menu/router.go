package m_menu

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMMenuRepository(initializer.DB)
	service := NewMMenuService(repo)
	handler := NewMMenuHandler(service, validate)

	api.Post("/m-menu", handler.MMenuCreate)
	api.Put("/m-menu", handler.MMenuUpdate)
	api.Get("/m-menu/:id", handler.MMenuIndex)
	api.Get("/m-menu", handler.MMenuPage)
	api.Delete("/m-menu/:id", handler.MMenuDelete)

	//api.Get("/generator/m-menu/:size", handler.GenerateMMenu)
}
