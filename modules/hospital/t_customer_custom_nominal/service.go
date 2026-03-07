package t_customer_custom_nominal

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    
    "time"
)

type TCustomerCustomNominalService interface {
    Get(id uint) (*TCustomerCustomNominal, error)
    Create(payload *TCustomerCustomNominalRequest, mUserId uint) error
    Update(payload *TCustomerCustomNominalRequest, mUserId uint) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type TCustomerCustomNominalServiceImpl struct {
    repo     TCustomerCustomNominalRepository
}

func NewTCustomerCustomNominalService(repo TCustomerCustomNominalRepository) TCustomerCustomNominalService {
    return &TCustomerCustomNominalServiceImpl{
        repo:     repo,
    }
}

func (s *TCustomerCustomNominalServiceImpl) Get(id uint) (*TCustomerCustomNominal, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *TCustomerCustomNominalServiceImpl) Create(payload *TCustomerCustomNominalRequest, mUserId uint) error {

    data := ToTCustomerCustomNominalEntity(*payload, nil, mUserId)

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

func (s *TCustomerCustomNominalServiceImpl) Update(payload *TCustomerCustomNominalRequest, mUserId uint) error {
    if payload.Id == nil {
        return errors.New("invalid payload")
    }

    existing, err := s.repo.Get(*payload.Id)
    if err != nil {
        return err
    }


    existing.CustomerId      = payload.CustomerId
    existing.Nominal         = payload.Nominal
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

func (s *TCustomerCustomNominalServiceImpl) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}

func (s *TCustomerCustomNominalServiceImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {
    return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
