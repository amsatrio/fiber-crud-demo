package m_role

import (
	"errors"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/util"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

type MRoleRepository interface {
	Get(id uint) (*MRole, error)
	Create(data *MRole) error
	Update(data *MRole) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MRoleRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewMRoleRepository(db *gorm.DB) MRoleRepository {
	return &MRoleRepositoryImpl{
		db: db,
	}
}

func (s *MRoleRepositoryImpl) Get(id uint) (*MRole, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	mRole := MRole{}
	result := s.db.First(&mRole, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &mRole, nil
}

func (s *MRoleRepositoryImpl) Create(mRole *MRole) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Create(&mRole)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MRoleRepositoryImpl) Update(mRole *MRole) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Model(&mRole).Updates(mRole)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MRoleRepositoryImpl) Delete(id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var mRole MRole
	result := s.db.Delete(&mRole, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *MRoleRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "service", "GetPageMRole", "")

	var mRoles []MRole
	var mRole MRole
	mRoleMap := util.GetJSONFieldTypes(mRole)

	// Create a DB instance and build the base query
	db := s.db

	// apply sorting
	db = util.ApplySorting(db, sortRequest)

	// apply filtering
	db = util.ApplyFiltering(db, filterRequest)

	// apply global search
	db = util.ApplyGlobalSearch(db, searchRequest, mRoleMap)

	util.Log("INFO", "service", "GetPageMRole", "")

	// Calculate the total data size without considering _size
	totalElements := db.Find(&mRoles).RowsAffected

	util.Log("INFO", "service", "GetPageMRole", "")

	// Calculate the total number of pages
	totalPages := totalElements / sizeInt64
	if totalElements%sizeInt64 != 0 {
		totalPages++
	}

	// paginate
	result := db.Scopes(util.ApplyPaginate(pageInt, sizeInt)).Find(&mRoles)

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
		Content:          mRoles,
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

	util.Log("INFO", "service", "GetPageMRole", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
