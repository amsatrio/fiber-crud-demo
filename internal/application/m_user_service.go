package application

import (
	"errors"
	"fiber-crud-demo/dto"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/internal/domain"
	"time"
)

type MUserService interface {
	Get(id uint) (*domain.MUser, error)
	Create(payload *domain.MUserRequest, mUserId uint) error
	Update(payload *domain.MUserRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MUserServiceImpl struct {
	repo domain.MUserRepository
}

func NewMUserService(repo domain.MUserRepository) MUserService {
	return &MUserServiceImpl{
		repo: repo,
	}
}

func (s *MUserServiceImpl) Get(id uint) (*domain.MUser, error) {
	return s.repo.Get(id)
}

func (s *MUserServiceImpl) Create(payload *domain.MUserRequest, mUserId uint) error {
	bool_true := false
	data := &domain.MUser{
		Id:           payload.Id,
		BiodataId:    payload.BiodataId,
		RoleId:       payload.RoleId,
		Email:        payload.Email,
		Password:     payload.Password,
		LoginAttempt: payload.LoginAttempt,
		IsLocked:     payload.IsLocked,
		LastLogin:    payload.LastLogin,
		CreatedOn:    dto.JSONTime{Time: time.Now()},
		CreatedBy:    mUserId,
		IsDelete:     &bool_true,
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

func (s *MUserServiceImpl) Update(payload *domain.MUserRequest, mUserId uint) error {

	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.BiodataId = payload.BiodataId
	existing.RoleId = payload.RoleId
	existing.Email = payload.Email
	existing.Password = payload.Password
	existing.LoginAttempt = payload.LoginAttempt
	existing.IsLocked = payload.IsLocked
	existing.LastLogin = payload.LastLogin
	existing.ModifiedBy = &mUserId
	existing.ModifiedOn = &dto.JSONTime{Time: time.Now()}
	if *payload.IsDelete {
		existing.DeletedBy = &mUserId
		existing.DeletedOn = &dto.JSONTime{Time: time.Now()}
		existing.IsDelete = payload.IsDelete
	}

	return s.repo.Update(existing)
}

func (s *MUserServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *MUserServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
