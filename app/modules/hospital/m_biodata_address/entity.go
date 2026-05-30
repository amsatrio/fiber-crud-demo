package m_biodata_address

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type MBiodataAddress struct {
	Id                   uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	BiodataId            *uint         `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
	Label                *string       `form:"label" json:"label" xml:"label" gorm:"size:100;type:varchar(100)" validate:"max=100"`
	Recipient            *string       `form:"recipient" json:"recipient" xml:"recipient" gorm:"size:100;type:varchar(100)" validate:"max=100"`
	RecipientPhoneNumber *string       `form:"recipientPhoneNumber" json:"recipientPhoneNumber" xml:"recipientPhoneNumber" gorm:"size:15;type:varchar(15)" validate:"max=15"`
	LocationId           *uint         `form:"locationId" json:"locationId" xml:"locationId" gorm:"type:bigint"`
	PostalCode           *string       `form:"postalCode" json:"postalCode" xml:"postalCode" gorm:"size:10;type:varchar(10)" validate:"max=10"`
	Address              *string       `form:"address" json:"address" xml:"address" gorm:"type:text"`
	CreatedBy            uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn            dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy           *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn           *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy            *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn            *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete             bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MBiodataAddress) TableName() string {
	return "m_biodata_address"
}

type MBiodataAddressRequest struct {
	Id                   *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	BiodataId            *uint   `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
	Label                *string `form:"label" json:"label" xml:"label" gorm:"size:100;type:varchar(100)" validate:"max=100"`
	Recipient            *string `form:"recipient" json:"recipient" xml:"recipient" gorm:"size:100;type:varchar(100)" validate:"max=100"`
	RecipientPhoneNumber *string `form:"recipientPhoneNumber" json:"recipientPhoneNumber" xml:"recipientPhoneNumber" gorm:"size:15;type:varchar(15)" validate:"max=15"`
	LocationId           *uint   `form:"locationId" json:"locationId" xml:"locationId" gorm:"type:bigint"`
	PostalCode           *string `form:"postalCode" json:"postalCode" xml:"postalCode" gorm:"size:10;type:varchar(10)" validate:"max=10"`
	Address              *string `form:"address" json:"address" xml:"address" gorm:"type:text"`
	IsDelete             *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MBiodataAddressResponse struct {
	Id                   *uint   `form:"id" json:"id" xml:"id"`
	BiodataId            *uint   `form:"biodataId" json:"biodataId" xml:"biodataId"`
	Label                *string `form:"label" json:"label" xml:"label" validate:"max=100"`
	Recipient            *string `form:"recipient" json:"recipient" xml:"recipient" validate:"max=100"`
	RecipientPhoneNumber *string `form:"recipientPhoneNumber" json:"recipientPhoneNumber" xml:"recipientPhoneNumber" validate:"max=15"`
	LocationId           *uint   `form:"locationId" json:"locationId" xml:"locationId"`
	PostalCode           *string `form:"postalCode" json:"postalCode" xml:"postalCode" validate:"max=10"`
	Address              *string `form:"address" json:"address" xml:"address"`
	IsDelete             bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToMBiodataAddressEntity(req MBiodataAddressRequest, imageData []byte, userId uint) MBiodataAddress {
	entity := MBiodataAddress{
		BiodataId:            req.BiodataId,
		Label:                req.Label,
		Recipient:            req.Recipient,
		RecipientPhoneNumber: req.RecipientPhoneNumber,
		LocationId:           req.LocationId,
		PostalCode:           req.PostalCode,
		Address:              req.Address,
		CreatedBy:            userId,
		CreatedOn:            dto.JSONTime{Time: time.Now()},
		IsDelete:             false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToMBiodataAddressResponse(entity MBiodataAddress) MBiodataAddressResponse {
	res := MBiodataAddressResponse{
		Id:                   &entity.Id,
		BiodataId:            entity.BiodataId,
		Label:                entity.Label,
		Recipient:            entity.Recipient,
		RecipientPhoneNumber: entity.RecipientPhoneNumber,
		LocationId:           entity.LocationId,
		PostalCode:           entity.PostalCode,
		Address:              entity.Address,
		IsDelete:             entity.IsDelete,
	}
	return res
}
func ToMBiodataAddressResponsesParallel(entities []MBiodataAddress) []MBiodataAddressResponse {
	numEntities := len(entities)
	responses := make([]MBiodataAddressResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity MBiodataAddress) {
			defer wg.Done()
			responses[index] = ToMBiodataAddressResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
