package m_user
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMUserRepository(initializer.DB)
	service := NewMUserService(repo)
	handler := NewMUserHandler(service, validate)
	api := app.Group("/v1/m-user")
	api.Post("", handler.MUserCreate)
	api.Put("", handler.MUserUpdate)
	api.Get(":id", handler.MUserIndex)
	api.Get("", handler.MUserPage)
	api.Delete(":id", handler.MUserDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-user/:size", handler.GenerateMUser)
}

