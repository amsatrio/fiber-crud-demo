package m_customer

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    
    "time"
)

type MCustomerService interface {
    Get(id uint) (*MCustomer, error)
    Create(payload *MCustomerRequest, mUserId uint) error
    Update(payload *MCustomerRequest, mUserId uint) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type MCustomerServiceImpl struct {
    repo     MCustomerRepository
}

func NewMCustomerService(repo MCustomerRepository) MCustomerService {
    return &MCustomerServiceImpl{
        repo:     repo,
    }
}

func (s *MCustomerServiceImpl) Get(id uint) (*MCustomer, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *MCustomerServiceImpl) Create(payload *MCustomerRequest, mUserId uint) error {

    data := ToMCustomerEntity(*payload, nil, mUserId)

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

func (s *MCustomerServiceImpl) Update(payload *MCustomerRequest, mUserId uint) error {
    if payload.Id == nil {
        return errors.New("invalid payload")
    }

    existing, err := s.repo.Get(*payload.Id)
    if err != nil {
        return err
    }


    existing.BiodataId       = payload.BiodataId
    existing.Dob             = payload.Dob
    existing.Gender          = payload.Gender
    existing.BloodGroupId    = payload.BloodGroupId
    existing.RhesusType      = payload.RhesusType
    existing.Height          = payload.Height
    existing.Weight          = payload.Weight
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

func (s *MCustomerServiceImpl) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}

func (s *MCustomerServiceImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {
    return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
