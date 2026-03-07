package m_medical_facility_category
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMMedicalFacilityCategoryRepository(initializer.DB)
	service := NewMMedicalFacilityCategoryService(repo)
	handler := NewMMedicalFacilityCategoryHandler(service, validate)
	api := app.Group("/v1/m-medical-facility-category")
	api.Post("", handler.MMedicalFacilityCategoryCreate)
	api.Put("", handler.MMedicalFacilityCategoryUpdate)
	api.Get(":id", handler.MMedicalFacilityCategoryIndex)
	api.Get("", handler.MMedicalFacilityCategoryPage)
	api.Delete(":id", handler.MMedicalFacilityCategoryDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-medical-facility-category/:size", handler.GenerateMMedicalFacilityCategory)
}

