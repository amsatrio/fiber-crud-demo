package t_treatment_discount

import (
    "encoding/json"
    "errors"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    "fiber-crud-demo/util"
    "regexp"
    "strconv"

    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v3"
    "gorm.io/gorm"
)

type TTreatmentDiscountHandler struct {
    service  TTreatmentDiscountService
    validate *validator.Validate
}

func NewTTreatmentDiscountHandler(service TTreatmentDiscountService, validate *validator.Validate) *TTreatmentDiscountHandler {
    return &TTreatmentDiscountHandler{
        service:  service,
        validate: validate,
    }
}

// TTreatmentDiscountCreate godoc
//
//  @Summary        TTreatmentDiscountCreate
//  @Description    Create TTreatmentDiscount
//  @Tags           tTreatmentDiscount
//  @Accept         json
//  @Produce        json
//  @Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//  @Param          tTreatmentDiscount   body        TTreatmentDiscountRequest    true    "Add TTreatmentDiscountRequest"
//  @Success        200 {object}    response.Response
//  @Failure        400 {object}    response.Response
//  @Failure        404 {object}    response.Response
//  @Failure        500 {object}    response.Response
//  @Router         /v1/t-treatment-discount [post]
func (h *TTreatmentDiscountHandler) TTreatmentDiscountCreate(c fiber.Ctx) error {

    res := &response.Response{}
    payload := new(TTreatmentDiscountRequest)

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
        util.Log("ERROR", "controllers", "TTreatmentDiscountCreate", "create data error: "+err.Error())
        res.ErrMessage(c.Path(), fiber.StatusBadRequest, "create data error: "+err.Error())
        return c.Status(res.Status).JSON(res)
    }

    res.Ok(c.Path(), nil)
    return c.Status(res.Status).JSON(res)
}

// TTreatmentDiscountUpdate godoc
//
//  @Summary        TTreatmentDiscountUpdate
//  @Description    Update TTreatmentDiscount
//  @Tags           tTreatmentDiscount
//  @Accept         json
//  @Produce        json
//  @Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//  @Param          tTreatmentDiscount   body        TTreatmentDiscountRequest    true    "Update TTreatmentDiscountRequest"
//  @Success        200 {object}    response.Response
//  @Failure        400 {object}    response.Response
//  @Failure        404 {object}    response.Response
//  @Failure        500 {object}    response.Response
//  @Router         /v1/t-treatment-discount [put]
func (h *TTreatmentDiscountHandler) TTreatmentDiscountUpdate(c fiber.Ctx) error {

    res := &response.Response{}
    payload := new(TTreatmentDiscountRequest)

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
        util.Log("ERROR", "controllers", "TTreatmentDiscountUpdate", "update data error: "+err.Error())
        res.ErrMessage(c.Path(), fiber.StatusBadRequest, "update data error: "+err.Error())
        return c.Status(res.Status).JSON(res)
    }

    res.Ok(c.Path(), nil)
    return c.Status(res.Status).JSON(res)
}

// TTreatmentDiscountIndex godoc
//
//  @Summary        TTreatmentDiscountIndex
//  @Description    Get TTreatmentDiscount by id
//  @Tags           tTreatmentDiscount
//  @Accept         json
//  @Produce        json
//  @Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//  @Param          id  path        int true    "TTreatmentDiscount id"
//  @Success        200 {object}    response.Response
//  @Failure        400 {object}    response.Response
//  @Failure        404 {object}    response.Response
//  @Failure        500 {object}    response.Response
//  @Router         /v1/t-treatment-discount/{id} [get]
func (h *TTreatmentDiscountHandler) TTreatmentDiscountIndex(c fiber.Ctx) error {

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
        util.Log("ERROR", "controllers", "TTreatmentDiscountIndex", err.Error())
        res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
        return c.Status(res.Status).JSON(res)
    }

    res.Ok(c.Path(), entity)
    return c.Status(res.Status).JSON(res)
}

// TTreatmentDiscountDelete godoc
//
//  @Summary        TTreatmentDiscountDelete
//  @Description    Delete TTreatmentDiscount by id
//  @Tags           tTreatmentDiscount
//  @Accept         json
//  @Produce        json
//  @Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//  @Param          id  path        int true    "TTreatmentDiscount id"
//  @Success        200 {object}    response.Response
//  @Failure        400 {object}    response.Response
//  @Failure        404 {object}    response.Response
//  @Failure        500 {object}    response.Response
//  @Router         /v1/t-treatment-discount/{id} [delete]
func (h *TTreatmentDiscountHandler) TTreatmentDiscountDelete(c fiber.Ctx) error {
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

// TTreatmentDiscountPage godoc
//
//  @Summary        TTreatmentDiscountPage
//  @Description    Get Page TTreatmentDiscount
//  @Tags           tTreatmentDiscount
//  @Accept         json
//  @Produce        json
//  @Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//  @Param          page    query       string  false   "page" default(0)
//  @Param          size    query       string  false   "size" default(5)
//  @Param          sort    query       string  false   "sort"
//  @Param          filter  query       string  false   "filter"
//  @Param          search  query       string  false   "global filter"
//  @Success        200 {object}    response.Response
//  @Failure        400 {object}    response.Response
//  @Failure        404 {object}    response.Response
//  @Failure        500 {object}    response.Response
//  @Router         /v1/t-treatment-discount [get]
func (h *TTreatmentDiscountHandler) TTreatmentDiscountPage(c fiber.Ctx) error {
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
        util.Log("ERROR", "controllers", "TTreatmentDiscountPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
        res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
        return c.Status(res.Status).JSON(res)
    }
    var filters []request.Filter
    jsonUnmarshalErr = json.Unmarshal([]byte(filterRequest), &filters)
    if jsonUnmarshalErr != nil {
        util.Log("ERROR", "controllers", "TTreatmentDiscountPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
        res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
        return c.Status(res.Status).JSON(res)
    }

    result, err := h.service.GetPage(sorts, filters, searchRequest, pageInt, sizeInt64, sizeInt)

    if err != nil {
        util.Log("ERROR", "controllers", "TTreatmentDiscountPage", "error: "+err.Error())
        res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
        return c.Status(res.Status).JSON(res)
    }

    res.Ok(c.Path(), result)
    return c.Status(res.Status).JSON(res)
}
