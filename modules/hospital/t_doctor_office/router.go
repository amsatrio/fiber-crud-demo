package t_doctor_office
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTDoctorOfficeRepository(initializer.DB)
	service := NewTDoctorOfficeService(repo)
	handler := NewTDoctorOfficeHandler(service, validate)
	api := app.Group("/v1/t-doctor-office")
	api.Post("", handler.TDoctorOfficeCreate)
	api.Put("", handler.TDoctorOfficeUpdate)
	api.Get(":id", handler.TDoctorOfficeIndex)
	api.Get("", handler.TDoctorOfficePage)
	api.Delete(":id", handler.TDoctorOfficeDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-doctor-office/:size", handler.GenerateTDoctorOffice)
}

