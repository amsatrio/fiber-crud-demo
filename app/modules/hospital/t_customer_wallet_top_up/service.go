package t_customer_wallet_top_up

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type TCustomerWalletTopUpService interface {
	Get(id uint) (*TCustomerWalletTopUp, error)
	Create(payload *TCustomerWalletTopUpRequest, mUserId uint) error
	Update(payload *TCustomerWalletTopUpRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TCustomerWalletTopUpServiceImpl struct {
	repo TCustomerWalletTopUpRepository
}

func NewTCustomerWalletTopUpService(repo TCustomerWalletTopUpRepository) TCustomerWalletTopUpService {
	return &TCustomerWalletTopUpServiceImpl{
		repo: repo,
	}
}

func (s *TCustomerWalletTopUpServiceImpl) Get(id uint) (*TCustomerWalletTopUp, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TCustomerWalletTopUpServiceImpl) Create(payload *TCustomerWalletTopUpRequest, mUserId uint) error {

	data := ToTCustomerWalletTopUpEntity(*payload, nil, mUserId)

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

func (s *TCustomerWalletTopUpServiceImpl) Update(payload *TCustomerWalletTopUpRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.CustomerWalletId = payload.CustomerWalletId
	existing.Amount = payload.Amount
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

func (s *TCustomerWalletTopUpServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TCustomerWalletTopUpServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
