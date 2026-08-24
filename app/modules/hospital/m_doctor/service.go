package m_doctor

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type MDoctorService interface {
	Get(id uint) (*MDoctor, error)
	Create(payload *MDoctorRequest, mUserId uint) error
	Update(payload *MDoctorRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MDoctorServiceImpl struct {
	repo MDoctorRepository
}

func NewMDoctorService(repo MDoctorRepository) MDoctorService {
	return &MDoctorServiceImpl{
		repo: repo,
	}
}

func (s *MDoctorServiceImpl) Get(id uint) (*MDoctor, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *MDoctorServiceImpl) Create(payload *MDoctorRequest, mUserId uint) error {

	data := ToMDoctorEntity(*payload, nil, mUserId)

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

func (s *MDoctorServiceImpl) Update(payload *MDoctorRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.BiodataId = payload.BiodataId
	existing.StrCode = payload.StrCode
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

func (s *MDoctorServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MDoctorServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
