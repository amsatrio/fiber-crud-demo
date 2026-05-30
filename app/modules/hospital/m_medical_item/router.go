package m_medical_item

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMMedicalItemRepository(initializer.DB)
	service := NewMMedicalItemService(repo)
	handler := NewMMedicalItemHandler(service, validate)

	api.Post("/m-medical-item", handler.MMedicalItemCreate)
	api.Put("/m-medical-item", handler.MMedicalItemUpdate)
	api.Get("/m-medical-item/:id", handler.MMedicalItemIndex)
	api.Get("/m-medical-item", handler.MMedicalItemPage)
	api.Delete("/m-medical-item/:id", handler.MMedicalItemDelete)

	//api.Get("/generator/m-medical-item/:size", handler.GenerateMMedicalItem)
}
