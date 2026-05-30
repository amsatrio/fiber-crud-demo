package t_doctor_office_treatment

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type TDoctorOfficeTreatmentService interface {
	Get(id uint) (*TDoctorOfficeTreatment, error)
	Create(payload *TDoctorOfficeTreatmentRequest, mUserId uint) error
	Update(payload *TDoctorOfficeTreatmentRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TDoctorOfficeTreatmentServiceImpl struct {
	repo TDoctorOfficeTreatmentRepository
}

func NewTDoctorOfficeTreatmentService(repo TDoctorOfficeTreatmentRepository) TDoctorOfficeTreatmentService {
	return &TDoctorOfficeTreatmentServiceImpl{
		repo: repo,
	}
}

func (s *TDoctorOfficeTreatmentServiceImpl) Get(id uint) (*TDoctorOfficeTreatment, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TDoctorOfficeTreatmentServiceImpl) Create(payload *TDoctorOfficeTreatmentRequest, mUserId uint) error {

	data := ToTDoctorOfficeTreatmentEntity(*payload, nil, mUserId)

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

func (s *TDoctorOfficeTreatmentServiceImpl) Update(payload *TDoctorOfficeTreatmentRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.DoctorTreatmentId = payload.DoctorTreatmentId
	existing.DoctorOfficeId = payload.DoctorOfficeId
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

func (s *TDoctorOfficeTreatmentServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TDoctorOfficeTreatmentServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
