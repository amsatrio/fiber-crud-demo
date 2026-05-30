package m_biodata_attachment

import (
	"github.com/amsatrio/fiber-crud-demo/app/initializer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetRouter(api fiber.Router, validate *validator.Validate) {
	repo := NewMBiodataAttachmentRepository(initializer.DB)
	service := NewMBiodataAttachmentService(repo)
	handler := NewMBiodataAttachmentHandler(service, validate)

	api.Post("/m-biodata-attachment", handler.MBiodataAttachmentCreate)
	api.Put("/m-biodata-attachment", handler.MBiodataAttachmentUpdate)
	api.Get("/m-biodata-attachment/:id", handler.MBiodataAttachmentIndex)
	api.Get("/m-biodata-attachment", handler.MBiodataAttachmentPage)
	api.Delete("/m-biodata-attachment/:id", handler.MBiodataAttachmentDelete)

	//api.Get("/generator/m-biodata-attachment/:size", handler.GenerateMBiodataAttachment)
}
