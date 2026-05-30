package m_customer_member

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/util"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type MCustomerMemberHandler struct {
	service  MCustomerMemberService
	validate *validator.Validate
}

func NewMCustomerMemberHandler(service MCustomerMemberService, validate *validator.Validate) *MCustomerMemberHandler {
	return &MCustomerMemberHandler{
		service:  service,
		validate: validate,
	}
}

// MCustomerMemberCreate godoc
//
//	@Summary        MCustomerMemberCreate
//	@Description    Create MCustomerMember
//	@Tags           mCustomerMember
//	@Accept         json
//	@Produce        json
//	@Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//	@Param          mCustomerMember   body        MCustomerMemberRequest    true    "Add MCustomerMemberRequest"
//	@Success        200 {object}    response.Response
//	@Failure        400 {object}    response.Response
//	@Failure        404 {object}    response.Response
//	@Failure        500 {object}    response.Response
//	@Router         /v1/m-customer-member [post]
func (h *MCustomerMemberHandler) MCustomerMemberCreate(c fiber.Ctx) error {

	res := &response.Response{}
	payload := new(MCustomerMemberRequest)

	if err := c.Bind().Body(payload); err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse body error: "+err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if err := h.validate.Struct(payload); err != nil {
		out, _ := util.ValidateError(err)
		if out != nil {
			res.ErrMessagePayload(c.Path(), fiber.StatusBadRequest, "invalid payload", out)
			return c.Status(res.Status).JSON(res)
		}
	}

	err := h.service.Create(payload, 0)
	if err != nil {
		util.Log("ERROR", "controllers", "MCustomerMemberCreate", "create data error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "create data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MCustomerMemberUpdate godoc
//
//	@Summary        MCustomerMemberUpdate
//	@Description    Update MCustomerMember
//	@Tags           mCustomerMember
//	@Accept         json
//	@Produce        json
//	@Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//	@Param          mCustomerMember   body        MCustomerMemberRequest    true    "Update MCustomerMemberRequest"
//	@Success        200 {object}    response.Response
//	@Failure        400 {object}    response.Response
//	@Failure        404 {object}    response.Response
//	@Failure        500 {object}    response.Response
//	@Router         /v1/m-customer-member [put]
func (h *MCustomerMemberHandler) MCustomerMemberUpdate(c fiber.Ctx) error {

	res := &response.Response{}
	payload := new(MCustomerMemberRequest)

	if err := c.Bind().Body(payload); err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse body error: "+err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if err := h.validate.Struct(payload); err != nil {
		out, _ := util.ValidateError(err)
		if out != nil {
			res.ErrMessagePayload(c.Path(), fiber.StatusBadRequest, "invalid payload", out)
			return c.Status(res.Status).JSON(res)
		}
	}

	err := h.service.Update(payload, 0)
	if err != nil {
		util.Log("ERROR", "controllers", "MCustomerMemberUpdate", "update data error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "update data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MCustomerMemberIndex godoc
//
//	@Summary        MCustomerMemberIndex
//	@Description    Get MCustomerMember by id
//	@Tags           mCustomerMember
//	@Accept         json
//	@Produce        json
//	@Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//	@Param          id  path        int true    "MCustomerMember id"
//	@Success        200 {object}    response.Response
//	@Failure        400 {object}    response.Response
//	@Failure        404 {object}    response.Response
//	@Failure        500 {object}    response.Response
//	@Router         /v1/m-customer-member/{id} [get]
func (h *MCustomerMemberHandler) MCustomerMemberIndex(c fiber.Ctx) error {

	res := &response.Response{}

	id := c.Params("id")
	var idUint uint
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse json error")
		return c.Status(res.Status).JSON(res)
	}
	idUint = uint(idUint64)

	entity, err := h.service.Get(idUint)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "data not found")
		return c.Status(res.Status).JSON(res)
	}

	if err != nil {
		util.Log("ERROR", "controllers", "MCustomerMemberIndex", err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), entity)
	return c.Status(res.Status).JSON(res)
}

// MCustomerMemberDelete godoc
//
//	@Summary        MCustomerMemberDelete
//	@Description    Delete MCustomerMember by id
//	@Tags           mCustomerMember
//	@Accept         json
//	@Produce        json
//	@Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//	@Param          id  path        int true    "MCustomerMember id"
//	@Success        200 {object}    response.Response
//	@Failure        400 {object}    response.Response
//	@Failure        404 {object}    response.Response
//	@Failure        500 {object}    response.Response
//	@Router         /v1/m-customer-member/{id} [delete]
func (h *MCustomerMemberHandler) MCustomerMemberDelete(c fiber.Ctx) error {
	res := &response.Response{}

	idParam := c.Params("id")
	var idUint uint
	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}
	idUint = uint(idUint64)

	err = h.service.Delete(idUint)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "delete data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// MCustomerMemberPage godoc
//
//	@Summary        MCustomerMemberPage
//	@Description    Get Page MCustomerMember
//	@Tags           mCustomerMember
//	@Accept         json
//	@Produce        json
//	@Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//	@Param          page    query       string  false   "page" default(0)
//	@Param          size    query       string  false   "size" default(5)
//	@Param          sort    query       string  false   "sort"
//	@Param          filter  query       string  false   "filter"
//	@Param          search  query       string  false   "global filter"
//	@Success        200 {object}    response.Response
//	@Failure        400 {object}    response.Response
//	@Failure        404 {object}    response.Response
//	@Failure        500 {object}    response.Response
//	@Router         /v1/m-customer-member [get]
func (h *MCustomerMemberHandler) MCustomerMemberPage(c fiber.Ctx) error {
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
		util.Log("ERROR", "controllers", "MCustomerMemberPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}
	var filters []request.Filter
	jsonUnmarshalErr = json.Unmarshal([]byte(filterRequest), &filters)
	if jsonUnmarshalErr != nil {
		util.Log("ERROR", "controllers", "MCustomerMemberPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}

	result, err := h.service.GetPage(sorts, filters, searchRequest, pageInt, sizeInt64, sizeInt)

	if err != nil {
		util.Log("ERROR", "controllers", "MCustomerMemberPage", "error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), result)
	return c.Status(res.Status).JSON(res)
}
