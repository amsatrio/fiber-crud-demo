package t_customer_va

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type TCustomerVa struct {
	Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
	VaNumber   *string       `form:"vaNumber" json:"vaNumber" xml:"vaNumber" gorm:"size:30;type:varchar(30)" validate:"max=30"`
	CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCustomerVa) TableName() string {
	return "t_customer_va"
}

type TCustomerVaRequest struct {
	Id         *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	CustomerId *uint   `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
	VaNumber   *string `form:"vaNumber" json:"vaNumber" xml:"vaNumber" gorm:"size:30;type:varchar(30)" validate:"max=30"`
	IsDelete   *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCustomerVaResponse struct {
	Id         *uint   `form:"id" json:"id" xml:"id"`
	CustomerId *uint   `form:"customerId" json:"customerId" xml:"customerId"`
	VaNumber   *string `form:"vaNumber" json:"vaNumber" xml:"vaNumber" validate:"max=30"`
	IsDelete   bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToTCustomerVaEntity(req TCustomerVaRequest, imageData []byte, userId uint) TCustomerVa {
	entity := TCustomerVa{
		CustomerId: req.CustomerId,
		VaNumber:   req.VaNumber,
		CreatedBy:  userId,
		CreatedOn:  dto.JSONTime{Time: time.Now()},
		IsDelete:   false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToTCustomerVaResponse(entity TCustomerVa) TCustomerVaResponse {
	res := TCustomerVaResponse{
		Id:         &entity.Id,
		CustomerId: entity.CustomerId,
		VaNumber:   entity.VaNumber,
		IsDelete:   entity.IsDelete,
	}
	return res
}
func ToTCustomerVaResponsesParallel(entities []TCustomerVa) []TCustomerVaResponse {
	numEntities := len(entities)
	responses := make([]TCustomerVaResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity TCustomerVa) {
			defer wg.Done()
			responses[index] = ToTCustomerVaResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
