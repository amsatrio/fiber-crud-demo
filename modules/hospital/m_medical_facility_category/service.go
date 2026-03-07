package m_medical_facility_category

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    
    "time"
)

type MMedicalFacilityCategoryService interface {
    Get(id uint) (*MMedicalFacilityCategory, error)
    Create(payload *MMedicalFacilityCategoryRequest, mUserId uint) error
    Update(payload *MMedicalFacilityCategoryRequest, mUserId uint) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type MMedicalFacilityCategoryServiceImpl struct {
    repo     MMedicalFacilityCategoryRepository
}

func NewMMedicalFacilityCategoryService(repo MMedicalFacilityCategoryRepository) MMedicalFacilityCategoryService {
    return &MMedicalFacilityCategoryServiceImpl{
        repo:     repo,
    }
}

func (s *MMedicalFacilityCategoryServiceImpl) Get(id uint) (*MMedicalFacilityCategory, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *MMedicalFacilityCategoryServiceImpl) Create(payload *MMedicalFacilityCategoryRequest, mUserId uint) error {

    data := ToMMedicalFacilityCategoryEntity(*payload, nil, mUserId)

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

func (s *MMedicalFacilityCategoryServiceImpl) Update(payload *MMedicalFacilityCategoryRequest, mUserId uint) error {
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

func (s *MMedicalFacilityCategoryServiceImpl) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}

func (s *MMedicalFacilityCategoryServiceImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {
    return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
