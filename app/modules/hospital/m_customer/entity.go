package m_customer

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type MCustomer struct {
	Id           uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	BiodataId    *uint         `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
	Dob          *string       `form:"dob" json:"dob" xml:"dob" gorm:"type:date"`
	Gender       *string       `form:"gender" json:"gender" xml:"gender" gorm:"size:1;type:varchar(1)" validate:"max=1"`
	BloodGroupId *uint         `form:"bloodGroupId" json:"bloodGroupId" xml:"bloodGroupId" gorm:"type:bigint"`
	RhesusType   *string       `form:"rhesusType" json:"rhesusType" xml:"rhesusType" gorm:"size:5;type:varchar(5)" validate:"max=5"`
	Height       *string       `form:"height" json:"height" xml:"height" gorm:"type:decimal"`
	Weight       *string       `form:"weight" json:"weight" xml:"weight" gorm:"type:decimal"`
	CreatedBy    uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn    dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy   *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn   *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy    *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn    *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete     bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MCustomer) TableName() string {
	return "m_customer"
}

type MCustomerRequest struct {
	Id           *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	BiodataId    *uint   `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
	Dob          *string `form:"dob" json:"dob" xml:"dob" gorm:"type:date"`
	Gender       *string `form:"gender" json:"gender" xml:"gender" gorm:"size:1;type:varchar(1)" validate:"max=1"`
	BloodGroupId *uint   `form:"bloodGroupId" json:"bloodGroupId" xml:"bloodGroupId" gorm:"type:bigint"`
	RhesusType   *string `form:"rhesusType" json:"rhesusType" xml:"rhesusType" gorm:"size:5;type:varchar(5)" validate:"max=5"`
	Height       *string `form:"height" json:"height" xml:"height" gorm:"type:decimal"`
	Weight       *string `form:"weight" json:"weight" xml:"weight" gorm:"type:decimal"`
	IsDelete     *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MCustomerResponse struct {
	Id           *uint   `form:"id" json:"id" xml:"id"`
	BiodataId    *uint   `form:"biodataId" json:"biodataId" xml:"biodataId"`
	Dob          *string `form:"dob" json:"dob" xml:"dob"`
	Gender       *string `form:"gender" json:"gender" xml:"gender" validate:"max=1"`
	BloodGroupId *uint   `form:"bloodGroupId" json:"bloodGroupId" xml:"bloodGroupId"`
	RhesusType   *string `form:"rhesusType" json:"rhesusType" xml:"rhesusType" validate:"max=5"`
	Height       *string `form:"height" json:"height" xml:"height"`
	Weight       *string `form:"weight" json:"weight" xml:"weight"`
	IsDelete     bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToMCustomerEntity(req MCustomerRequest, imageData []byte, userId uint) MCustomer {
	entity := MCustomer{
		BiodataId:    req.BiodataId,
		Dob:          req.Dob,
		Gender:       req.Gender,
		BloodGroupId: req.BloodGroupId,
		RhesusType:   req.RhesusType,
		Height:       req.Height,
		Weight:       req.Weight,
		CreatedBy:    userId,
		CreatedOn:    dto.JSONTime{Time: time.Now()},
		IsDelete:     false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToMCustomerResponse(entity MCustomer) MCustomerResponse {
	res := MCustomerResponse{
		Id:           &entity.Id,
		BiodataId:    entity.BiodataId,
		Dob:          entity.Dob,
		Gender:       entity.Gender,
		BloodGroupId: entity.BloodGroupId,
		RhesusType:   entity.RhesusType,
		Height:       entity.Height,
		Weight:       entity.Weight,
		IsDelete:     entity.IsDelete,
	}
	return res
}
func ToMCustomerResponsesParallel(entities []MCustomer) []MCustomerResponse {
	numEntities := len(entities)
	responses := make([]MCustomerResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity MCustomer) {
			defer wg.Done()
			responses[index] = ToMCustomerResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
