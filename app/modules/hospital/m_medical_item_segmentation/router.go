package m_medical_item_segmentation

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMMedicalItemSegmentationRepository(initializer.DB)
	service := NewMMedicalItemSegmentationService(repo)
	handler := NewMMedicalItemSegmentationHandler(service, validate)

	api.Post("/m-medical-item-segmentation", handler.MMedicalItemSegmentationCreate)
	api.Put("/m-medical-item-segmentation", handler.MMedicalItemSegmentationUpdate)
	api.Get("/m-medical-item-segmentation/:id", handler.MMedicalItemSegmentationIndex)
	api.Get("/m-medical-item-segmentation", handler.MMedicalItemSegmentationPage)
	api.Delete("/m-medical-item-segmentation/:id", handler.MMedicalItemSegmentationDelete)

	//api.Get("/generator/m-medical-item-segmentation/:size", handler.GenerateMMedicalItemSegmentation)
}
