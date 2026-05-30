package m_medical_item_category

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMMedicalItemCategoryRepository(initializer.DB)
	service := NewMMedicalItemCategoryService(repo)
	handler := NewMMedicalItemCategoryHandler(service, validate)

	api.Post("/m-medical-item-category", handler.MMedicalItemCategoryCreate)
	api.Put("/m-medical-item-category", handler.MMedicalItemCategoryUpdate)
	api.Get("/m-medical-item-category/:id", handler.MMedicalItemCategoryIndex)
	api.Get("/m-medical-item-category", handler.MMedicalItemCategoryPage)
	api.Delete("/m-medical-item-category/:id", handler.MMedicalItemCategoryDelete)

	//api.Get("/generator/m-medical-item-category/:size", handler.GenerateMMedicalItemCategory)
}
