package t_courier_discount
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTCourierDiscountRepository(initializer.DB)
	service := NewTCourierDiscountService(repo)
	handler := NewTCourierDiscountHandler(service, validate)
	api := app.Group("/v1/t-courier-discount")
	api.Post("", handler.TCourierDiscountCreate)
	api.Put("", handler.TCourierDiscountUpdate)
	api.Get(":id", handler.TCourierDiscountIndex)
	api.Get("", handler.TCourierDiscountPage)
	api.Delete(":id", handler.TCourierDiscountDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-courier-discount/:size", handler.GenerateTCourierDiscount)
}

