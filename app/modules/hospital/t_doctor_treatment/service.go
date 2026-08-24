package t_doctor_treatment

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type TDoctorTreatmentService interface {
	Get(id uint) (*TDoctorTreatment, error)
	Create(payload *TDoctorTreatmentRequest, mUserId uint) error
	Update(payload *TDoctorTreatmentRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TDoctorTreatmentServiceImpl struct {
	repo TDoctorTreatmentRepository
}

func NewTDoctorTreatmentService(repo TDoctorTreatmentRepository) TDoctorTreatmentService {
	return &TDoctorTreatmentServiceImpl{
		repo: repo,
	}
}

func (s *TDoctorTreatmentServiceImpl) Get(id uint) (*TDoctorTreatment, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TDoctorTreatmentServiceImpl) Create(payload *TDoctorTreatmentRequest, mUserId uint) error {

	data := ToTDoctorTreatmentEntity(*payload, nil, mUserId)

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

func (s *TDoctorTreatmentServiceImpl) Update(payload *TDoctorTreatmentRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.DoctorId = payload.DoctorId
	existing.Name = payload.Name
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

func (s *TDoctorTreatmentServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TDoctorTreatmentServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
