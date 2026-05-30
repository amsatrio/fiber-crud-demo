package t_token

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type TToken struct {
	Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	Email      *string       `form:"email" json:"email" xml:"email" gorm:"size:100;type:varchar(100)" validate:"max=100"`
	UserId     *uint         `form:"userId" json:"userId" xml:"userId" gorm:"type:bigint"`
	Token      *string       `form:"token" json:"token" xml:"token" gorm:"size:50;type:varchar(50)" validate:"max=50"`
	ExpiredOn  *dto.JSONTime `form:"expiredOn" json:"expiredOn" xml:"expiredOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsExpired  *bool         `form:"isExpired" json:"isExpired" xml:"isExpired" gorm:"type:boolean"`
	UsedFor    *string       `form:"usedFor" json:"usedFor" xml:"usedFor" gorm:"size:20;type:varchar(20)" validate:"max=20"`
	CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TToken) TableName() string {
	return "t_token"
}

type TTokenRequest struct {
	Id        *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	Email     *string       `form:"email" json:"email" xml:"email" gorm:"size:100;type:varchar(100)" validate:"max=100"`
	UserId    *uint         `form:"userId" json:"userId" xml:"userId" gorm:"type:bigint"`
	Token     *string       `form:"token" json:"token" xml:"token" gorm:"size:50;type:varchar(50)" validate:"max=50"`
	ExpiredOn *dto.JSONTime `form:"expiredOn" json:"expiredOn" xml:"expiredOn" gorm:"type:datetime"`
	IsExpired *bool         `form:"isExpired" json:"isExpired" xml:"isExpired" gorm:"type:boolean"`
	UsedFor   *string       `form:"usedFor" json:"usedFor" xml:"usedFor" gorm:"size:20;type:varchar(20)" validate:"max=20"`
	IsDelete  *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TTokenResponse struct {
	Id        *uint         `form:"id" json:"id" xml:"id"`
	Email     *string       `form:"email" json:"email" xml:"email" validate:"max=100"`
	UserId    *uint         `form:"userId" json:"userId" xml:"userId"`
	Token     *string       `form:"token" json:"token" xml:"token" validate:"max=50"`
	ExpiredOn *dto.JSONTime `form:"expiredOn" json:"expiredOn" xml:"expiredOn"`
	IsExpired *bool         `form:"isExpired" json:"isExpired" xml:"isExpired"`
	UsedFor   *string       `form:"usedFor" json:"usedFor" xml:"usedFor" validate:"max=20"`
	IsDelete  bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToTTokenEntity(req TTokenRequest, imageData []byte, userId uint) TToken {
	entity := TToken{
		Email:     req.Email,
		UserId:    req.UserId,
		Token:     req.Token,
		ExpiredOn: req.ExpiredOn,
		IsExpired: req.IsExpired,
		UsedFor:   req.UsedFor,
		CreatedBy: userId,
		CreatedOn: dto.JSONTime{Time: time.Now()},
		IsDelete:  false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToTTokenResponse(entity TToken) TTokenResponse {
	res := TTokenResponse{
		Id:        &entity.Id,
		Email:     entity.Email,
		UserId:    entity.UserId,
		Token:     entity.Token,
		ExpiredOn: entity.ExpiredOn,
		IsExpired: entity.IsExpired,
		UsedFor:   entity.UsedFor,
		IsDelete:  entity.IsDelete,
	}
	return res
}
func ToTTokenResponsesParallel(entities []TToken) []TTokenResponse {
	numEntities := len(entities)
	responses := make([]TTokenResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity TToken) {
			defer wg.Done()
			responses[index] = ToTTokenResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
