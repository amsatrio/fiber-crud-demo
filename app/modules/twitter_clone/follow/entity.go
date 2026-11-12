package follow

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type Follow struct {
	FollowerId uint         `form:"followerId" json:"followerId" xml:"followerId" gorm:"primaryKey;not null;type:bigint"`
	FollowedId uint         `form:"followedId" json:"followedId" xml:"followedId" gorm:"primaryKey;not null;type:bigint"`
	CreatedAt  dto.JSONTime `form:"createdAt" json:"createdAt" xml:"createdAt" gorm:"not null;type:timestamp" swaggertype:"string" example:"2024-02-16 10:33:10"`
}

func (Follow) TableName() string {
	return "follows"
}

type FollowRequest struct {
	FollowerId *uint `form:"followerId" json:"followerId" xml:"followerId" gorm:"not null;type:bigint" validate:"required"`
	FollowedId *uint `form:"followedId" json:"followedId" xml:"followedId" gorm:"not null;type:bigint" validate:"required"`
}
type FollowResponse struct {
	FollowerId *uint        `form:"followerId" json:"followerId" xml:"followerId"`
	FollowedId *uint        `form:"followedId" json:"followedId" xml:"followedId"`
	CreatedAt  dto.JSONTime `form:"createdAt" json:"createdAt" xml:"createdAt"`
}

func ToFollowEntity(req FollowRequest) Follow {
	entity := Follow{
		CreatedAt: dto.JSONTime{Time: time.Now()},
	}
	if req.FollowerId != nil {
		entity.FollowerId = *req.FollowerId
	}
	if req.FollowedId != nil {
		entity.FollowedId = *req.FollowedId
	}
	return entity
}

func ToFollowResponse(entity Follow) FollowResponse {
	res := FollowResponse{
		FollowerId: &entity.FollowerId,
		FollowedId: &entity.FollowedId,
		CreatedAt:  entity.CreatedAt,
	}
	return res
}
func ToFollowResponsesParallel(entities []Follow) []FollowResponse {
	numEntities := len(entities)
	responses := make([]FollowResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity Follow) {
			defer wg.Done()
			responses[index] = ToFollowResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
