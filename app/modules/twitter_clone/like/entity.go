package like

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type Like struct {
	UserId    uint         `form:"userId" json:"userId" xml:"userId" gorm:"primaryKey;not null;type:bigint"`
	TweetId   uint         `form:"tweetId" json:"tweetId" xml:"tweetId" gorm:"primaryKey;not null;type:bigint"`
	CreatedAt dto.JSONTime `form:"createdAt" json:"createdAt" xml:"createdAt" gorm:"not null;type:timestamp" swaggertype:"string" example:"2024-02-16 10:33:10"`
}

func (Like) TableName() string {
	return "likes"
}

type LikeRequest struct {
	UserId  *uint `form:"userId" json:"userId" xml:"userId" gorm:"not null;type:bigint" validate:"required"`
	TweetId *uint `form:"tweetId" json:"tweetId" xml:"tweetId" gorm:"not null;type:bigint" validate:"required"`
}
type LikeResponse struct {
	UserId    *uint        `form:"userId" json:"userId" xml:"userId"`
	TweetId   *uint        `form:"tweetId" json:"tweetId" xml:"tweetId"`
	CreatedAt dto.JSONTime `form:"createdAt" json:"createdAt" xml:"createdAt"`
}

func ToLikeEntity(req LikeRequest) Like {
	entity := Like{
		CreatedAt: dto.JSONTime{Time: time.Now()},
	}
	if req.UserId != nil {
		entity.UserId = *req.UserId
	}
	if req.TweetId != nil {
		entity.TweetId = *req.TweetId
	}
	return entity
}

func ToLikeResponse(entity Like) LikeResponse {
	res := LikeResponse{
		UserId:    &entity.UserId,
		TweetId:   &entity.TweetId,
		CreatedAt: entity.CreatedAt,
	}
	return res
}
func ToLikeResponsesParallel(entities []Like) []LikeResponse {
	numEntities := len(entities)
	responses := make([]LikeResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity Like) {
			defer wg.Done()
			responses[index] = ToLikeResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
