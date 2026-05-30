package t_treatment_discount

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTTreatmentDiscountRepository(initializer.DB)
	service := NewTTreatmentDiscountService(repo)
	handler := NewTTreatmentDiscountHandler(service, validate)

	api.Post("/t-treatment-discount", handler.TTreatmentDiscountCreate)
	api.Put("/t-treatment-discount", handler.TTreatmentDiscountUpdate)
	api.Get("/t-treatment-discount/:id", handler.TTreatmentDiscountIndex)
	api.Get("/t-treatment-discount", handler.TTreatmentDiscountPage)
	api.Delete("/t-treatment-discount/:id", handler.TTreatmentDiscountDelete)

	//api.Get("/generator/t-treatment-discount/:size", handler.GenerateTTreatmentDiscount)
}
