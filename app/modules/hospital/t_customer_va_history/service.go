package t_customer_va_history

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type TCustomerVaHistoryService interface {
	Get(id uint) (*TCustomerVaHistory, error)
	Create(payload *TCustomerVaHistoryRequest, mUserId uint) error
	Update(payload *TCustomerVaHistoryRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TCustomerVaHistoryServiceImpl struct {
	repo TCustomerVaHistoryRepository
}

func NewTCustomerVaHistoryService(repo TCustomerVaHistoryRepository) TCustomerVaHistoryService {
	return &TCustomerVaHistoryServiceImpl{
		repo: repo,
	}
}

func (s *TCustomerVaHistoryServiceImpl) Get(id uint) (*TCustomerVaHistory, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TCustomerVaHistoryServiceImpl) Create(payload *TCustomerVaHistoryRequest, mUserId uint) error {

	data := ToTCustomerVaHistoryEntity(*payload, nil, mUserId)

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

func (s *TCustomerVaHistoryServiceImpl) Update(payload *TCustomerVaHistoryRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.CustomerVaId = payload.CustomerVaId
	existing.Amount = payload.Amount
	existing.ExpiredOn = payload.ExpiredOn
	existing.ModifiedBy = &mUserId
	existing.ModifiedOn = &dto.JSONTime{Time: time.Now()}
	existing.DeletedBy = nil
	existing.DeletedOn = nil

	if payload.IsDelete != nil && *payload.IsDelete {
		existing.DeletedBy = &mUserId
		existing.DeletedOn = &dto.JSONTime{Time: time.Now()}
		existing.IsDelete = *payload.IsDelete
	}

	return s.repo.Update(existing)
}

func (s *TCustomerVaHistoryServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TCustomerVaHistoryServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
