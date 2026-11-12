package like

import (
	"errors"
	"strconv"
	"sync"

	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/modules/twitter_clone/tweet"
	"github.com/amsatrio/fiber-crud-demo/app/util"

	"gorm.io/gorm"
)

type LikeRepository interface {
	Get(userId uint, tweetId uint) (*Like, error)
	Create(data *Like) error
	Delete(userId uint, tweetId uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type LikeRepositoryImpl struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &LikeRepositoryImpl{
		db: db,
	}
}

func (s *LikeRepositoryImpl) Get(userId uint, tweetId uint) (*Like, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	entity := Like{}
	result := s.db.First(&entity, "user_id = ? AND tweet_id = ?", userId, tweetId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (s *LikeRepositoryImpl) Create(entity *Like) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&entity).Error; err != nil {
			return err
		}

		return tx.Model(&tweet.Tweet{}).
			Where("id = ?", entity.TweetId).
			UpdateColumn("likes_count", gorm.Expr("likes_count + 1")).Error
	})
}

func (s *LikeRepositoryImpl) Delete(userId uint, tweetId uint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&Like{}, "user_id = ? AND tweet_id = ?", userId, tweetId)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("data not found")
		}

		return tx.Model(&tweet.Tweet{}).
			Where("id = ? AND likes_count > 0", tweetId).
			UpdateColumn("likes_count", gorm.Expr("likes_count - 1")).Error
	})
}

func (s *LikeRepositoryImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	util.Log("INFO", "repository", "GetPageLike", "")

	var entities []Like
	var entity Like
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
		Content:          ToLikeResponsesParallel(entities),
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

	util.Log("INFO", "repository", "GetPageLike", "sort is empty: "+strconv.FormatBool(sort.Empty))

	return &page, nil
}
