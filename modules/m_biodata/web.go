package m_biodata

import (
	"github.com/gofiber/fiber/v2"
)

type MBiodataWebHandler struct {
}

func NewMBiodataWebHandler() *MBiodataWebHandler {
	return &MBiodataWebHandler{}
}

func (h *MBiodataWebHandler) MBiodataWebIndex(c *fiber.Ctx) error {
	return c.Render("pages/m-biodata", fiber.Map{
		"Title":      "Biodata CRUD",
		"ModalTitle": "Biodata Details",
	}, "layouts/main")
}
