package m_menu_role
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMMenuRoleRepository(initializer.DB)
	service := NewMMenuRoleService(repo)
	handler := NewMMenuRoleHandler(service, validate)
	api := app.Group("/v1/m-menu-role")
	api.Post("", handler.MMenuRoleCreate)
	api.Put("", handler.MMenuRoleUpdate)
	api.Get(":id", handler.MMenuRoleIndex)
	api.Get("", handler.MMenuRolePage)
	api.Delete(":id", handler.MMenuRoleDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-menu-role/:size", handler.GenerateMMenuRole)
}

