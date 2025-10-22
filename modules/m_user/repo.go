package m_user

import (
	"errors"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/util"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

type MUserRepository interface {
	Get(id uint) (*MUser, error)
	Create(data *MUser) error
	Update(data *MUser) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MUserRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewMUserRepository(db *gorm.DB) MUserRepository {
	return &MUserRepositoryImpl{
		db: db,
	}
}

func (s *MUserRepositoryImpl) Get(id uint) (*MUser, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	mUser := MUser{}
	result := s.db.First(&mUser, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &mUser, nil
}

func (s *MUserRepositoryImpl) Create(mUser *MUser) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Create(&mUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MUserRepositoryImpl) Update(mUser *MUser) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Model(&mUser).Updates(mUser)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MUserRepositoryImpl) Delete(id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var mUser MUser
	result := s.db.Delete(&mUser, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *MUserRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "service", "GetPageMUser", "")

	var mRoles []MUser
	var mUser MUser
	mRoleMap := util.GetJSONFieldTypes(mUser)

	// Create a DB instance and build the base query
	db := s.db

	// apply sorting
	db = util.ApplySorting(db, sortRequest)

	// apply filtering
	db = util.ApplyFiltering(db, filterRequest)

	// apply global search
	db = util.ApplyGlobalSearch(db, searchRequest, mRoleMap)

	util.Log("INFO", "service", "GetPageMUser", "")

	// Calculate the total data size without considering _size
	totalElements := db.Find(&mRoles).RowsAffected

	util.Log("INFO", "service", "GetPageMUser", "")

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

	util.Log("INFO", "service", "GetPageMUser", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
