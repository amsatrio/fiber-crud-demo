package like

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

type LikeHandler struct {
	service  LikeService
	validate *validator.Validate
}

func NewLikeHandler(service LikeService, validate *validator.Validate) *LikeHandler {
	return &LikeHandler{
		service:  service,
		validate: validate,
	}
}

// LikeCreate godoc
//
//	@Summary        LikeCreate
//	@Description    Create Like
//	@Tags           like
//	@Accept         json
//	@Produce        json
//	@Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//	@Param          like   body        LikeRequest    true    "Add LikeRequest"
//	@Success        200 {object}    response.Response
//	@Failure        400 {object}    response.Response
//	@Failure        404 {object}    response.Response
//	@Failure        500 {object}    response.Response
//	@Router         /v1/twitter-clone/like [post]
func (h *LikeHandler) LikeCreate(c fiber.Ctx) error {

	res := &response.Response{}
	payload := new(LikeRequest)

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

	err := h.service.Create(payload)
	if err != nil {
		util.Log("ERROR", "controllers", "LikeCreate", "create data error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "create data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// LikeIndex godoc
//
//	@Summary        LikeIndex
//	@Description    Get Like by user id and tweet id
//	@Tags           like
//	@Accept         json
//	@Produce        json
//	@Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//	@Param          userId  path        int true    "Like user id"
//	@Param          tweetId  path        int true    "Like tweet id"
//	@Success        200 {object}    response.Response
//	@Failure        400 {object}    response.Response
//	@Failure        404 {object}    response.Response
//	@Failure        500 {object}    response.Response
//	@Router         /v1/twitter-clone/like/{userId}/{tweetId} [get]
func (h *LikeHandler) LikeIndex(c fiber.Ctx) error {

	res := &response.Response{}

	userId, err := strconv.ParseUint(c.Params("userId"), 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse json error")
		return c.Status(res.Status).JSON(res)
	}
	tweetId, err := strconv.ParseUint(c.Params("tweetId"), 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse json error")
		return c.Status(res.Status).JSON(res)
	}

	entity, err := h.service.Get(uint(userId), uint(tweetId))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "data not found")
		return c.Status(res.Status).JSON(res)
	}

	if err != nil {
		util.Log("ERROR", "controllers", "LikeIndex", err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), entity)
	return c.Status(res.Status).JSON(res)
}

// LikeDelete godoc
//
//	@Summary        LikeDelete
//	@Description    Delete Like by user id and tweet id
//	@Tags           like
//	@Accept         json
//	@Produce        json
//	@Param          Accept-Encoding header  string  false   "gzip" default(gzip)
//	@Param          userId  path        int true    "Like user id"
//	@Param          tweetId  path        int true    "Like tweet id"
//	@Success        200 {object}    response.Response
//	@Failure        400 {object}    response.Response
//	@Failure        404 {object}    response.Response
//	@Failure        500 {object}    response.Response
//	@Router         /v1/twitter-clone/like/{userId}/{tweetId} [delete]
func (h *LikeHandler) LikeDelete(c fiber.Ctx) error {
	res := &response.Response{}

	userId, err := strconv.ParseUint(c.Params("userId"), 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}
	tweetId, err := strconv.ParseUint(c.Params("tweetId"), 10, 32)
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	err = h.service.Delete(uint(userId), uint(tweetId))
	if err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "delete data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

// LikePage godoc
//
//	@Summary        LikePage
//	@Description    Get Page Like
//	@Tags           like
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
//	@Router         /v1/twitter-clone/like [get]
func (h *LikeHandler) LikePage(c fiber.Ctx) error {
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
		util.Log("ERROR", "controllers", "LikePage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}
	var filters []request.Filter
	jsonUnmarshalErr = json.Unmarshal([]byte(filterRequest), &filters)
	if jsonUnmarshalErr != nil {
		util.Log("ERROR", "controllers", "LikePage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}

	for i := range sorts {
		sorts[i].Id = util.CamelCaseToSnakeCase(sorts[i].Id)
	}
	for i := range filters {
		filters[i].Id = util.CamelCaseToSnakeCase(filters[i].Id)
	}

	result, err := h.service.GetPage(sorts, filters, searchRequest, pageInt, sizeInt64, sizeInt)

	if err != nil {
		util.Log("ERROR", "controllers", "LikePage", "error: "+err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), result)
	return c.Status(res.Status).JSON(res)
}
