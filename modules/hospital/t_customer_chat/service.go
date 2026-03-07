package t_customer_chat

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    
    "time"
)

type TCustomerChatService interface {
    Get(id uint) (*TCustomerChat, error)
    Create(payload *TCustomerChatRequest, mUserId uint) error
    Update(payload *TCustomerChatRequest, mUserId uint) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type TCustomerChatServiceImpl struct {
    repo     TCustomerChatRepository
}

func NewTCustomerChatService(repo TCustomerChatRepository) TCustomerChatService {
    return &TCustomerChatServiceImpl{
        repo:     repo,
    }
}

func (s *TCustomerChatServiceImpl) Get(id uint) (*TCustomerChat, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *TCustomerChatServiceImpl) Create(payload *TCustomerChatRequest, mUserId uint) error {

    data := ToTCustomerChatEntity(*payload, nil, mUserId)

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

func (s *TCustomerChatServiceImpl) Update(payload *TCustomerChatRequest, mUserId uint) error {
    if payload.Id == nil {
        return errors.New("invalid payload")
    }

    existing, err := s.repo.Get(*payload.Id)
    if err != nil {
        return err
    }


    existing.CustomerId      = payload.CustomerId
    existing.DoctorId        = payload.DoctorId
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

func (s *TCustomerChatServiceImpl) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}

func (s *TCustomerChatServiceImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {
    return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
