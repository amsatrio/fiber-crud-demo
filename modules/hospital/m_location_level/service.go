package m_location_level

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    
    "time"
)

type MLocationLevelService interface {
    Get(id uint) (*MLocationLevel, error)
    Create(payload *MLocationLevelRequest, mUserId uint) error
    Update(payload *MLocationLevelRequest, mUserId uint) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type MLocationLevelServiceImpl struct {
    repo     MLocationLevelRepository
}

func NewMLocationLevelService(repo MLocationLevelRepository) MLocationLevelService {
    return &MLocationLevelServiceImpl{
        repo:     repo,
    }
}

func (s *MLocationLevelServiceImpl) Get(id uint) (*MLocationLevel, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *MLocationLevelServiceImpl) Create(payload *MLocationLevelRequest, mUserId uint) error {

    data := ToMLocationLevelEntity(*payload, nil, mUserId)

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

func (s *MLocationLevelServiceImpl) Update(payload *MLocationLevelRequest, mUserId uint) error {
    if payload.Id == nil {
        return errors.New("invalid payload")
    }

    existing, err := s.repo.Get(*payload.Id)
    if err != nil {
        return err
    }


    existing.Name            = payload.Name
    existing.Abbreviation    = payload.Abbreviation
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

func (s *MLocationLevelServiceImpl) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}

func (s *MLocationLevelServiceImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {
    return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
