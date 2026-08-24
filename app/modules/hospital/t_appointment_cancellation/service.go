package t_appointment_cancellation

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type TAppointmentCancellationService interface {
	Get(id uint) (*TAppointmentCancellation, error)
	Create(payload *TAppointmentCancellationRequest, mUserId uint) error
	Update(payload *TAppointmentCancellationRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TAppointmentCancellationServiceImpl struct {
	repo TAppointmentCancellationRepository
}

func NewTAppointmentCancellationService(repo TAppointmentCancellationRepository) TAppointmentCancellationService {
	return &TAppointmentCancellationServiceImpl{
		repo: repo,
	}
}

func (s *TAppointmentCancellationServiceImpl) Get(id uint) (*TAppointmentCancellation, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TAppointmentCancellationServiceImpl) Create(payload *TAppointmentCancellationRequest, mUserId uint) error {

	data := ToTAppointmentCancellationEntity(*payload, nil, mUserId)

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

func (s *TAppointmentCancellationServiceImpl) Update(payload *TAppointmentCancellationRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.AppointmentId = payload.AppointmentId
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

func (s *TAppointmentCancellationServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TAppointmentCancellationServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
