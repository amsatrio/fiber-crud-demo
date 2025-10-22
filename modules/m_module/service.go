package m_module

import (
	"errors"
	"fiber-crud-demo/dto"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"time"
)

type MModuleService interface {
	Get(id uint) (*MModule, error)
	Create(payload *MModuleRequest, mUserId uint) error
	Update(payload *MModuleRequest, mUserId uint) error
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
	repo MModuleRepository
}

func NewMModuleService(repo MModuleRepository) MModuleService {
	return &MModuleServiceImpl{
		repo: repo,
	}
}

func (s *MModuleServiceImpl) Get(id uint) (*MModule, error) {
	return s.repo.Get(id)
}

func (s *MModuleServiceImpl) Create(payload *MModuleRequest, mUserId uint) error {
	bool_true := false
	data := &MModule{
		Id:        0,
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

	data.Id = *payload.Id

	return s.repo.Create(data)
}

func (s *MModuleServiceImpl) Update(payload *MModuleRequest, mUserId uint) error {

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
	existing.IsDelete = payload.IsDelete
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
