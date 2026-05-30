package m_biodata_attachment

import (
	"errors"
	"strconv"
	"sync"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/util"

	"gorm.io/gorm"
)

type MBiodataAttachmentRepository interface {
	Get(id uint) (*MBiodataAttachment, error)
	Create(data *MBiodataAttachment) error
	Update(data *MBiodataAttachment) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MBiodataAttachmentRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewMBiodataAttachmentRepository(db *gorm.DB) MBiodataAttachmentRepository {
	return &MBiodataAttachmentRepositoryImpl{
		db: db,
	}
}

func (s *MBiodataAttachmentRepositoryImpl) Get(id uint) (*MBiodataAttachment, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	entity := MBiodataAttachment{}
	result := s.db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (s *MBiodataAttachmentRepositoryImpl) Create(entity *MBiodataAttachment) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MBiodataAttachmentRepositoryImpl) Update(entity *MBiodataAttachment) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Model(&entity).Updates(entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MBiodataAttachmentRepositoryImpl) Delete(id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var entity MBiodataAttachment
	result := s.db.Delete(&entity, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *MBiodataAttachmentRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "repository", "GetPageMBiodataAttachment", "")

	var entities []MBiodataAttachment
	var entity MBiodataAttachment
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
		Content:          ToMBiodataAttachmentResponsesParallel(entities),
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

	util.Log("INFO", "repository", "GetPageMBiodataAttachment", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
