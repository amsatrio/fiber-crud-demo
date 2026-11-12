package like

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type LikeService interface {
	Get(userId uint, tweetId uint) (*Like, error)
	Create(payload *LikeRequest) error
	Delete(userId uint, tweetId uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type LikeServiceImpl struct {
	repo LikeRepository
}

func NewLikeService(repo LikeRepository) LikeService {
	return &LikeServiceImpl{
		repo: repo,
	}
}

func (s *LikeServiceImpl) Get(userId uint, tweetId uint) (*Like, error) {
	existing, err := s.repo.Get(userId, tweetId)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *LikeServiceImpl) Create(payload *LikeRequest) error {
	if payload.UserId == nil || payload.TweetId == nil {
		return errors.New("invalid payload")
	}

	data := ToLikeEntity(*payload)
	return s.repo.Create(&data)
}

func (s *LikeServiceImpl) Delete(userId uint, tweetId uint) error {
	_, err := s.repo.Get(userId, tweetId)
	if err != nil {
		return err
	}
	return s.repo.Delete(userId, tweetId)
}

func (s *LikeServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
