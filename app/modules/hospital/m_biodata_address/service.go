package m_biodata_address

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type MBiodataAddressService interface {
	Get(id uint) (*MBiodataAddress, error)
	Create(payload *MBiodataAddressRequest, mUserId uint) error
	Update(payload *MBiodataAddressRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MBiodataAddressServiceImpl struct {
	repo MBiodataAddressRepository
}

func NewMBiodataAddressService(repo MBiodataAddressRepository) MBiodataAddressService {
	return &MBiodataAddressServiceImpl{
		repo: repo,
	}
}

func (s *MBiodataAddressServiceImpl) Get(id uint) (*MBiodataAddress, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *MBiodataAddressServiceImpl) Create(payload *MBiodataAddressRequest, mUserId uint) error {

	data := ToMBiodataAddressEntity(*payload, nil, mUserId)

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

func (s *MBiodataAddressServiceImpl) Update(payload *MBiodataAddressRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.BiodataId = payload.BiodataId
	existing.Label = payload.Label
	existing.Recipient = payload.Recipient
	existing.RecipientPhoneNumber = payload.RecipientPhoneNumber
	existing.LocationId = payload.LocationId
	existing.PostalCode = payload.PostalCode
	existing.Address = payload.Address
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

func (s *MBiodataAddressServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MBiodataAddressServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
