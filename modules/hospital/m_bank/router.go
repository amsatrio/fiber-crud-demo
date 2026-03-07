package m_bank
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMBankRepository(initializer.DB)
	service := NewMBankService(repo)
	handler := NewMBankHandler(service, validate)
	api := app.Group("/v1/m-bank")
	api.Post("", handler.MBankCreate)
	api.Put("", handler.MBankUpdate)
	api.Get(":id", handler.MBankIndex)
	api.Get("", handler.MBankPage)
	api.Delete(":id", handler.MBankDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-bank/:size", handler.GenerateMBank)
}

