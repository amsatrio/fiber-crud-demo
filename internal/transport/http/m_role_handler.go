package http

import (
	"encoding/json"
	"errors"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/internal/application"
	"fiber-crud-demo/internal/domain"
	"fiber-crud-demo/util"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MRoleHandler struct {
	service  application.MRoleService
	validate *validator.Validate
}

func NewMRoleHandler(service application.MRoleService, validate *validator.Validate) *MRoleHandler {
	return &MRoleHandler{
		service:  service,
		validate: validate,
	}
}

// MRoleCreate godoc
//
//	@Summary		MRoleCreate
//	@Description	Create MRole
//	@Tags			mRole
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			mRole	body		domain.MRoleRequest	true	"Add MRoleRequest"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-role [post]
func (h *MRoleHandler) MRoleCreate(c *fiber.Ctx) error {

	res := &response.Response{}
	payload := new(domain.MRoleRequest)

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

	// insert data

	err := h.service.Create(payload, 0)
	if err != nil {
		util.Log("ERROR", "controllers", "MRoleCreate", "create data error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "create data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MRoleUpdate godoc
//
//	@Summary		MRoleUpdate
//	@Description	Update MRole
//	@Tags			mRole
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			mRole	body		domain.MRoleRequest	true	"Add MRoleRequest"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-role [put]
func (h *MRoleHandler) MRoleUpdate(c *fiber.Ctx) error {

	res := &response.Response{}
	payload := new(domain.MRoleRequest)

	// parse payload
	if err := c.BodyParser(payload); err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse json error")
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

	// update data
	err := h.service.Update(payload, 0)
	if err != nil {
		util.Log("ERROR", "controllers", "MRoleUpdate", "update data error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "update data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MRoleIndex godoc
//
//	@Summary		MRoleIndex
//	@Description	Get MRole by id
//	@Tags			mRole
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			id	path		int	true	"MRole id"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-role/{id} [get]
func (h *MRoleHandler) MRoleIndex(c *fiber.Ctx) error {

	res := &response.Response{}

	id := c.Params("id")
	var idUint uint
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse json error")
		return c.Status(res.Status).JSON(res)
	}
	idUint = uint(idUint64)

	mRole, err := h.service.Get(idUint)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "data not found")
		return c.Status(res.Status).JSON(res)
	}

	if err != nil {
		util.Log("ERROR", "controllers", "MRoleIndex", err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), mRole)
	return c.Status(res.Status).JSON(res)
}

// MRoleDelete godoc
//
//	@Summary		MRoleDelete
//	@Description	Delete MRole by id
//	@Tags			mRole
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			id	path		int	true	"MRole id"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-role/{id} [delete]
func (h *MRoleHandler) MRoleDelete(c *fiber.Ctx) error {
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

	// delete mRole
	err = h.service.Delete(idUint)

	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "delete data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	// return response
	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MRolePage godoc
//
//	@Summary		MRolePage
//	@Description	Get Page MRole
//	@Tags			mRole
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
//	@Router			/m-role [get]
func (h *MRoleHandler) MRolePage(c *fiber.Ctx) error {
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
		util.Log("ERROR", "controllers", "MRolePage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}
	var filters []request.Filter
	jsonUnmarshalErr = json.Unmarshal([]byte(filterRequest), &filters)
	if jsonUnmarshalErr != nil {
		util.Log("ERROR", "controllers", "MRolePage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
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
		util.Log("ERROR", "controllers", "MRolePage", "error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	// return response
	res.Ok(c.Path(), result)
	return c.Status(res.Status).JSON(res)
}
