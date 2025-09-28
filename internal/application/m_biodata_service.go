package application

import (
	"errors"
	"fiber-crud-demo/dto"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/internal/domain"
	"time"
)

type MBiodataService interface {
	Get(id uint) (*domain.MBiodata, error)
	Create(payload *domain.MBiodataRequest, mUserId uint) error
	Update(payload *domain.MBiodataRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MBiodataServiceImpl struct {
	repo domain.MBiodataRepository
}

func NewMBiodataService(repo domain.MBiodataRepository) MBiodataService {
	return &MBiodataServiceImpl{
		repo: repo,
	}
}

func (s *MBiodataServiceImpl) Get(id uint) (*domain.MBiodata, error) {
	return s.repo.Get(id)
}

func (s *MBiodataServiceImpl) Create(payload *domain.MBiodataRequest, mUserId uint) error {
	bool_true := false
	data := &domain.MBiodata{
		Id:          payload.Id,
		Fullname:    payload.Fullname,
		MobilePhone: payload.MobilePhone,
		Image:       payload.Image,
		ImagePath:   &payload.ImagePath,
		CreatedOn:   dto.JSONTime{Time: time.Now()},
		CreatedBy:   mUserId,
		IsDelete:    &bool_true,
	}

	if payload.Id == nil {
		return s.repo.Create(data)
	}

	_, err := s.repo.Get(*payload.Id)
	if err == nil {
		return errors.New("data exists")
	}

	data.Id = nil

	return s.repo.Create(data)
}

func (s *MBiodataServiceImpl) Update(payload *domain.MBiodataRequest, mUserId uint) error {

	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.Fullname = payload.Fullname
	existing.MobilePhone = payload.MobilePhone
	existing.Image = payload.Image
	existing.ImagePath = &payload.ImagePath
	existing.ModifiedBy = &mUserId
	existing.ModifiedOn = &dto.JSONTime{Time: time.Now()}
	if *payload.IsDelete {
		existing.DeletedBy = &mUserId
		existing.DeletedOn = &dto.JSONTime{Time: time.Now()}
		existing.IsDelete = payload.IsDelete
	}

	return s.repo.Update(existing)
}

func (s *MBiodataServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *MBiodataServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
