package m_specialization

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type MSpecialization struct {
	Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
	CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MSpecialization) TableName() string {
	return "m_specialization"
}

type MSpecializationRequest struct {
	Id       *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	Name     *string `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
	IsDelete *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MSpecializationResponse struct {
	Id       *uint   `form:"id" json:"id" xml:"id"`
	Name     *string `form:"name" json:"name" xml:"name" validate:"max=50"`
	IsDelete bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToMSpecializationEntity(req MSpecializationRequest, imageData []byte, userId uint) MSpecialization {
	entity := MSpecialization{
		Name:      req.Name,
		CreatedBy: userId,
		CreatedOn: dto.JSONTime{Time: time.Now()},
		IsDelete:  false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToMSpecializationResponse(entity MSpecialization) MSpecializationResponse {
	res := MSpecializationResponse{
		Id:       &entity.Id,
		Name:     entity.Name,
		IsDelete: entity.IsDelete,
	}
	return res
}
func ToMSpecializationResponsesParallel(entities []MSpecialization) []MSpecializationResponse {
	numEntities := len(entities)
	responses := make([]MSpecializationResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity MSpecialization) {
			defer wg.Done()
			responses[index] = ToMSpecializationResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
