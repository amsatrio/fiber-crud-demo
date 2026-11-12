package follow

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type FollowService interface {
	Get(followerId uint, followedId uint) (*Follow, error)
	Create(payload *FollowRequest) error
	Delete(followerId uint, followedId uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type FollowServiceImpl struct {
	repo FollowRepository
}

func NewFollowService(repo FollowRepository) FollowService {
	return &FollowServiceImpl{
		repo: repo,
	}
}

func (s *FollowServiceImpl) Get(followerId uint, followedId uint) (*Follow, error) {
	existing, err := s.repo.Get(followerId, followedId)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *FollowServiceImpl) Create(payload *FollowRequest) error {
	if payload.FollowerId == nil || payload.FollowedId == nil {
		return errors.New("invalid payload")
	}

	data := ToFollowEntity(*payload)
	return s.repo.Create(&data)
}

func (s *FollowServiceImpl) Delete(followerId uint, followedId uint) error {
	_, err := s.repo.Get(followerId, followedId)
	if err != nil {
		return err
	}
	return s.repo.Delete(followerId, followedId)
}

func (s *FollowServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
