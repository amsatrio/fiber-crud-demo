package t_customer_registered_card

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type TCustomerRegisteredCard struct {
	Id             uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	CustomerId     *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
	CardNumber     *string       `form:"cardNumber" json:"cardNumber" xml:"cardNumber" gorm:"size:20;type:varchar(20)" validate:"max=20"`
	ValidityPeriod *string       `form:"validityPeriod" json:"validityPeriod" xml:"validityPeriod" gorm:"type:date"`
	Cvv            *string       `form:"cvv" json:"cvv" xml:"cvv" gorm:"size:5;type:varchar(5)" validate:"max=5"`
	CreatedBy      uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn      dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy     *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn     *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy      *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn      *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete       bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCustomerRegisteredCard) TableName() string {
	return "t_customer_registered_card"
}

type TCustomerRegisteredCardRequest struct {
	Id             *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	CustomerId     *uint   `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
	CardNumber     *string `form:"cardNumber" json:"cardNumber" xml:"cardNumber" gorm:"size:20;type:varchar(20)" validate:"max=20"`
	ValidityPeriod *string `form:"validityPeriod" json:"validityPeriod" xml:"validityPeriod" gorm:"type:date"`
	Cvv            *string `form:"cvv" json:"cvv" xml:"cvv" gorm:"size:5;type:varchar(5)" validate:"max=5"`
	IsDelete       *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCustomerRegisteredCardResponse struct {
	Id             *uint   `form:"id" json:"id" xml:"id"`
	CustomerId     *uint   `form:"customerId" json:"customerId" xml:"customerId"`
	CardNumber     *string `form:"cardNumber" json:"cardNumber" xml:"cardNumber" validate:"max=20"`
	ValidityPeriod *string `form:"validityPeriod" json:"validityPeriod" xml:"validityPeriod"`
	Cvv            *string `form:"cvv" json:"cvv" xml:"cvv" validate:"max=5"`
	IsDelete       bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToTCustomerRegisteredCardEntity(req TCustomerRegisteredCardRequest, imageData []byte, userId uint) TCustomerRegisteredCard {
	entity := TCustomerRegisteredCard{
		CustomerId:     req.CustomerId,
		CardNumber:     req.CardNumber,
		ValidityPeriod: req.ValidityPeriod,
		Cvv:            req.Cvv,
		CreatedBy:      userId,
		CreatedOn:      dto.JSONTime{Time: time.Now()},
		IsDelete:       false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToTCustomerRegisteredCardResponse(entity TCustomerRegisteredCard) TCustomerRegisteredCardResponse {
	res := TCustomerRegisteredCardResponse{
		Id:             &entity.Id,
		CustomerId:     entity.CustomerId,
		CardNumber:     entity.CardNumber,
		ValidityPeriod: entity.ValidityPeriod,
		Cvv:            entity.Cvv,
		IsDelete:       entity.IsDelete,
	}
	return res
}
func ToTCustomerRegisteredCardResponsesParallel(entities []TCustomerRegisteredCard) []TCustomerRegisteredCardResponse {
	numEntities := len(entities)
	responses := make([]TCustomerRegisteredCardResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity TCustomerRegisteredCard) {
			defer wg.Done()
			responses[index] = ToTCustomerRegisteredCardResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
