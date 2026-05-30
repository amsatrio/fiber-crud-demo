package m_menu_role

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMMenuRoleRepository(initializer.DB)
	service := NewMMenuRoleService(repo)
	handler := NewMMenuRoleHandler(service, validate)

	api.Post("/m-menu-role", handler.MMenuRoleCreate)
	api.Put("/m-menu-role", handler.MMenuRoleUpdate)
	api.Get("/m-menu-role/:id", handler.MMenuRoleIndex)
	api.Get("/m-menu-role", handler.MMenuRolePage)
	api.Delete("/m-menu-role/:id", handler.MMenuRoleDelete)

	//api.Get("/generator/m-menu-role/:size", handler.GenerateMMenuRole)
}
