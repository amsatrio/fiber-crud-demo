package t_courier_discount

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type TCourierDiscount struct {
	Id            uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	CourierTypeId *uint         `form:"courierTypeId" json:"courierTypeId" xml:"courierTypeId" gorm:"type:bigint"`
	Value         *string       `form:"value" json:"value" xml:"value" gorm:"type:decimal"`
	CreatedBy     uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn     dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy    *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn    *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy     *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn     *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete      bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCourierDiscount) TableName() string {
	return "t_courier_discount"
}

type TCourierDiscountRequest struct {
	Id            *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	CourierTypeId *uint   `form:"courierTypeId" json:"courierTypeId" xml:"courierTypeId" gorm:"type:bigint"`
	Value         *string `form:"value" json:"value" xml:"value" gorm:"type:decimal"`
	IsDelete      *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCourierDiscountResponse struct {
	Id            *uint   `form:"id" json:"id" xml:"id"`
	CourierTypeId *uint   `form:"courierTypeId" json:"courierTypeId" xml:"courierTypeId"`
	Value         *string `form:"value" json:"value" xml:"value"`
	IsDelete      bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToTCourierDiscountEntity(req TCourierDiscountRequest, imageData []byte, userId uint) TCourierDiscount {
	entity := TCourierDiscount{
		CourierTypeId: req.CourierTypeId,
		Value:         req.Value,
		CreatedBy:     userId,
		CreatedOn:     dto.JSONTime{Time: time.Now()},
		IsDelete:      false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToTCourierDiscountResponse(entity TCourierDiscount) TCourierDiscountResponse {
	res := TCourierDiscountResponse{
		Id:            &entity.Id,
		CourierTypeId: entity.CourierTypeId,
		Value:         entity.Value,
		IsDelete:      entity.IsDelete,
	}
	return res
}
func ToTCourierDiscountResponsesParallel(entities []TCourierDiscount) []TCourierDiscountResponse {
	numEntities := len(entities)
	responses := make([]TCourierDiscountResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity TCourierDiscount) {
			defer wg.Done()
			responses[index] = ToTCourierDiscountResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
