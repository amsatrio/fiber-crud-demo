package t_reset_password
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTResetPasswordRepository(initializer.DB)
	service := NewTResetPasswordService(repo)
	handler := NewTResetPasswordHandler(service, validate)
	api := app.Group("/v1/t-reset-password")
	api.Post("", handler.TResetPasswordCreate)
	api.Put("", handler.TResetPasswordUpdate)
	api.Get(":id", handler.TResetPasswordIndex)
	api.Get("", handler.TResetPasswordPage)
	api.Delete(":id", handler.TResetPasswordDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-reset-password/:size", handler.GenerateTResetPassword)
}

