package t_doctor_office_treatment_price

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewTDoctorOfficeTreatmentPriceRepository(initializer.DB)
	service := NewTDoctorOfficeTreatmentPriceService(repo)
	handler := NewTDoctorOfficeTreatmentPriceHandler(service, validate)

	api.Post("/t-doctor-office-treatment-price", handler.TDoctorOfficeTreatmentPriceCreate)
	api.Put("/t-doctor-office-treatment-price", handler.TDoctorOfficeTreatmentPriceUpdate)
	api.Get("/t-doctor-office-treatment-price/:id", handler.TDoctorOfficeTreatmentPriceIndex)
	api.Get("/t-doctor-office-treatment-price", handler.TDoctorOfficeTreatmentPricePage)
	api.Delete("/t-doctor-office-treatment-price/:id", handler.TDoctorOfficeTreatmentPriceDelete)

	//api.Get("/generator/t-doctor-office-treatment-price/:size", handler.GenerateTDoctorOfficeTreatmentPrice)
}
