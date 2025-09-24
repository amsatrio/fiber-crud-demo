package m_role

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

type MRoleService interface {
	GetMRole(id uint) (*schema.MRole, error)
	CreateMRole(mRole *schema.MRoleRequest, mUserId uint) error
	UpdateMRole(mRole *schema.MRoleRequest, mUserId uint) error
	DeleteMRole(id uint, mUserId uint) error
	GetPageMRole(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MRoleServiceImpl struct {
	db *gorm.DB
}

func NewMRoleServiceImpl(db *gorm.DB) MRoleService {
	return &MRoleServiceImpl{
		db: db,
	}
}

func (s *MRoleServiceImpl) GetMRole(id uint) (*schema.MRole, error) {
	mRole := schema.MRole{}
	result := s.db.First(&mRole, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &mRole, nil
}

func (s *MRoleServiceImpl) CreateMRole(payload *schema.MRoleRequest, mUserId uint) error {

	mRole := payload.ToModelNew(mUserId)

	util.Log("INFO", "service", "MRoleService", "CreateMRole: ")

	var oldMRole schema.MRole

	// find data
	result := s.db.First(&oldMRole, mRole.Id)
	if result.Error == nil {
		return errors.New("data exist")
	}

	result = s.db.Create(&mRole)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MRoleServiceImpl) UpdateMRole(payload *schema.MRoleRequest, mUserId uint) error {

	var oldMRole *schema.MRole

	// find data
	result := s.db.First(&oldMRole, payload.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("data not found")
	}
	if result.Error != nil {
		return result.Error
	}

	// update data
	oldMRole.Name = payload.Name
	oldMRole.Code = payload.Code
	oldMRole.Level = payload.Level
	oldMRole.ModifiedBy = &mUserId
	oldMRole.ModifiedOn = &dto.JSONTime{Time: time.Now()}
	if *payload.IsDelete {
		oldMRole.DeletedBy = &mUserId
		oldMRole.DeletedOn = &dto.JSONTime{Time: time.Now()}
		oldMRole.IsDelete = payload.IsDelete
	}

	// update data for response
	result = s.db.Model(&oldMRole).Updates(oldMRole)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *MRoleServiceImpl) DeleteMRole(id uint, mUserId uint) error {
	var mRole schema.MRole
	result := s.db.Delete(&mRole, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *MRoleServiceImpl) GetPageMRole(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "service", "GetPageMRole", "")

	var mRoles []schema.MRole
	var mRole schema.MRole
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
