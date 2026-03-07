package m_medical_facility
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMMedicalFacilityRepository(initializer.DB)
	service := NewMMedicalFacilityService(repo)
	handler := NewMMedicalFacilityHandler(service, validate)
	api := app.Group("/v1/m-medical-facility")
	api.Post("", handler.MMedicalFacilityCreate)
	api.Put("", handler.MMedicalFacilityUpdate)
	api.Get(":id", handler.MMedicalFacilityIndex)
	api.Get("", handler.MMedicalFacilityPage)
	api.Delete(":id", handler.MMedicalFacilityDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-medical-facility/:size", handler.GenerateMMedicalFacility)
}

