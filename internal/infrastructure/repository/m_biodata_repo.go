package repository

import (
	"errors"
	"fiber-crud-demo/dto"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/internal/domain"
	"fiber-crud-demo/util"
	"strconv"
	"sync"
	"time"

	"gorm.io/gorm"
)

type MBiodataRepository struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewMBiodataRepository(db *gorm.DB) domain.MBiodataRepository {
	return &MBiodataRepository{
		db: db,
	}
}

func (s *MBiodataRepository) GetMBiodata(id uint) (*domain.MBiodata, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	mBiodata := domain.MBiodata{}
	result := s.db.First(&mBiodata, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &mBiodata, nil
}

func (s *MBiodataRepository) CreateMBiodata(payload *domain.MBiodataRequest, mUserId uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	mBiodata := payload.ToModelNew(mUserId)

	util.Log("INFO", "service", "MBiodataService", "CreateMBiodata: ")

	var oldMBiodata domain.MBiodata

	// find data
	result := s.db.First(&oldMBiodata, mBiodata.Id)
	if result.Error == nil {
		return errors.New("data exist")
	}

	result = s.db.Create(&mBiodata)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MBiodataRepository) UpdateMBiodata(payload *domain.MBiodataRequest, mUserId uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var oldMBiodata *domain.MBiodata

	// find data
	result := s.db.First(&oldMBiodata, payload.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("data not found")
	}
	if result.Error != nil {
		return result.Error
	}

	// update data
	oldMBiodata.Fullname = payload.Fullname
	oldMBiodata.MobilePhone = payload.MobilePhone
	oldMBiodata.Image = payload.Image
	oldMBiodata.ImagePath = &payload.ImagePath
	oldMBiodata.ModifiedBy = &mUserId
	oldMBiodata.ModifiedOn = &dto.JSONTime{Time: time.Now()}
	if *payload.IsDelete {
		oldMBiodata.DeletedBy = &mUserId
		oldMBiodata.DeletedOn = &dto.JSONTime{Time: time.Now()}
		oldMBiodata.IsDelete = payload.IsDelete
	}

	// update data for response
	result = s.db.Model(&oldMBiodata).Updates(oldMBiodata)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MBiodataRepository) DeleteMBiodata(id uint, mUserId uint) error {
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

func (s *MBiodataRepository) GetPageMBiodata(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

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
