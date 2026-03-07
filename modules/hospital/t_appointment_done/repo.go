package t_appointment_done

import (
    "errors"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    "fiber-crud-demo/util"
    "strconv"
    "sync"

    "gorm.io/gorm"
)

type TAppointmentDoneRepository interface {
    Get(id uint) (*TAppointmentDone, error)
    Create(data *TAppointmentDone) error
    Update(data *TAppointmentDone) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type TAppointmentDoneRepositoryImpl struct {
    mutex sync.Mutex
    db    *gorm.DB
}

func NewTAppointmentDoneRepository(db *gorm.DB) TAppointmentDoneRepository {
    return &TAppointmentDoneRepositoryImpl{
        db: db,
    }
}

func (s *TAppointmentDoneRepositoryImpl) Get(id uint) (*TAppointmentDone, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    entity := TAppointmentDone{}
    result := s.db.First(&entity, id)
    if result.Error != nil {
        return nil, result.Error
    }

    return &entity, nil
}

func (s *TAppointmentDoneRepositoryImpl) Create(entity *TAppointmentDone) error {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    result := s.db.Create(&entity)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func (s *TAppointmentDoneRepositoryImpl) Update(entity *TAppointmentDone) error {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    result := s.db.Model(&entity).Updates(entity)

    if result.Error != nil {
        return result.Error
    }

    return nil
}

func (s *TAppointmentDoneRepositoryImpl) Delete(id uint) error {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    var entity TAppointmentDone
    result := s.db.Delete(&entity, id)

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("data not found")
    }
    return nil
}

func (s *TAppointmentDoneRepositoryImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {

    util.Log("INFO", "repository", "GetPageTAppointmentDone", "")

    var entities []TAppointmentDone
    var entity TAppointmentDone
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
        Content:          ToTAppointmentDoneResponsesParallel(entities),
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

    util.Log("INFO", "repository", "GetPageTAppointmentDone", "sort is empty: "+strconv.FormatBool(sort.Empty))

    return &page, nil
}
