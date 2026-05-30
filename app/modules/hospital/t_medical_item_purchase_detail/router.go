package t_medical_item_purchase_detail

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTMedicalItemPurchaseDetailRepository(initializer.DB)
	service := NewTMedicalItemPurchaseDetailService(repo)
	handler := NewTMedicalItemPurchaseDetailHandler(service, validate)

	api.Post("/t-medical-item-purchase-detail", handler.TMedicalItemPurchaseDetailCreate)
	api.Put("/t-medical-item-purchase-detail", handler.TMedicalItemPurchaseDetailUpdate)
	api.Get("/t-medical-item-purchase-detail/:id", handler.TMedicalItemPurchaseDetailIndex)
	api.Get("/t-medical-item-purchase-detail", handler.TMedicalItemPurchaseDetailPage)
	api.Delete("/t-medical-item-purchase-detail/:id", handler.TMedicalItemPurchaseDetailDelete)

	//api.Get("/generator/t-medical-item-purchase-detail/:size", handler.GenerateTMedicalItemPurchaseDetail)
}
