package follow

import (
	"errors"
	"strconv"
	"sync"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/util"

	"gorm.io/gorm"
)

type FollowRepository interface {
	Get(followerId uint, followedId uint) (*Follow, error)
	Create(data *Follow) error
	Delete(followerId uint, followedId uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type FollowRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &FollowRepositoryImpl{
		db: db,
	}
}

func (s *FollowRepositoryImpl) Get(followerId uint, followedId uint) (*Follow, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	entity := Follow{}
	result := s.db.First(&entity, "follower_id = ? AND followed_id = ?", followerId, followedId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (s *FollowRepositoryImpl) Create(entity *Follow) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *FollowRepositoryImpl) Delete(followerId uint, followedId uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := s.db.Delete(&Follow{}, "follower_id = ? AND followed_id = ?", followerId, followedId)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (s *FollowRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "repository", "GetPageFollow", "")

	var entities []Follow
	var entity Follow
	fieldMap := util.GetJSONFieldTypes(entity)

	db := s.db
	if len(sortRequest) > 0 {
		db = util.ApplySorting(db, sortRequest)
	} else {
		db = db.Order("created_at DESC")
	}
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
		Content:          ToFollowResponsesParallel(entities),
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

	util.Log("INFO", "repository", "GetPageFollow", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
