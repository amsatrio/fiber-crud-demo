package m_biodata

import (
	"encoding/json"
	"errors"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/util"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MBiodataHandler struct {
	service  MBiodataService
	validate *validator.Validate
}

func NewMBiodataHandler(service MBiodataService, validate *validator.Validate) *MBiodataHandler {
	return &MBiodataHandler{
		service:  service,
		validate: validate,
	}
}

// MBiodataCreate godoc
//
//	@Summary		MBiodataCreate
//	@Description	Create MBiodata
//	@Tags			mBiodata
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			mBiodata	body		MBiodataRequest	true	"Add MBiodataRequest"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-biodata [post]
func (h *MBiodataHandler) MBiodataCreate(c *fiber.Ctx) error {

	res := &response.Response{}
	payload := new(MBiodataRequest)

	// parse payload
	if err := c.BodyParser(payload); err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse json error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	// validate payload
	if err := h.validate.Struct(payload); err != nil {
		out, _ := util.ValidateError(err)
		if out != nil {
			res.ErrMessagePayload(c.Path(), fiber.StatusBadRequest, "invalid payload", out)
			return c.Status(res.Status).JSON(res)
		}
	}

	image, err := c.FormFile("image")
	if err == nil && image != nil {
		payload.Image = image
	}

	// insert data
	err = h.service.Create(payload, 0)
	if err != nil {
		util.Log("ERROR", "controllers", "MBiodataCreate", "create data error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "create data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MBiodataUpdate godoc
//
//	@Summary		MBiodataUpdate
//	@Description	Update MBiodata
//	@Tags			mBiodata
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			mBiodata	body		MBiodataRequest	true	"Add MBiodataRequest"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-biodata [put]
func (h *MBiodataHandler) MBiodataUpdate(c *fiber.Ctx) error {

	res := &response.Response{}
	payload := new(MBiodataRequest)

	// parse payload
	if err := c.BodyParser(payload); err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse json error")
		return c.Status(res.Status).JSON(res)
	}

	// h.validate payload
	if err := h.validate.Struct(payload); err != nil {
		out, _ := util.ValidateError(err)
		if out != nil {
			res.ErrMessagePayload(c.Path(), fiber.StatusBadRequest, "invalid payload", out)
			return c.Status(res.Status).JSON(res)
		}
	}

	image, err := c.FormFile("image")
	if err == nil && image != nil {
		payload.Image = image
	}

	// update data
	err = h.service.Update(payload, 0)
	if err != nil {
		util.Log("ERROR", "controllers", "MBiodataUpdate", "update data error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "update data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MBiodataIndex godoc
//
//	@Summary		MBiodataIndex
//	@Description	Get MBiodata by id
//	@Tags			mBiodata
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			id	path		int	true	"MBiodata id"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-biodata/{id} [get]
func (h *MBiodataHandler) MBiodataIndex(c *fiber.Ctx) error {

	res := &response.Response{}

	id := c.Params("id")
	var idUint uint
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse json error")
		return c.Status(res.Status).JSON(res)
	}
	idUint = uint(idUint64)

	mBiodata, err := h.service.Get(idUint)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "data not found")
		return c.Status(res.Status).JSON(res)
	}

	if err != nil {
		util.Log("ERROR", "controllers", "MBiodataIndex", err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), mBiodata)
	return c.Status(res.Status).JSON(res)
}

// MBiodataDelete godoc
//
//	@Summary		MBiodataDelete
//	@Description	Delete MBiodata by id
//	@Tags			mBiodata
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			id	path		int	true	"MBiodata id"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-biodata/{id} [delete]
func (h *MBiodataHandler) MBiodataDelete(c *fiber.Ctx) error {
	res := &response.Response{}

	// get id from request param
	idParam := c.Params("id")
	var idUint uint
	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}
	idUint = uint(idUint64)

	// delete mBiodata
	err = h.service.Delete(idUint)

	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "delete data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	// return response
	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MBiodataPage godoc
//
//	@Summary		MBiodataPage
//	@Description	Get Page MBiodata
//	@Tags			mBiodata
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			page	query		string	false	"page" default(0)
//	@Param			size	query		string	false	"size" default(5)
//	@Param			sort	query		string	false	"sort"
//	@Param			filter	query		string	false	"filter"
//	@Param			search	query		string	false	"global filter"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-biodata [get]
func (h *MBiodataHandler) MBiodataPage(c *fiber.Ctx) error {
	res := &response.Response{}

	sortRequest := c.Query("sort", "[]")
	pageRequest := c.Query("page", "0")
	sizeRequest := c.Query("size", "5")
	filterRequest := c.Query("filter", "[]")
	searchRequest := c.Query("search", "")

	pageInt, errorPageInt := strconv.Atoi(pageRequest)
	sizeInt64, errorLimitInt64 := strconv.ParseInt(sizeRequest, 10, 64)
	sizeInt, errorLimitInt := strconv.Atoi(sizeRequest)

	if errorPageInt != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+errorPageInt.Error())
		return c.Status(res.Status).JSON(res)
	}
	if errorLimitInt64 != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+errorLimitInt64.Error())
		return c.Status(res.Status).JSON(res)
	}
	if errorLimitInt != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+errorLimitInt.Error())
		return c.Status(res.Status).JSON(res)
	}

	isLetterNumber := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`).MatchString
	if !isLetterNumber(searchRequest) && searchRequest != "" {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: global search must not contains special character")
		return c.Status(res.Status).JSON(res)
	}

	var sorts []request.Sort
	jsonUnmarshalErr := json.Unmarshal([]byte(sortRequest), &sorts)
	if jsonUnmarshalErr != nil {
		util.Log("ERROR", "controllers", "MBiodataPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}
	var filters []request.Filter
	jsonUnmarshalErr = json.Unmarshal([]byte(filterRequest), &filters)
	if jsonUnmarshalErr != nil {
		util.Log("ERROR", "controllers", "MBiodataPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}

	result, err := h.service.GetPage(
		sorts,
		filters,
		searchRequest,
		pageInt,
		sizeInt64,
		sizeInt)

	if err != nil {
		util.Log("ERROR", "controllers", "MBiodataPage", "error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	// return response
	res.Ok(c.Path(), result)
	return c.Status(res.Status).JSON(res)
}
