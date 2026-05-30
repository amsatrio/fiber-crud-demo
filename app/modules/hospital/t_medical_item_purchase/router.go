package t_medical_item_purchase

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTMedicalItemPurchaseRepository(initializer.DB)
	service := NewTMedicalItemPurchaseService(repo)
	handler := NewTMedicalItemPurchaseHandler(service, validate)

	api.Post("/t-medical-item-purchase", handler.TMedicalItemPurchaseCreate)
	api.Put("/t-medical-item-purchase", handler.TMedicalItemPurchaseUpdate)
	api.Get("/t-medical-item-purchase/:id", handler.TMedicalItemPurchaseIndex)
	api.Get("/t-medical-item-purchase", handler.TMedicalItemPurchasePage)
	api.Delete("/t-medical-item-purchase/:id", handler.TMedicalItemPurchaseDelete)

	//api.Get("/generator/t-medical-item-purchase/:size", handler.GenerateTMedicalItemPurchase)
}
