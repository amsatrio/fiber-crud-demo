package m_biodata_attachment
import (
	"fiber-crud-demo/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(app *fiber.App, validate *validator.Validate) {
	repo := NewMBiodataAttachmentRepository(initializer.DB)
	service := NewMBiodataAttachmentService(repo)
	handler := NewMBiodataAttachmentHandler(service, validate)
	api := app.Group("/v1/m-biodata-attachment")
	api.Post("", handler.MBiodataAttachmentCreate)
	api.Put("", handler.MBiodataAttachmentUpdate)
	api.Get(":id", handler.MBiodataAttachmentIndex)
	api.Get("", handler.MBiodataAttachmentPage)
	api.Delete(":id", handler.MBiodataAttachmentDelete)

	//generatorApi := app.Group("/v1/generator")
	//generatorApi.Get("/m-biodata-attachment/:size", handler.GenerateMBiodataAttachment)
}

