package m_menu
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMMenuRepository(initializer.DB)
	service := NewMMenuService(repo)
	handler := NewMMenuHandler(service, validate)
	api := app.Group("/v1/m-menu")
	api.Post("", handler.MMenuCreate)
	api.Put("", handler.MMenuUpdate)
	api.Get(":id", handler.MMenuIndex)
	api.Get("", handler.MMenuPage)
	api.Delete(":id", handler.MMenuDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-menu/:size", handler.GenerateMMenu)
}

