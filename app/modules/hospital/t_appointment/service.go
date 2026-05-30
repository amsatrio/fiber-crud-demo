package t_appointment

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type TAppointmentService interface {
	Get(id uint) (*TAppointment, error)
	Create(payload *TAppointmentRequest, mUserId uint) error
	Update(payload *TAppointmentRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TAppointmentServiceImpl struct {
	repo TAppointmentRepository
}

func NewTAppointmentService(repo TAppointmentRepository) TAppointmentService {
	return &TAppointmentServiceImpl{
		repo: repo,
	}
}

func (s *TAppointmentServiceImpl) Get(id uint) (*TAppointment, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TAppointmentServiceImpl) Create(payload *TAppointmentRequest, mUserId uint) error {

	data := ToTAppointmentEntity(*payload, nil, mUserId)

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

func (s *TAppointmentServiceImpl) Update(payload *TAppointmentRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.CustomerId = payload.CustomerId
	existing.DoctorOfficeId = payload.DoctorOfficeId
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

func (s *TAppointmentServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TAppointmentServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
