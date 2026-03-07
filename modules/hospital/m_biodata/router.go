package m_biodata
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMBiodataRepository(initializer.DB)
	service := NewMBiodataService(repo)
	handler := NewMBiodataHandler(service, validate)
	api := app.Group("/v1/m-biodata")
	api.Post("", handler.MBiodataCreate)
	api.Put("", handler.MBiodataUpdate)
	api.Get(":id", handler.MBiodataIndex)
	api.Get("", handler.MBiodataPage)
	api.Delete(":id", handler.MBiodataDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-biodata/:size", handler.GenerateMBiodata)
}

