package m_wallet_default_nominal

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type MWalletDefaultNominalService interface {
	Get(id uint) (*MWalletDefaultNominal, error)
	Create(payload *MWalletDefaultNominalRequest, mUserId uint) error
	Update(payload *MWalletDefaultNominalRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MWalletDefaultNominalServiceImpl struct {
	repo MWalletDefaultNominalRepository
}

func NewMWalletDefaultNominalService(repo MWalletDefaultNominalRepository) MWalletDefaultNominalService {
	return &MWalletDefaultNominalServiceImpl{
		repo: repo,
	}
}

func (s *MWalletDefaultNominalServiceImpl) Get(id uint) (*MWalletDefaultNominal, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *MWalletDefaultNominalServiceImpl) Create(payload *MWalletDefaultNominalRequest, mUserId uint) error {

	data := ToMWalletDefaultNominalEntity(*payload, nil, mUserId)

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

func (s *MWalletDefaultNominalServiceImpl) Update(payload *MWalletDefaultNominalRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.Nominal = payload.Nominal
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

func (s *MWalletDefaultNominalServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MWalletDefaultNominalServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
