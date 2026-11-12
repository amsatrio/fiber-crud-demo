package post

import (
	"errors"
	"strconv"
	"sync"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/util"
	"gorm.io/gorm"
)

type PostRepository interface {
	Get(id uint) (*Post, error)
	Create(data *Post) error
	Update(data *Post) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type PostRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{
		db: db,
	}
}

func (s *PostRepositoryImpl) Get(id uint) (*Post, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	entity := Post{}
	result := s.db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (s *PostRepositoryImpl) Create(entity *Post) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *PostRepositoryImpl) Update(entity *Post) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Model(&entity).Updates(entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *PostRepositoryImpl) Delete(id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var entity Post
	result := s.db.Delete(&entity, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *PostRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "repository", "GetPagePost", "")

	var entities []Post
	var entity Post
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
		Content:          ToPostResponsesParallel(entities),
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

	util.Log("INFO", "repository", "GetPagePost", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
