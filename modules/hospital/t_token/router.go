package t_token
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTTokenRepository(initializer.DB)
	service := NewTTokenService(repo)
	handler := NewTTokenHandler(service, validate)
	api := app.Group("/v1/t-token")
	api.Post("", handler.TTokenCreate)
	api.Put("", handler.TTokenUpdate)
	api.Get(":id", handler.TTokenIndex)
	api.Get("", handler.TTokenPage)
	api.Delete(":id", handler.TTokenDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-token/:size", handler.GenerateTToken)
}

