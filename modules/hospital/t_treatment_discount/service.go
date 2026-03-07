package t_treatment_discount

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    
    "time"
)

type TTreatmentDiscountService interface {
    Get(id uint) (*TTreatmentDiscount, error)
    Create(payload *TTreatmentDiscountRequest, mUserId uint) error
    Update(payload *TTreatmentDiscountRequest, mUserId uint) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type TTreatmentDiscountServiceImpl struct {
    repo     TTreatmentDiscountRepository
}

func NewTTreatmentDiscountService(repo TTreatmentDiscountRepository) TTreatmentDiscountService {
    return &TTreatmentDiscountServiceImpl{
        repo:     repo,
    }
}

func (s *TTreatmentDiscountServiceImpl) Get(id uint) (*TTreatmentDiscount, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *TTreatmentDiscountServiceImpl) Create(payload *TTreatmentDiscountRequest, mUserId uint) error {

    data := ToTTreatmentDiscountEntity(*payload, nil, mUserId)

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

func (s *TTreatmentDiscountServiceImpl) Update(payload *TTreatmentDiscountRequest, mUserId uint) error {
    if payload.Id == nil {
        return errors.New("invalid payload")
    }

    existing, err := s.repo.Get(*payload.Id)
    if err != nil {
        return err
    }


    existing.DoctorOfficeTreatmentPriceId = payload.DoctorOfficeTreatmentPriceId
    existing.Value           = payload.Value
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

func (s *TTreatmentDiscountServiceImpl) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}

func (s *TTreatmentDiscountServiceImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {
    return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
