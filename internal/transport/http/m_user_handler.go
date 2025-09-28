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

type MUserHandler struct {
	service  application.MUserService
	validate *validator.Validate
}

func NewMUserHandler(service application.MUserService, validate *validator.Validate) *MUserHandler {
	return &MUserHandler{
		service:  service,
		validate: validate,
	}
}

// var validate = validator.New()

// MUserCreate godoc
//
//	@Summary		MUserCreate
//	@Description	Create MUser
//	@Tags			mUser
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			mUser	body		domain.MUserRequest	true	"Add MUserRequest"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-user [post]
func (h *MUserHandler) MUserCreate(c *fiber.Ctx) error {

	res := &response.Response{}
	payload := new(domain.MUserRequest)

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
		util.Log("ERROR", "controllers", "MUserCreate", "create data error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "create data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MUserUpdate godoc
//
//	@Summary		MUserUpdate
//	@Description	Update MUser
//	@Tags			mUser
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			mUser	body		domain.MUserRequest	true	"Add MUserRequest"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-user [put]
func (h *MUserHandler) MUserUpdate(c *fiber.Ctx) error {

	res := &response.Response{}
	payload := new(domain.MUserRequest)

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
		util.Log("ERROR", "controllers", "MUserUpdate", "update data error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "update data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MUserIndex godoc
//
//	@Summary		MUserIndex
//	@Description	Get MUser by id
//	@Tags			mUser
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			id	path		int	true	"MUser id"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-user/{id} [get]
func (h *MUserHandler) MUserIndex(c *fiber.Ctx) error {

	res := &response.Response{}

	id := c.Params("id")
	var idUint uint
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse json error")
		return c.Status(res.Status).JSON(res)
	}
	idUint = uint(idUint64)

	mUser, err := h.service.Get(idUint)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "data not found")
		return c.Status(res.Status).JSON(res)
	}

	if err != nil {
		util.Log("ERROR", "controllers", "MUserIndex", err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), mUser)
	return c.Status(res.Status).JSON(res)
}

// MUserDelete godoc
//
//	@Summary		MUserDelete
//	@Description	Delete MUser by id
//	@Tags			mUser
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			id	path		int	true	"MUser id"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-user/{id} [delete]
func (h *MUserHandler) MUserDelete(c *fiber.Ctx) error {
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

	// delete mUser
	err = h.service.Delete(idUint)

	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "delete data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	// return response
	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MUserPage godoc
//
//	@Summary		MUserPage
//	@Description	Get Page MUser
//	@Tags			mUser
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Encoding	header	string	false	"gzip" default(gzip)
//	@Param			_page	query		string	false	"page" default(0)
//	@Param			_size	query		string	false	"size" default(5)
//	@Param			_sort	query		string	false	"sort"
//	@Param			_filter	query		string	false	"filter"
//	@Param			_q	query		string	false	"global filter"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/m-user [get]
func (h *MUserHandler) MUserPage(c *fiber.Ctx) error {
	res := &response.Response{}

	sortRequest := c.Query("_sort", "[]")
	pageRequest := c.Query("_page", "0")
	sizeRequest := c.Query("_size", "10")
	filterRequest := c.Query("_filter", "[]")
	searchRequest := c.Query("_q", "")

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
		util.Log("ERROR", "controllers", "MUserPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}
	var filters []request.Filter
	jsonUnmarshalErr = json.Unmarshal([]byte(filterRequest), &filters)
	if jsonUnmarshalErr != nil {
		util.Log("ERROR", "controllers", "MUserPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
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
		util.Log("ERROR", "controllers", "MUserPage", "error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	// return response
	res.Ok(c.Path(), result)
	return c.Status(res.Status).JSON(res)
}
