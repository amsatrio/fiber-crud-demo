package m_blood_group

import (
    "errors"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    "fiber-crud-demo/util"
    "strconv"
    "sync"

    "gorm.io/gorm"
)

type MBloodGroupRepository interface {
    Get(id uint) (*MBloodGroup, error)
    Create(data *MBloodGroup) error
    Update(data *MBloodGroup) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type MBloodGroupRepositoryImpl struct {
    mutex sync.Mutex
    db    *gorm.DB
}

func NewMBloodGroupRepository(db *gorm.DB) MBloodGroupRepository {
    return &MBloodGroupRepositoryImpl{
        db: db,
    }
}

func (s *MBloodGroupRepositoryImpl) Get(id uint) (*MBloodGroup, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    entity := MBloodGroup{}
    result := s.db.First(&entity, id)
    if result.Error != nil {
        return nil, result.Error
    }

    return &entity, nil
}

func (s *MBloodGroupRepositoryImpl) Create(entity *MBloodGroup) error {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    result := s.db.Create(&entity)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func (s *MBloodGroupRepositoryImpl) Update(entity *MBloodGroup) error {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    result := s.db.Model(&entity).Updates(entity)

    if result.Error != nil {
        return result.Error
    }

    return nil
}

func (s *MBloodGroupRepositoryImpl) Delete(id uint) error {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    var entity MBloodGroup
    result := s.db.Delete(&entity, id)

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("data not found")
    }
    return nil
}

func (s *MBloodGroupRepositoryImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {

    util.Log("INFO", "repository", "GetPageMBloodGroup", "")

    var entities []MBloodGroup
    var entity MBloodGroup
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
        Content:          ToMBloodGroupResponsesParallel(entities),
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

    util.Log("INFO", "repository", "GetPageMBloodGroup", "sort is empty: "+strconv.FormatBool(sort.Empty))

    return &page, nil
}
