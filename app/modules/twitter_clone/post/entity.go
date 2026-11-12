package post

import (
	"encoding/base64"
	"mime/multipart"
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type Post struct {
	Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	UserId     *uint         `form:"userId" json:"userId" xml:"userId" gorm:"type:bigint" validate:"required"`
	Message    *[]byte       `form:"message" json:"message" xml:"message" gorm:"type:blob"`
	CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (Post) TableName() string {
	return "post"
}

type PostRequest struct {
	Id       *uint                 `form:"id" json:"id" xml:"id" validate:"required"`
	UserId   *uint                 `form:"userId" json:"userId" xml:"userId"`
	Message  *multipart.FileHeader `form:"message" json:"message" xml:"message" gorm:"type:blob"`
	IsDelete *bool                 `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
type PostResponse struct {
	Id       *uint   `form:"id" json:"id" xml:"id" validate:"required"`
	UserId   *uint   `form:"userId" json:"userId" xml:"userId"`
	Message  *string `form:"message" json:"message" xml:"message" gorm:"type:blob"`
	IsDelete bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToPostEntity(req PostRequest, imageData []byte, userId uint) Post {
	entity := Post{
		UserId:    req.UserId,
		Message:   &imageData,
		CreatedBy: userId,
		CreatedOn: dto.JSONTime{Time: time.Now()},
		IsDelete:  false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToPostResponse(entity Post) PostResponse {
	res := PostResponse{
		Id:       &entity.Id,
		UserId:   entity.UserId,
		Message:  nil,
		IsDelete: entity.IsDelete,
	}
	if entity.Message != nil && len(*entity.Message) > 0 {
		encoded := base64.StdEncoding.EncodeToString(*entity.Message)
		res.Message = &encoded
	}
	return res
}
func ToPostResponsesParallel(entities []Post) []PostResponse {
	numEntities := len(entities)
	responses := make([]PostResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity Post) {
			defer wg.Done()
			responses[index] = ToPostResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
