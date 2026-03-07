package m_medical_item_segmentation
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMMedicalItemSegmentationRepository(initializer.DB)
	service := NewMMedicalItemSegmentationService(repo)
	handler := NewMMedicalItemSegmentationHandler(service, validate)
	api := app.Group("/v1/m-medical-item-segmentation")
	api.Post("", handler.MMedicalItemSegmentationCreate)
	api.Put("", handler.MMedicalItemSegmentationUpdate)
	api.Get(":id", handler.MMedicalItemSegmentationIndex)
	api.Get("", handler.MMedicalItemSegmentationPage)
	api.Delete(":id", handler.MMedicalItemSegmentationDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-medical-item-segmentation/:size", handler.GenerateMMedicalItemSegmentation)
}

