package m_courier

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    
    "time"
)

type MCourierService interface {
    Get(id uint) (*MCourier, error)
    Create(payload *MCourierRequest, mUserId uint) error
    Update(payload *MCourierRequest, mUserId uint) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type MCourierServiceImpl struct {
    repo     MCourierRepository
}

func NewMCourierService(repo MCourierRepository) MCourierService {
    return &MCourierServiceImpl{
        repo:     repo,
    }
}

func (s *MCourierServiceImpl) Get(id uint) (*MCourier, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *MCourierServiceImpl) Create(payload *MCourierRequest, mUserId uint) error {

    data := ToMCourierEntity(*payload, nil, mUserId)

    if payload.Id == nil {
        return s.repo.Create(&data)
    }

    _, err := s.repo.Get(*payload.Id)
    if err == nil {
        return errors.New("data exists")
    }

    data.Id = *payload.Id
    return s.repo.Create(&data)
}

func (s *MCourierServiceImpl) Update(payload *MCourierRequest, mUserId uint) error {
    if payload.Id == nil {
        return errors.New("invalid payload")
    }

    existing, err := s.repo.Get(*payload.Id)
    if err != nil {
        return err
    }


    existing.Name            = payload.Name
    existing.ModifiedBy = &mUserId
    existing.ModifiedOn = &dto.JSONTime{Time: time.Now()}
    existing.DeletedBy  = nil
    existing.DeletedOn  = nil

    if payload.IsDelete != nil && *payload.IsDelete {
        existing.DeletedBy = &mUserId
        existing.DeletedOn = &dto.JSONTime{Time: time.Now()}
        existing.IsDelete = *payload.IsDelete
    }

    return s.repo.Update(existing)
}

func (s *MCourierServiceImpl) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}

func (s *MCourierServiceImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {
    return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
