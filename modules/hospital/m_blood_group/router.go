package m_blood_group
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMBloodGroupRepository(initializer.DB)
	service := NewMBloodGroupService(repo)
	handler := NewMBloodGroupHandler(service, validate)
	api := app.Group("/v1/m-blood-group")
	api.Post("", handler.MBloodGroupCreate)
	api.Put("", handler.MBloodGroupUpdate)
	api.Get(":id", handler.MBloodGroupIndex)
	api.Get("", handler.MBloodGroupPage)
	api.Delete(":id", handler.MBloodGroupDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-blood-group/:size", handler.GenerateMBloodGroup)
}

