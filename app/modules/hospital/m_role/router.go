package m_role

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMRoleRepository(initializer.DB)
	service := NewMRoleService(repo)
	handler := NewMRoleHandler(service, validate)

	api.Post("/m-role", handler.MRoleCreate)
	api.Put("/m-role", handler.MRoleUpdate)
	api.Get("/m-role/:id", handler.MRoleIndex)
	api.Get("/m-role", handler.MRolePage)
	api.Delete("/m-role/:id", handler.MRoleDelete)

	//api.Get("/generator/m-role/:size", handler.GenerateMRole)
}
