package nomen

import (
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/initializer"
	"github.com/amsatrio/fiber-crud-demo/app/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var validate = validator.New()

func NomenPage(c fiber.Ctx) error {
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
		util.Log("ERROR", "controllers", "NomenPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}
	var filters []request.Filter
	jsonUnmarshalErr = json.Unmarshal([]byte(filterRequest), &filters)
	if jsonUnmarshalErr != nil {
		util.Log("ERROR", "controllers", "NomenPage", "jsonUnmarshalErr error: "+jsonUnmarshalErr.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse data error: "+jsonUnmarshalErr.Error())
		return c.Status(res.Status).JSON(res)
	}

	for i := range sorts {
		sorts[i].Id = util.CamelCaseToSnakeCase(sorts[i].Id)
	}
	for i := range filters {
		filters[i].Id = util.CamelCaseToSnakeCase(filters[i].Id)
	}

	var nouns []Noun
	db := initializer.DB_GERMAN.Model(&Noun{})

	if searchRequest != "" {
		db = util.ApplyGlobalSearch(db, searchRequest, util.GetJSONFieldTypes(Noun{}))
	}

	if len(filters) > 0 {
		db = util.ApplyFiltering(db, filters)
	}

	var totalElements int64
	filterDB_GERMAN := db.Session(&gorm.Session{})
	filterDB_GERMAN.Count(&totalElements)

	if len(sorts) > 0 {
		db = util.ApplySorting(db, sorts)
	}

	totalPages := totalElements / sizeInt64
	if totalElements%sizeInt64 != 0 {
		totalPages++
	}

	db.Scopes(util.ApplyPaginate(pageInt, sizeInt)).Find(&nouns)

	lastPage := int64(pageInt) == totalPages-1
	firstPage := pageInt == 0

	sort := response.Sort{
		Empty:    totalElements <= 0,
		Sorted:   len(sorts) > 0,
		Unsorted: len(sorts) == 0,
	}

	pageable := response.Pageable{
		Offset:     pageInt * sizeInt,
		PageNumber: pageInt,
		PageSize:   sizeInt,
		Paged:      true,
		UnPaged:    false,
		Sort:       sort,
	}

	pageResponse := response.Page{
		Content:          nouns,
		Pageable:         pageable,
		Sort:             sort,
		TotalPages:       totalPages,
		TotalElements:    totalElements,
		Size:             sizeInt,
		Number:           pageInt,
		NumberOfElements: len(nouns),
		Last:             lastPage,
		First:            firstPage,
		Empty:            sort.Empty,
	}

	res.Ok(c.Path(), pageResponse)
	return c.Status(res.Status).JSON(res)
}

func GetByID(c fiber.Ctx) error {
	util.Log("INFO", "nomen", "api", "GetByID()")

	res := &response.Response{}
	id := c.Params("id")

	var noun Noun
	if err := initializer.DB_GERMAN.First(&noun, id).Error; err != nil {
		res.ErrMessage(c.Path(), fiber.StatusNotFound, "noun not found")
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), noun)
	return c.Status(res.Status).JSON(res)
}

func Create(c fiber.Ctx) error {
	util.Log("INFO", "nomen", "api", "Create()")

	res := &response.Response{}
	req := new(CreateNounRequest)

	if err := c.Bind().JSON(req); err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "invalid request body")
		return c.Status(res.Status).JSON(res)
	}

	if err := validate.Struct(req); err != nil {
		validationErrors, _ := util.ValidateError(err)
		res.Err(c.Path(), validationErrors, fiber.StatusBadRequest)
		return c.Status(res.Status).JSON(res)
	}

	noun := Noun{
		Singular:          req.Singular,
		Gender:            req.Gender,
		Plural:            req.Plural,
		GenitiveSingular:  req.GenitiveSingular,
		IsNDeklination:    req.IsNDeklination,
		TranslationEn:     req.TranslationEn,
		ExampleSentenceDe: req.ExampleSentenceDe,
		ExampleSentenceEn: req.ExampleSentenceEn,
		Level:             req.Level,
	}

	if err := initializer.DB_GERMAN.Create(&noun).Error; err != nil {
		res.Err(c.Path(), err.Error(), fiber.StatusInternalServerError)
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), noun)
	return c.Status(res.Status).JSON(res)
}

func Update(c fiber.Ctx) error {
	util.Log("INFO", "nomen", "api", "Update()")

	res := &response.Response{}
	id := c.Params("id")
	req := new(UpdateNounRequest)

	if err := c.Bind().JSON(req); err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "invalid request body")
		return c.Status(res.Status).JSON(res)
	}

	var noun Noun
	if err := initializer.DB_GERMAN.First(&noun, id).Error; err != nil {
		res.ErrMessage(c.Path(), fiber.StatusNotFound, "noun not found")
		return c.Status(res.Status).JSON(res)
	}

	updates := make(map[string]interface{})
	if req.Singular != "" {
		updates["singular"] = req.Singular
	}
	if req.Gender != "" {
		updates["gender"] = req.Gender
	}
	if req.Plural != nil {
		updates["plural"] = req.Plural
	}
	if req.GenitiveSingular != nil {
		updates["genitive_singular"] = req.GenitiveSingular
	}
	if req.IsNDeklination != nil {
		updates["is_n_deklination"] = *req.IsNDeklination
	}
	if req.TranslationEn != "" {
		updates["translation_en"] = req.TranslationEn
	}
	if req.ExampleSentenceDe != nil {
		updates["example_sentence_de"] = req.ExampleSentenceDe
	}
	if req.ExampleSentenceEn != nil {
		updates["example_sentence_en"] = req.ExampleSentenceEn
	}
	if req.Level != "" {
		updates["level"] = req.Level
	}

	if err := initializer.DB_GERMAN.Model(&noun).Updates(updates).Error; err != nil {
		res.Err(c.Path(), err.Error(), fiber.StatusInternalServerError)
		return c.Status(res.Status).JSON(res)
	}

	initializer.DB_GERMAN.First(&noun, id)

	res.Ok(c.Path(), noun)
	return c.Status(res.Status).JSON(res)
}

func Delete(c fiber.Ctx) error {
	util.Log("INFO", "nomen", "api", "Delete()")

	res := &response.Response{}
	id := c.Params("id")

	var noun Noun
	if err := initializer.DB_GERMAN.First(&noun, id).Error; err != nil {
		res.ErrMessage(c.Path(), fiber.StatusNotFound, "noun not found")
		return c.Status(res.Status).JSON(res)
	}

	if err := initializer.DB_GERMAN.Delete(&noun).Error; err != nil {
		res.Err(c.Path(), err.Error(), fiber.StatusInternalServerError)
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), "noun deleted successfully")
	return c.Status(res.Status).JSON(res)
}
