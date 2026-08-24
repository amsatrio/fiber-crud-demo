package t_customer_wallet

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type TCustomerWalletService interface {
	Get(id uint) (*TCustomerWallet, error)
	Create(payload *TCustomerWalletRequest, mUserId uint) error
	Update(payload *TCustomerWalletRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TCustomerWalletServiceImpl struct {
	repo TCustomerWalletRepository
}

func NewTCustomerWalletService(repo TCustomerWalletRepository) TCustomerWalletService {
	return &TCustomerWalletServiceImpl{
		repo: repo,
	}
}

func (s *TCustomerWalletServiceImpl) Get(id uint) (*TCustomerWallet, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TCustomerWalletServiceImpl) Create(payload *TCustomerWalletRequest, mUserId uint) error {

	data := ToTCustomerWalletEntity(*payload, nil, mUserId)

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

func (s *TCustomerWalletServiceImpl) Update(payload *TCustomerWalletRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.CustomerId = payload.CustomerId
	existing.Pin = payload.Pin
	existing.Balance = payload.Balance
	existing.Barcode = payload.Barcode
	existing.Points = payload.Points
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

func (s *TCustomerWalletServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TCustomerWalletServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
