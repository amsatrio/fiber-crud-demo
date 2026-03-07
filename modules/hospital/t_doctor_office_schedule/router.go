package t_doctor_office_schedule
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTDoctorOfficeScheduleRepository(initializer.DB)
	service := NewTDoctorOfficeScheduleService(repo)
	handler := NewTDoctorOfficeScheduleHandler(service, validate)
	api := app.Group("/v1/t-doctor-office-schedule")
	api.Post("", handler.TDoctorOfficeScheduleCreate)
	api.Put("", handler.TDoctorOfficeScheduleUpdate)
	api.Get(":id", handler.TDoctorOfficeScheduleIndex)
	api.Get("", handler.TDoctorOfficeSchedulePage)
	api.Delete(":id", handler.TDoctorOfficeScheduleDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-doctor-office-schedule/:size", handler.GenerateTDoctorOfficeSchedule)
}

