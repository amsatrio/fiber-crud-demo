package application

import (
	"errors"
	"fiber-crud-demo/dto"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/internal/domain"
	"time"
)

type MModuleService interface {
	Get(id uint) (*domain.MModule, error)
	Create(payload *domain.MModuleRequest, mUserId uint) error
	Update(payload *domain.MModuleRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MModuleServiceImpl struct {
	repo domain.MModuleRepository
}

func NewMModuleService(repo domain.MModuleRepository) MModuleService {
	return &MModuleServiceImpl{
		repo: repo,
	}
}

func (s *MModuleServiceImpl) Get(id uint) (*domain.MModule, error) {
	return s.repo.Get(id)
}

func (s *MModuleServiceImpl) Create(payload *domain.MModuleRequest, mUserId uint) error {
	bool_true := false
	data := &domain.MModule{
		Id:        payload.Id,
		Name:      payload.Name,
		CreatedOn: dto.JSONTime{Time: time.Now()},
		CreatedBy: mUserId,
		IsDelete:  &bool_true,
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

func (s *MModuleServiceImpl) Update(payload *domain.MModuleRequest, mUserId uint) error {

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
	if *payload.IsDelete {
		existing.DeletedBy = &mUserId
		existing.DeletedOn = &dto.JSONTime{Time: time.Now()}
		existing.IsDelete = payload.IsDelete
	}

	return s.repo.Update(existing)
}

func (s *MModuleServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *MModuleServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
