package m_medical_facility_schedule
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMMedicalFacilityScheduleRepository(initializer.DB)
	service := NewMMedicalFacilityScheduleService(repo)
	handler := NewMMedicalFacilityScheduleHandler(service, validate)
	api := app.Group("/v1/m-medical-facility-schedule")
	api.Post("", handler.MMedicalFacilityScheduleCreate)
	api.Put("", handler.MMedicalFacilityScheduleUpdate)
	api.Get(":id", handler.MMedicalFacilityScheduleIndex)
	api.Get("", handler.MMedicalFacilitySchedulePage)
	api.Delete(":id", handler.MMedicalFacilityScheduleDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-medical-facility-schedule/:size", handler.GenerateMMedicalFacilitySchedule)
}

