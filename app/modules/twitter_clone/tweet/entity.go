package tweet

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type Tweet struct {
	Id            uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	UserId        *uint        `form:"userId" json:"userId" xml:"userId" gorm:"not null;type:bigint" validate:"required"`
	Content       *string      `form:"content" json:"content" xml:"content" gorm:"size:280;type:varchar(280)" validate:"max=280"`
	ParentTweetId *uint        `form:"parentTweetId" json:"parentTweetId" xml:"parentTweetId" gorm:"type:bigint"`
	LikesCount    *uint        `form:"likesCount" json:"likesCount" xml:"likesCount" gorm:"type:int unsigned"`
	RetweetsCount *uint        `form:"retweetsCount" json:"retweetsCount" xml:"retweetsCount" gorm:"type:int unsigned"`
	RepliesCount  *uint        `form:"repliesCount" json:"repliesCount" xml:"repliesCount" gorm:"type:int unsigned"`
	CreatedAt     dto.JSONTime `form:"createdAt" json:"createdAt" xml:"createdAt" gorm:"not null;type:timestamp" swaggertype:"string" example:"2024-02-16 10:33:10"`
	Username      *string      `form:"username" json:"username" xml:"username" gorm:"-"`
}

func (Tweet) TableName() string {
	return "tweets"
}

type TweetRequest struct {
	Id            *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	UserId        *uint   `form:"userId" json:"userId" xml:"userId" gorm:"not null;type:bigint" validate:"required"`
	Content       *string `form:"content" json:"content" xml:"content" gorm:"size:280;type:varchar(280)" validate:"required,max=280"`
	ParentTweetId *uint   `form:"parentTweetId" json:"parentTweetId" xml:"parentTweetId" gorm:"type:bigint"`
	LikesCount    *uint   `form:"likesCount" json:"likesCount" xml:"likesCount" gorm:"type:int unsigned"`
	RetweetsCount *uint   `form:"retweetsCount" json:"retweetsCount" xml:"retweetsCount" gorm:"type:int unsigned"`
	RepliesCount  *uint   `form:"repliesCount" json:"repliesCount" xml:"repliesCount" gorm:"type:int unsigned"`
}
type TweetResponse struct {
	Id            *uint        `form:"id" json:"id" xml:"id"`
	UserId        *uint        `form:"userId" json:"userId" xml:"userId"`
	Content       *string      `form:"content" json:"content" xml:"content" validate:"max=280"`
	ParentTweetId *uint        `form:"parentTweetId" json:"parentTweetId" xml:"parentTweetId"`
	LikesCount    *uint        `form:"likesCount" json:"likesCount" xml:"likesCount"`
	RetweetsCount *uint        `form:"retweetsCount" json:"retweetsCount" xml:"retweetsCount"`
	RepliesCount  *uint        `form:"repliesCount" json:"repliesCount" xml:"repliesCount"`
	CreatedAt     dto.JSONTime `form:"createdAt" json:"createdAt" xml:"createdAt"`
	Username      *string      `form:"username" json:"username" xml:"username"`
}

func ToTweetEntity(req TweetRequest) Tweet {
	entity := Tweet{
		UserId:        req.UserId,
		Content:       req.Content,
		ParentTweetId: req.ParentTweetId,
		LikesCount:    req.LikesCount,
		RetweetsCount: req.RetweetsCount,
		RepliesCount:  req.RepliesCount,
		CreatedAt:     dto.JSONTime{Time: time.Now()},
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	zero := uint(0)
	if entity.LikesCount == nil {
		entity.LikesCount = &zero
	}
	if entity.RetweetsCount == nil {
		entity.RetweetsCount = &zero
	}
	if entity.RepliesCount == nil {
		entity.RepliesCount = &zero
	}
	return entity
}

func ToTweetResponse(entity Tweet) TweetResponse {
	res := TweetResponse{
		Id:            &entity.Id,
		UserId:        entity.UserId,
		Content:       entity.Content,
		ParentTweetId: entity.ParentTweetId,
		LikesCount:    entity.LikesCount,
		RetweetsCount: entity.RetweetsCount,
		RepliesCount:  entity.RepliesCount,
		CreatedAt:     entity.CreatedAt,
		Username:      entity.Username,
	}
	return res
}
func ToTweetResponsesParallel(entities []Tweet) []TweetResponse {
	numEntities := len(entities)
	responses := make([]TweetResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity Tweet) {
			defer wg.Done()
			responses[index] = ToTweetResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
