package m_medical_item_segmentation

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type MMedicalItemSegmentationService interface {
	Get(id uint) (*MMedicalItemSegmentation, error)
	Create(payload *MMedicalItemSegmentationRequest, mUserId uint) error
	Update(payload *MMedicalItemSegmentationRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MMedicalItemSegmentationServiceImpl struct {
	repo MMedicalItemSegmentationRepository
}

func NewMMedicalItemSegmentationService(repo MMedicalItemSegmentationRepository) MMedicalItemSegmentationService {
	return &MMedicalItemSegmentationServiceImpl{
		repo: repo,
	}
}

func (s *MMedicalItemSegmentationServiceImpl) Get(id uint) (*MMedicalItemSegmentation, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *MMedicalItemSegmentationServiceImpl) Create(payload *MMedicalItemSegmentationRequest, mUserId uint) error {

	data := ToMMedicalItemSegmentationEntity(*payload, nil, mUserId)

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

func (s *MMedicalItemSegmentationServiceImpl) Update(payload *MMedicalItemSegmentationRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

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

func (s *MMedicalItemSegmentationServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MMedicalItemSegmentationServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
