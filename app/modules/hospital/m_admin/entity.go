package m_admin

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type MAdmin struct {
	Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
	Code       *string       `form:"code" json:"code" xml:"code" gorm:"size:10;type:varchar(10)" validate:"max=10"`
	CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MAdmin) TableName() string {
	return "m_admin"
}

type MAdminRequest struct {
	Id        *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	BiodataId *uint   `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
	Code      *string `form:"code" json:"code" xml:"code" gorm:"size:10;type:varchar(10)" validate:"max=10"`
	IsDelete  *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MAdminResponse struct {
	Id        *uint   `form:"id" json:"id" xml:"id"`
	BiodataId *uint   `form:"biodataId" json:"biodataId" xml:"biodataId"`
	Code      *string `form:"code" json:"code" xml:"code" validate:"max=10"`
	IsDelete  bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToMAdminEntity(req MAdminRequest, imageData []byte, userId uint) MAdmin {
	entity := MAdmin{
		BiodataId: req.BiodataId,
		Code:      req.Code,
		CreatedBy: userId,
		CreatedOn: dto.JSONTime{Time: time.Now()},
		IsDelete:  false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToMAdminResponse(entity MAdmin) MAdminResponse {
	res := MAdminResponse{
		Id:        &entity.Id,
		BiodataId: entity.BiodataId,
		Code:      entity.Code,
		IsDelete:  entity.IsDelete,
	}
	return res
}
func ToMAdminResponsesParallel(entities []MAdmin) []MAdminResponse {
	numEntities := len(entities)
	responses := make([]MAdminResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity MAdmin) {
			defer wg.Done()
			responses[index] = ToMAdminResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
