package m_blood_group

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMBloodGroupRepository(initializer.DB)
	service := NewMBloodGroupService(repo)
	handler := NewMBloodGroupHandler(service, validate)

	api.Post("/m-blood-group", handler.MBloodGroupCreate)
	api.Put("/m-blood-group", handler.MBloodGroupUpdate)
	api.Get("/m-blood-group/:id", handler.MBloodGroupIndex)
	api.Get("/m-blood-group", handler.MBloodGroupPage)
	api.Delete("/m-blood-group/:id", handler.MBloodGroupDelete)

	//api.Get("/generator/m-blood-group/:size", handler.GenerateMBloodGroup)
}
