package t_doctor_office_schedule

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type TDoctorOfficeScheduleService interface {
	Get(id uint) (*TDoctorOfficeSchedule, error)
	Create(payload *TDoctorOfficeScheduleRequest, mUserId uint) error
	Update(payload *TDoctorOfficeScheduleRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TDoctorOfficeScheduleServiceImpl struct {
	repo TDoctorOfficeScheduleRepository
}

func NewTDoctorOfficeScheduleService(repo TDoctorOfficeScheduleRepository) TDoctorOfficeScheduleService {
	return &TDoctorOfficeScheduleServiceImpl{
		repo: repo,
	}
}

func (s *TDoctorOfficeScheduleServiceImpl) Get(id uint) (*TDoctorOfficeSchedule, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TDoctorOfficeScheduleServiceImpl) Create(payload *TDoctorOfficeScheduleRequest, mUserId uint) error {

	data := ToTDoctorOfficeScheduleEntity(*payload, nil, mUserId)

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

func (s *TDoctorOfficeScheduleServiceImpl) Update(payload *TDoctorOfficeScheduleRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.DoctorId = payload.DoctorId
	existing.MedicalFacilityScheduleId = payload.MedicalFacilityScheduleId
	existing.Slot = payload.Slot
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

func (s *TDoctorOfficeScheduleServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TDoctorOfficeScheduleServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
