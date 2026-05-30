package t_appointment_reschedule_history

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type TAppointmentRescheduleHistoryService interface {
	Get(id uint) (*TAppointmentRescheduleHistory, error)
	Create(payload *TAppointmentRescheduleHistoryRequest, mUserId uint) error
	Update(payload *TAppointmentRescheduleHistoryRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TAppointmentRescheduleHistoryServiceImpl struct {
	repo TAppointmentRescheduleHistoryRepository
}

func NewTAppointmentRescheduleHistoryService(repo TAppointmentRescheduleHistoryRepository) TAppointmentRescheduleHistoryService {
	return &TAppointmentRescheduleHistoryServiceImpl{
		repo: repo,
	}
}

func (s *TAppointmentRescheduleHistoryServiceImpl) Get(id uint) (*TAppointmentRescheduleHistory, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TAppointmentRescheduleHistoryServiceImpl) Create(payload *TAppointmentRescheduleHistoryRequest, mUserId uint) error {

	data := ToTAppointmentRescheduleHistoryEntity(*payload, nil, mUserId)

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

func (s *TAppointmentRescheduleHistoryServiceImpl) Update(payload *TAppointmentRescheduleHistoryRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.AppointmentId = payload.AppointmentId
	existing.DoctorOfficeScheduleId = payload.DoctorOfficeScheduleId
	existing.DoctorOfficeTreatmentId = payload.DoctorOfficeTreatmentId
	existing.AppointmentDate = payload.AppointmentDate
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

func (s *TAppointmentRescheduleHistoryServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TAppointmentRescheduleHistoryServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
