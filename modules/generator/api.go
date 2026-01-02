package generator

import (
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/modules/m_biodata"
	"fiber-crud-demo/util"
	"math/rand/v2"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type GeneratorHandler struct {
	m_biodata_service m_biodata.MBiodataService
}

func NewGeneratorHandler(m_biodata_service m_biodata.MBiodataService) *GeneratorHandler {
	return &GeneratorHandler{
		m_biodata_service: m_biodata_service,
	}
}

func generateWord(min, max int, isNumeric bool) string {
	length := min + rand.IntN(max-min+1)
	b := make([]byte, length)

	var charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if isNumeric {
		charset = "0123456789"
	}

	charsetLen := len(charset)
	for i := range b {
		b[i] = charset[rand.IntN(charsetLen)]
	}

	return string(b)
}

func generateNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func (h *GeneratorHandler) GenerateMBiodata(c *fiber.Ctx) error {
	util.Log("INFO", "generator", "api", "GenerateMBiodata()")

	res := &response.Response{}

	sizeParam := c.Params("size")
	sizeUint64, err := strconv.ParseUint(sizeParam, 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse path param error")
		return c.Status(res.Status).JSON(res)
	}

	for range sizeUint64 {
		isDelete := false
		m_biodata := new(m_biodata.MBiodataRequest)
		m_biodata.Fullname = generateWord(5, 10, false)
		m_biodata.MobilePhone = generateWord(10, 12, true)
		m_biodata.IsDelete = &isDelete

		h.m_biodata_service.Create(m_biodata, 0)
	}

	res.Ok(c.Path(), "ok")

	return c.Status(res.Status).JSON(res)
}
