package m_education_level

import (
	"errors"
	"strconv"
	"sync"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/util"

	"gorm.io/gorm"
)

type MEducationLevelRepository interface {
	Get(id uint) (*MEducationLevel, error)
	Create(data *MEducationLevel) error
	Update(data *MEducationLevel) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MEducationLevelRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewMEducationLevelRepository(db *gorm.DB) MEducationLevelRepository {
	return &MEducationLevelRepositoryImpl{
		db: db,
	}
}

func (s *MEducationLevelRepositoryImpl) Get(id uint) (*MEducationLevel, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	entity := MEducationLevel{}
	result := s.db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (s *MEducationLevelRepositoryImpl) Create(entity *MEducationLevel) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MEducationLevelRepositoryImpl) Update(entity *MEducationLevel) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Model(&entity).Updates(entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MEducationLevelRepositoryImpl) Delete(id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var entity MEducationLevel
	result := s.db.Delete(&entity, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *MEducationLevelRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "repository", "GetPageMEducationLevel", "")

	var entities []MEducationLevel
	var entity MEducationLevel
	fieldMap := util.GetJSONFieldTypes(entity)

	db := s.db
	db = util.ApplySorting(db, sortRequest)
	db = util.ApplyFiltering(db, filterRequest)
	db = util.ApplyGlobalSearch(db, searchRequest, fieldMap)

	totalElements := db.Find(&entities).RowsAffected

	totalPages := totalElements / sizeInt64
	if totalElements%sizeInt64 != 0 {
		totalPages++
	}

	result := db.Scopes(util.ApplyPaginate(pageInt, sizeInt)).Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	lastPage := int64(pageInt) == totalPages-1
	firstPage := pageInt == 0

	sort := response.Sort{
		Empty:    totalElements <= 0,
		Sorted:   true,
		Unsorted: false,
	}

	pageable := response.Pageable{
		Offset:     pageInt * sizeInt,
		PageNumber: pageInt,
		PageSize:   sizeInt,
		Paged:      true,
		UnPaged:    false,
		Sort:       sort,
	}

	page := response.Page{
		Content:          ToMEducationLevelResponsesParallel(entities),
		Pageable:         pageable,
		Sort:             sort,
		TotalPages:       totalPages,
		TotalElements:    totalElements,
		Size:             sizeInt,
		Number:           pageInt,
		NumberOfElements: sizeInt,
		Last:             lastPage,
		First:            firstPage,
		Empty:            sort.Empty,
	}

	util.Log("INFO", "repository", "GetPageMEducationLevel", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
