package m_role
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMRoleRepository(initializer.DB)
	service := NewMRoleService(repo)
	handler := NewMRoleHandler(service, validate)
	api := app.Group("/v1/m-role")
	api.Post("", handler.MRoleCreate)
	api.Put("", handler.MRoleUpdate)
	api.Get(":id", handler.MRoleIndex)
	api.Get("", handler.MRolePage)
	api.Delete(":id", handler.MRoleDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-role/:size", handler.GenerateMRole)
}

