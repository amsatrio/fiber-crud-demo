package m_biodata

import (
	"github.com/gofiber/fiber/v2"
)

type MBiodataWebHandler struct {
}

func NewMBiodataWebHandler() *MBiodataWebHandler {
	return &MBiodataWebHandler{}
}

func (h *MBiodataWebHandler) MBiodataDatatableWebIndex(c *fiber.Ctx) error {
	return c.Render("pages/m-biodata/datatable", fiber.Map{
		"Title":      "Biodata CRUD",
		"ModalTitle": "Biodata Details",
	}, "layouts/main")
}

func (h *MBiodataWebHandler) MBiodataTableWebIndex(c *fiber.Ctx) error {
	return c.Render("pages/m-biodata/table", fiber.Map{
		"Title":      "Biodata CRUD",
		"ModalTitle": "Biodata Details",
	}, "layouts/main")
}

func (h *MBiodataWebHandler) MBiodataWebIndex(c *fiber.Ctx) error {
	return c.Render("pages/m-biodata/index", fiber.Map{
		"Title":      "Biodata CRUD",
		"ModalTitle": "Biodata Details",
	}, "layouts/main")
}
