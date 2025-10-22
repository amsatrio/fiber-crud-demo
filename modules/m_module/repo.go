package m_module

import (
	"errors"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/util"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

type MModuleRepository interface {
	Get(id uint) (*MModule, error)
	Create(data *MModule) error
	Update(data *MModule) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}
type MModuleRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewMModuleRepository(db *gorm.DB) MModuleRepository {
	return &MModuleRepositoryImpl{
		db: db,
	}
}

func (s *MModuleRepositoryImpl) Get(id uint) (*MModule, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	mModule := MModule{}
	result := s.db.First(&mModule, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &mModule, nil
}

func (s *MModuleRepositoryImpl) Create(mModule *MModule) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Create(&mModule)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MModuleRepositoryImpl) Update(mModule *MModule) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Model(&mModule).Updates(mModule)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MModuleRepositoryImpl) Delete(id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var mModule MModule
	result := s.db.Delete(&mModule, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *MModuleRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "service", "GetPageMModule", "")

	var mModules []MModule
	var mModule MModule
	mModuleMap := util.GetJSONFieldTypes(mModule)

	// Create a DB instance and build the base query
	db := s.db

	// apply sorting
	db = util.ApplySorting(db, sortRequest)

	// apply filtering
	db = util.ApplyFiltering(db, filterRequest)

	// apply global search
	db = util.ApplyGlobalSearch(db, searchRequest, mModuleMap)

	util.Log("INFO", "service", "GetPageMModule", "")

	// Calculate the total data size without considering _size
	totalElements := db.Find(&mModules).RowsAffected

	util.Log("INFO", "service", "GetPageMModule", "")

	// Calculate the total number of pages
	totalPages := totalElements / sizeInt64
	if totalElements%sizeInt64 != 0 {
		totalPages++
	}

	// paginate
	result := db.Scopes(util.ApplyPaginate(pageInt, sizeInt)).Find(&mModules)

	if result.Error != nil {
		return nil, result.Error
	}

	lastPage := int64(pageInt) == totalPages-1
	firstPage := pageInt == 0

	// prepare page
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
		Content:          mModules,
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

	util.Log("INFO", "service", "GetPageMModule", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
