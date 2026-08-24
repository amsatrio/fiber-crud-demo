package m_medical_facility

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type MMedicalFacilityService interface {
	Get(id uint) (*MMedicalFacility, error)
	Create(payload *MMedicalFacilityRequest, mUserId uint) error
	Update(payload *MMedicalFacilityRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MMedicalFacilityServiceImpl struct {
	repo MMedicalFacilityRepository
}

func NewMMedicalFacilityService(repo MMedicalFacilityRepository) MMedicalFacilityService {
	return &MMedicalFacilityServiceImpl{
		repo: repo,
	}
}

func (s *MMedicalFacilityServiceImpl) Get(id uint) (*MMedicalFacility, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *MMedicalFacilityServiceImpl) Create(payload *MMedicalFacilityRequest, mUserId uint) error {

	data := ToMMedicalFacilityEntity(*payload, nil, mUserId)

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

func (s *MMedicalFacilityServiceImpl) Update(payload *MMedicalFacilityRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.Name = payload.Name
	existing.MedicalFacilityCategoryId = payload.MedicalFacilityCategoryId
	existing.LocationId = payload.LocationId
	existing.FullAddress = payload.FullAddress
	existing.Email = payload.Email
	existing.PhoneCode = payload.PhoneCode
	existing.Phone = payload.Phone
	existing.Fax = payload.Fax
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

func (s *MMedicalFacilityServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MMedicalFacilityServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
