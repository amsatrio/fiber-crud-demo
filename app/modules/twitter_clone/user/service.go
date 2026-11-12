package user

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type UserService interface {
	Get(id uint) (*User, error)
	Create(payload *UserRequest, userId uint) error
	Update(payload *UserRequest, userId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type UserServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) Get(id uint) (*User, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *UserServiceImpl) Create(payload *UserRequest, userId uint) error {

	data := ToUserEntity(*payload, userId)

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

func (s *UserServiceImpl) Update(payload *UserRequest, userId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.Username = payload.Username
	existing.Email = payload.Email
	existing.DisplayName = payload.DisplayName
	existing.PasswordHash = payload.PasswordHash
	existing.Bio = payload.Bio
	existing.ProfileImageUrl = payload.ProfileImageUrl

	return s.repo.Update(existing)
}

func (s *UserServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *UserServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
