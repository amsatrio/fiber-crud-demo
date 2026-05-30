package m_medical_item_category

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type MMedicalItemCategoryService interface {
	Get(id uint) (*MMedicalItemCategory, error)
	Create(payload *MMedicalItemCategoryRequest, mUserId uint) error
	Update(payload *MMedicalItemCategoryRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MMedicalItemCategoryServiceImpl struct {
	repo MMedicalItemCategoryRepository
}

func NewMMedicalItemCategoryService(repo MMedicalItemCategoryRepository) MMedicalItemCategoryService {
	return &MMedicalItemCategoryServiceImpl{
		repo: repo,
	}
}

func (s *MMedicalItemCategoryServiceImpl) Get(id uint) (*MMedicalItemCategory, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *MMedicalItemCategoryServiceImpl) Create(payload *MMedicalItemCategoryRequest, mUserId uint) error {

	data := ToMMedicalItemCategoryEntity(*payload, nil, mUserId)

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

func (s *MMedicalItemCategoryServiceImpl) Update(payload *MMedicalItemCategoryRequest, mUserId uint) error {
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

func (s *MMedicalItemCategoryServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MMedicalItemCategoryServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
