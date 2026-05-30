package t_current_doctor_specialization

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type TCurrentDoctorSpecializationService interface {
	Get(id uint) (*TCurrentDoctorSpecialization, error)
	Create(payload *TCurrentDoctorSpecializationRequest, mUserId uint) error
	Update(payload *TCurrentDoctorSpecializationRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TCurrentDoctorSpecializationServiceImpl struct {
	repo TCurrentDoctorSpecializationRepository
}

func NewTCurrentDoctorSpecializationService(repo TCurrentDoctorSpecializationRepository) TCurrentDoctorSpecializationService {
	return &TCurrentDoctorSpecializationServiceImpl{
		repo: repo,
	}
}

func (s *TCurrentDoctorSpecializationServiceImpl) Get(id uint) (*TCurrentDoctorSpecialization, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TCurrentDoctorSpecializationServiceImpl) Create(payload *TCurrentDoctorSpecializationRequest, mUserId uint) error {

	data := ToTCurrentDoctorSpecializationEntity(*payload, nil, mUserId)

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

func (s *TCurrentDoctorSpecializationServiceImpl) Update(payload *TCurrentDoctorSpecializationRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.DoctorId = payload.DoctorId
	existing.SpecializationId = payload.SpecializationId
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

func (s *TCurrentDoctorSpecializationServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TCurrentDoctorSpecializationServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
