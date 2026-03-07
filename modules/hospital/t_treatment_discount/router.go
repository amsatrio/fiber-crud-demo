package t_treatment_discount
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewTTreatmentDiscountRepository(initializer.DB)
	service := NewTTreatmentDiscountService(repo)
	handler := NewTTreatmentDiscountHandler(service, validate)
	api := app.Group("/v1/t-treatment-discount")
	api.Post("", handler.TTreatmentDiscountCreate)
	api.Put("", handler.TTreatmentDiscountUpdate)
	api.Get(":id", handler.TTreatmentDiscountIndex)
	api.Get("", handler.TTreatmentDiscountPage)
	api.Delete(":id", handler.TTreatmentDiscountDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/t-treatment-discount/:size", handler.GenerateTTreatmentDiscount)
}

