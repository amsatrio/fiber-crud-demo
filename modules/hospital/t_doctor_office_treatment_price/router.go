package t_doctor_office_treatment_price
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTDoctorOfficeTreatmentPriceRepository(initializer.DB)
	service := NewTDoctorOfficeTreatmentPriceService(repo)
	handler := NewTDoctorOfficeTreatmentPriceHandler(service, validate)
	api := app.Group("/v1/t-doctor-office-treatment-price")
	api.Post("", handler.TDoctorOfficeTreatmentPriceCreate)
	api.Put("", handler.TDoctorOfficeTreatmentPriceUpdate)
	api.Get(":id", handler.TDoctorOfficeTreatmentPriceIndex)
	api.Get("", handler.TDoctorOfficeTreatmentPricePage)
	api.Delete(":id", handler.TDoctorOfficeTreatmentPriceDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-doctor-office-treatment-price/:size", handler.GenerateTDoctorOfficeTreatmentPrice)
}

