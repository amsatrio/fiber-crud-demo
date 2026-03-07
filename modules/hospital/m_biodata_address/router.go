package m_biodata_address
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMBiodataAddressRepository(initializer.DB)
	service := NewMBiodataAddressService(repo)
	handler := NewMBiodataAddressHandler(service, validate)
	api := app.Group("/v1/m-biodata-address")
	api.Post("", handler.MBiodataAddressCreate)
	api.Put("", handler.MBiodataAddressUpdate)
	api.Get(":id", handler.MBiodataAddressIndex)
	api.Get("", handler.MBiodataAddressPage)
	api.Delete(":id", handler.MBiodataAddressDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-biodata-address/:size", handler.GenerateMBiodataAddress)
}

