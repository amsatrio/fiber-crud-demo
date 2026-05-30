package t_customer_wallet_withdraw

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type TCustomerWalletWithdrawService interface {
	Get(id uint) (*TCustomerWalletWithdraw, error)
	Create(payload *TCustomerWalletWithdrawRequest, mUserId uint) error
	Update(payload *TCustomerWalletWithdrawRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TCustomerWalletWithdrawServiceImpl struct {
	repo TCustomerWalletWithdrawRepository
}

func NewTCustomerWalletWithdrawService(repo TCustomerWalletWithdrawRepository) TCustomerWalletWithdrawService {
	return &TCustomerWalletWithdrawServiceImpl{
		repo: repo,
	}
}

func (s *TCustomerWalletWithdrawServiceImpl) Get(id uint) (*TCustomerWalletWithdraw, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TCustomerWalletWithdrawServiceImpl) Create(payload *TCustomerWalletWithdrawRequest, mUserId uint) error {

	data := ToTCustomerWalletWithdrawEntity(*payload, nil, mUserId)

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

func (s *TCustomerWalletWithdrawServiceImpl) Update(payload *TCustomerWalletWithdrawRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.CustomerId = payload.CustomerId
	existing.WalletDefaultNominalId = payload.WalletDefaultNominalId
	existing.Amount = payload.Amount
	existing.BankName = payload.BankName
	existing.AccountNumber = payload.AccountNumber
	existing.AccountName = payload.AccountName
	existing.Otp = payload.Otp
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

func (s *TCustomerWalletWithdrawServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TCustomerWalletWithdrawServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
