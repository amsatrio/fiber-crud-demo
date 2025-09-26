package repository

import (
	"errors"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/internal/domain"
	"fiber-crud-demo/util"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

type MBiodataRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewMBiodataRepository(db *gorm.DB) domain.MBiodataRepository {
	return &MBiodataRepositoryImpl{
		db: db,
	}
}

func (s *MBiodataRepositoryImpl) Get(id uint) (*domain.MBiodata, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	mBiodata := domain.MBiodata{}
	result := s.db.First(&mBiodata, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &mBiodata, nil
}

func (s *MBiodataRepositoryImpl) Create(mBiodata *domain.MBiodata) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Create(&mBiodata)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MBiodataRepositoryImpl) Update(mBiodata *domain.MBiodata) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Model(&mBiodata).Updates(mBiodata)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MBiodataRepositoryImpl) Delete(id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var mBiodata domain.MBiodata
	result := s.db.Delete(&mBiodata, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *MBiodataRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "service", "GetPageMBiodata", "")

	var mRoles []domain.MBiodata
	var mBiodata domain.MBiodata
	mRoleMap := util.GetJSONFieldTypes(mBiodata)

	// Create a DB instance and build the base query
	db := s.db

	// apply sorting
	db = util.ApplySorting(db, sortRequest)

	// apply filtering
	db = util.ApplyFiltering(db, filterRequest)

	// apply global search
	db = util.ApplyGlobalSearch(db, searchRequest, mRoleMap)

	util.Log("INFO", "service", "GetPageMBiodata", "")

	// Calculate the total data size without considering _size
	totalElements := db.Find(&mRoles).RowsAffected

	util.Log("INFO", "service", "GetPageMBiodata", "")

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

	util.Log("INFO", "service", "GetPageMBiodata", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
