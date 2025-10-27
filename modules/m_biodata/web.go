package m_biodata

import (
	"github.com/gofiber/fiber/v2"
)

type MBiodataWebHandler struct {
}

func NewMBiodataWebHandler() *MBiodataWebHandler {
	return &MBiodataWebHandler{}
}

func (h *MBiodataWebHandler) MBiodataTableDatatableWebIndex(c *fiber.Ctx) error {
	return c.Render("pages/m-biodata/table-datatable", fiber.Map{
		"Title":      "Biodata CRUD",
		"ModalTitle": "Biodata Details",
	}, "layouts/main")
}

func (h *MBiodataWebHandler) MBiodataTableHTMLWebIndex(c *fiber.Ctx) error {
	return c.Render("pages/m-biodata/table-html", fiber.Map{
		"Title":      "Biodata CRUD",
		"ModalTitle": "Biodata Details",
	}, "layouts/main")
}

func (h *MBiodataWebHandler) MBiodataTableTailwindCSSWebIndex(c *fiber.Ctx) error {
	return c.Render("pages/m-biodata/table-tailwindcss", fiber.Map{
		"Title":      "Biodata CRUD",
		"ModalTitle": "Biodata Details",
	}, "layouts/main")
}

func (h *MBiodataWebHandler) MBiodataTableBootstrapWebIndex(c *fiber.Ctx) error {
	return c.Render("pages/m-biodata/table-bootstrap", fiber.Map{
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
