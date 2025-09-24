package m_user

import (
	"errors"
	"fiber-crud-demo/dto"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/dto/schema"
	"fiber-crud-demo/util"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type MUserService interface {
	GetMUser(id uint) (*schema.MUser, error)
	CreateMUser(mUser *schema.MUserRequest, mUserId uint) error
	UpdateMUser(mUser *schema.MUserRequest, mUserId uint) error
	DeleteMUser(id uint, mUserId uint) error
	GetPageMUser(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MUserServiceImpl struct {
	db *gorm.DB
}

func NewMUserServiceImpl(db *gorm.DB) MUserService {
	return &MUserServiceImpl{
		db: db,
	}
}

func (s *MUserServiceImpl) GetMUser(id uint) (*schema.MUser, error) {
	mUser := schema.MUser{}
	result := s.db.First(&mUser, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &mUser, nil
}

func (s *MUserServiceImpl) CreateMUser(payload *schema.MUserRequest, mUserId uint) error {

	mUser := payload.ToModelNew(mUserId)

	util.Log("INFO", "service", "MUserService", "CreateMUser: ")

	var oldMUser schema.MUser

	// find data
	result := s.db.First(&oldMUser, mUser.Id)
	if result.Error == nil {
		return errors.New("data exist")
	}

	result = s.db.Create(&mUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MUserServiceImpl) UpdateMUser(payload *schema.MUserRequest, mUserId uint) error {

	var oldMUser *schema.MUser

	// find data
	result := s.db.First(&oldMUser, payload.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("data not found")
	}
	if result.Error != nil {
		return result.Error
	}

	// update data
	oldMUser.BiodataId = payload.BiodataId
	oldMUser.RoleId = payload.RoleId
	oldMUser.Email = payload.Email
	oldMUser.Password = payload.Password
	oldMUser.LoginAttempt = payload.LoginAttempt
	oldMUser.IsLocked = payload.IsLocked
	oldMUser.LastLogin = payload.LastLogin
	oldMUser.ModifiedBy = &mUserId
	oldMUser.ModifiedOn = &dto.JSONTime{Time: time.Now()}
	if *payload.IsDelete {
		oldMUser.DeletedBy = &mUserId
		oldMUser.DeletedOn = &dto.JSONTime{Time: time.Now()}
		oldMUser.IsDelete = payload.IsDelete
	}

	// update data for response
	result = s.db.Model(&oldMUser).Updates(oldMUser)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MUserServiceImpl) DeleteMUser(id uint, mUserId uint) error {
	var mUser schema.MUser
	result := s.db.Delete(&mUser, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *MUserServiceImpl) GetPageMUser(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "service", "GetPageMUser", "")

	var mRoles []schema.MUser
	var mUser schema.MUser
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
