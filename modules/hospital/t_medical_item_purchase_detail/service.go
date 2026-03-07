package t_medical_item_purchase_detail

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    
    "time"
)

type TMedicalItemPurchaseDetailService interface {
    Get(id uint) (*TMedicalItemPurchaseDetail, error)
    Create(payload *TMedicalItemPurchaseDetailRequest, mUserId uint) error
    Update(payload *TMedicalItemPurchaseDetailRequest, mUserId uint) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type TMedicalItemPurchaseDetailServiceImpl struct {
    repo     TMedicalItemPurchaseDetailRepository
}

func NewTMedicalItemPurchaseDetailService(repo TMedicalItemPurchaseDetailRepository) TMedicalItemPurchaseDetailService {
    return &TMedicalItemPurchaseDetailServiceImpl{
        repo:     repo,
    }
}

func (s *TMedicalItemPurchaseDetailServiceImpl) Get(id uint) (*TMedicalItemPurchaseDetail, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *TMedicalItemPurchaseDetailServiceImpl) Create(payload *TMedicalItemPurchaseDetailRequest, mUserId uint) error {

    data := ToTMedicalItemPurchaseDetailEntity(*payload, nil, mUserId)

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

func (s *TMedicalItemPurchaseDetailServiceImpl) Update(payload *TMedicalItemPurchaseDetailRequest, mUserId uint) error {
    if payload.Id == nil {
        return errors.New("invalid payload")
    }

    existing, err := s.repo.Get(*payload.Id)
    if err != nil {
        return err
    }


    existing.MedicalItemPurchaseId = payload.MedicalItemPurchaseId
    existing.MedicalItemId   = payload.MedicalItemId
    existing.Qty             = payload.Qty
    existing.MedicalFacilityId = payload.MedicalFacilityId
    existing.CourirId        = payload.CourirId
    existing.SubTotal        = payload.SubTotal
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

func (s *TMedicalItemPurchaseDetailServiceImpl) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}

func (s *TMedicalItemPurchaseDetailServiceImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {
    return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
