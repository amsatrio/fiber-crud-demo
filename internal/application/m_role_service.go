package application

import (
	"errors"
	"fiber-crud-demo/dto"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/internal/domain"
	"time"
)

type MRoleService interface {
	Get(id uint) (*domain.MRole, error)
	Create(payload *domain.MRoleRequest, mUserId uint) error
	Update(payload *domain.MRoleRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MRoleServiceImpl struct {
	repo domain.MRoleRepository
}

func NewMRoleService(repo domain.MRoleRepository) MRoleService {
	return &MRoleServiceImpl{
		repo: repo,
	}
}

func (s *MRoleServiceImpl) Get(id uint) (*domain.MRole, error) {
	return s.repo.Get(id)
}

func (s *MRoleServiceImpl) Create(payload *domain.MRoleRequest, mUserId uint) error {
	bool_true := false
	data := &domain.MRole{
		Id:        payload.Id,
		Name:      payload.Name,
		Code:      payload.Code,
		Level:     payload.Level,
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

func (s *MRoleServiceImpl) Update(payload *domain.MRoleRequest, mUserId uint) error {

	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.Name = payload.Name
	existing.Code = payload.Code
	existing.Level = payload.Level
	existing.ModifiedBy = &mUserId
	existing.ModifiedOn = &dto.JSONTime{Time: time.Now()}
	if *payload.IsDelete {
		existing.DeletedBy = &mUserId
		existing.DeletedOn = &dto.JSONTime{Time: time.Now()}
		existing.IsDelete = payload.IsDelete
	}

	return s.repo.Update(existing)
}

func (s *MRoleServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *MRoleServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
