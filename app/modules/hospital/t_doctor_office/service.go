package t_doctor_office

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type TDoctorOfficeService interface {
	Get(id uint) (*TDoctorOffice, error)
	Create(payload *TDoctorOfficeRequest, mUserId uint) error
	Update(payload *TDoctorOfficeRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TDoctorOfficeServiceImpl struct {
	repo TDoctorOfficeRepository
}

func NewTDoctorOfficeService(repo TDoctorOfficeRepository) TDoctorOfficeService {
	return &TDoctorOfficeServiceImpl{
		repo: repo,
	}
}

func (s *TDoctorOfficeServiceImpl) Get(id uint) (*TDoctorOffice, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TDoctorOfficeServiceImpl) Create(payload *TDoctorOfficeRequest, mUserId uint) error {

	data := ToTDoctorOfficeEntity(*payload, nil, mUserId)

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

func (s *TDoctorOfficeServiceImpl) Update(payload *TDoctorOfficeRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.DoctorId = payload.DoctorId
	existing.MedicalFacilityId = payload.MedicalFacilityId
	existing.Specialization = payload.Specialization
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

func (s *TDoctorOfficeServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TDoctorOfficeServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
