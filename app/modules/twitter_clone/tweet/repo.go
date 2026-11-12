package tweet

import (
	"errors"
	"strconv"
	"sync"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/util"

	"gorm.io/gorm"
)

type TweetRepository interface {
	Get(id uint) (*Tweet, error)
	Create(data *Tweet) error
	Update(data *Tweet) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TweetRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewTweetRepository(db *gorm.DB) TweetRepository {
	return &TweetRepositoryImpl{
		db: db,
	}
}

func (s *TweetRepositoryImpl) Get(id uint) (*Tweet, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	entity := Tweet{}
	result := s.db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}

	list := []Tweet{entity}
	s.attachUsernames(list)
	return &list[0], nil
}

func (s *TweetRepositoryImpl) usernamesFor(userIds []uint) map[uint]string {
	type row struct {
		Id       uint
		Username *string
	}

	var rows []row
	s.db.Table("users").Select("id, username").Where("id IN ?", userIds).Scan(&rows)

	names := make(map[uint]string, len(rows))
	for _, r := range rows {
		if r.Username != nil {
			names[r.Id] = *r.Username
		}
	}
	return names
}

func (s *TweetRepositoryImpl) attachUsernames(entities []Tweet) {
	userIds := make([]uint, 0, len(entities))
	for _, e := range entities {
		if e.UserId != nil {
			userIds = append(userIds, *e.UserId)
		}
	}
	if len(userIds) == 0 {
		return
	}

	names := s.usernamesFor(userIds)
	for i := range entities {
		if entities[i].UserId == nil {
			continue
		}
		if name, ok := names[*entities[i].UserId]; ok {
			entities[i].Username = &name
		}
	}
}

func (s *TweetRepositoryImpl) Create(entity *Tweet) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *TweetRepositoryImpl) Update(entity *Tweet) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Model(&entity).Updates(entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *TweetRepositoryImpl) Delete(id uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var entity Tweet
	result := s.db.Delete(&entity, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *TweetRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "repository", "GetPageTweet", "")

	var entities []Tweet
	var entity Tweet
	fieldMap := util.GetJSONFieldTypes(entity)

	db := s.db
	db = util.ApplySorting(db, sortRequest)
	db = util.ApplyFiltering(db, filterRequest)
	db = util.ApplyGlobalSearch(db, searchRequest, fieldMap)

	totalElements := db.Find(&entities).RowsAffected

	totalPages := totalElements / sizeInt64
	if totalElements%sizeInt64 != 0 {
		totalPages++
	}

	result := db.Scopes(util.ApplyPaginate(pageInt, sizeInt)).Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	s.attachUsernames(entities)

	lastPage := int64(pageInt) == totalPages-1
	firstPage := pageInt == 0

	sort := response.Sort{
		Empty:    totalElements <= 0,
		Sorted:   true,
		Unsorted: false,
	}

	pageable := response.Pageable{
		Offset:     pageInt * sizeInt,
		PageNumber: pageInt,
		PageSize:   sizeInt,
		Paged:      true,
		UnPaged:    false,
		Sort:       sort,
	}

	page := response.Page{
		Content:          ToTweetResponsesParallel(entities),
		Pageable:         pageable,
		Sort:             sort,
		TotalPages:       totalPages,
		TotalElements:    totalElements,
		Size:             sizeInt,
		Number:           pageInt,
		NumberOfElements: sizeInt,
		Last:             lastPage,
		First:            firstPage,
		Empty:            sort.Empty,
	}

	util.Log("INFO", "repository", "GetPageTweet", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
