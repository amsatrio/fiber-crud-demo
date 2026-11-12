package tweet

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type TweetService interface {
	Get(id uint) (*Tweet, error)
	Create(payload *TweetRequest, userId uint) error
	Update(payload *TweetRequest, userId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TweetServiceImpl struct {
	repo TweetRepository
}

func NewTweetService(repo TweetRepository) TweetService {
	return &TweetServiceImpl{
		repo: repo,
	}
}

func (s *TweetServiceImpl) Get(id uint) (*Tweet, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TweetServiceImpl) Create(payload *TweetRequest, userId uint) error {

	data := ToTweetEntity(*payload)

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

func (s *TweetServiceImpl) Update(payload *TweetRequest, userId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.UserId = payload.UserId
	existing.Content = payload.Content
	existing.ParentTweetId = payload.ParentTweetId
	if payload.LikesCount != nil {
		existing.LikesCount = payload.LikesCount
	}
	if payload.RetweetsCount != nil {
		existing.RetweetsCount = payload.RetweetsCount
	}
	if payload.RepliesCount != nil {
		existing.RepliesCount = payload.RepliesCount
	}

	return s.repo.Update(existing)
}

func (s *TweetServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TweetServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
