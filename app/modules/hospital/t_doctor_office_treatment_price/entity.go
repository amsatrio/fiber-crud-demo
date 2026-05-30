package t_doctor_office_treatment_price

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type TDoctorOfficeTreatmentPrice struct {
	Id                      uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	DoctorOfficeTreatmentId *uint         `form:"doctorOfficeTreatmentId" json:"doctorOfficeTreatmentId" xml:"doctorOfficeTreatmentId" gorm:"type:bigint"`
	Price                   *string       `form:"price" json:"price" xml:"price" gorm:"type:decimal"`
	PriceStartFrom          *string       `form:"priceStartFrom" json:"priceStartFrom" xml:"priceStartFrom" gorm:"type:decimal"`
	PriceUntilFrom          *string       `form:"priceUntilFrom" json:"priceUntilFrom" xml:"priceUntilFrom" gorm:"type:decimal"`
	CreatedBy               uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn               dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy              *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn              *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy               *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn               *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete                bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TDoctorOfficeTreatmentPrice) TableName() string {
	return "t_doctor_office_treatment_price"
}

type TDoctorOfficeTreatmentPriceRequest struct {
	Id                      *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	DoctorOfficeTreatmentId *uint   `form:"doctorOfficeTreatmentId" json:"doctorOfficeTreatmentId" xml:"doctorOfficeTreatmentId" gorm:"type:bigint"`
	Price                   *string `form:"price" json:"price" xml:"price" gorm:"type:decimal"`
	PriceStartFrom          *string `form:"priceStartFrom" json:"priceStartFrom" xml:"priceStartFrom" gorm:"type:decimal"`
	PriceUntilFrom          *string `form:"priceUntilFrom" json:"priceUntilFrom" xml:"priceUntilFrom" gorm:"type:decimal"`
	IsDelete                *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TDoctorOfficeTreatmentPriceResponse struct {
	Id                      *uint   `form:"id" json:"id" xml:"id"`
	DoctorOfficeTreatmentId *uint   `form:"doctorOfficeTreatmentId" json:"doctorOfficeTreatmentId" xml:"doctorOfficeTreatmentId"`
	Price                   *string `form:"price" json:"price" xml:"price"`
	PriceStartFrom          *string `form:"priceStartFrom" json:"priceStartFrom" xml:"priceStartFrom"`
	PriceUntilFrom          *string `form:"priceUntilFrom" json:"priceUntilFrom" xml:"priceUntilFrom"`
	IsDelete                bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToTDoctorOfficeTreatmentPriceEntity(req TDoctorOfficeTreatmentPriceRequest, imageData []byte, userId uint) TDoctorOfficeTreatmentPrice {
	entity := TDoctorOfficeTreatmentPrice{
		DoctorOfficeTreatmentId: req.DoctorOfficeTreatmentId,
		Price:                   req.Price,
		PriceStartFrom:          req.PriceStartFrom,
		PriceUntilFrom:          req.PriceUntilFrom,
		CreatedBy:               userId,
		CreatedOn:               dto.JSONTime{Time: time.Now()},
		IsDelete:                false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToTDoctorOfficeTreatmentPriceResponse(entity TDoctorOfficeTreatmentPrice) TDoctorOfficeTreatmentPriceResponse {
	res := TDoctorOfficeTreatmentPriceResponse{
		Id:                      &entity.Id,
		DoctorOfficeTreatmentId: entity.DoctorOfficeTreatmentId,
		Price:                   entity.Price,
		PriceStartFrom:          entity.PriceStartFrom,
		PriceUntilFrom:          entity.PriceUntilFrom,
		IsDelete:                entity.IsDelete,
	}
	return res
}
func ToTDoctorOfficeTreatmentPriceResponsesParallel(entities []TDoctorOfficeTreatmentPrice) []TDoctorOfficeTreatmentPriceResponse {
	numEntities := len(entities)
	responses := make([]TDoctorOfficeTreatmentPriceResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity TDoctorOfficeTreatmentPrice) {
			defer wg.Done()
			responses[index] = ToTDoctorOfficeTreatmentPriceResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
