package m_admin
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMAdminRepository(initializer.DB)
	service := NewMAdminService(repo)
	handler := NewMAdminHandler(service, validate)
	api := app.Group("/v1/m-admin")
	api.Post("", handler.MAdminCreate)
	api.Put("", handler.MAdminUpdate)
	api.Get(":id", handler.MAdminIndex)
	api.Get("", handler.MAdminPage)
	api.Delete(":id", handler.MAdminDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-admin/:size", handler.GenerateMAdmin)
}

