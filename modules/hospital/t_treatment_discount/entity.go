package t_treatment_discount

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TTreatmentDiscount struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    DoctorOfficeTreatmentPriceId *uint         `form:"doctorOfficeTreatmentPriceId" json:"doctorOfficeTreatmentPriceId" xml:"doctorOfficeTreatmentPriceId" gorm:"type:bigint"`
    Value      *string       `form:"value" json:"value" xml:"value" gorm:"type:decimal"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TTreatmentDiscount) TableName() string {
    return "t_treatment_discount"
}

type TTreatmentDiscountRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    DoctorOfficeTreatmentPriceId *uint         `form:"doctorOfficeTreatmentPriceId" json:"doctorOfficeTreatmentPriceId" xml:"doctorOfficeTreatmentPriceId" gorm:"type:bigint"`
    Value      *string       `form:"value" json:"value" xml:"value" gorm:"type:decimal"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TTreatmentDiscountResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    DoctorOfficeTreatmentPriceId *uint         `form:"doctorOfficeTreatmentPriceId" json:"doctorOfficeTreatmentPriceId" xml:"doctorOfficeTreatmentPriceId"`
    Value      *string       `form:"value" json:"value" xml:"value"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTTreatmentDiscountEntity(req TTreatmentDiscountRequest, imageData []byte, userId uint) TTreatmentDiscount {
    entity := TTreatmentDiscount{
        DoctorOfficeTreatmentPriceId:  req.DoctorOfficeTreatmentPriceId,
        Value:         req.Value,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTTreatmentDiscountResponse(entity TTreatmentDiscount) TTreatmentDiscountResponse {
    res := TTreatmentDiscountResponse{
        Id: &entity.Id,
        DoctorOfficeTreatmentPriceId: entity.DoctorOfficeTreatmentPriceId,
        Value: entity.Value,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTTreatmentDiscountResponsesParallel(entities []TTreatmentDiscount) []TTreatmentDiscountResponse {
    numEntities := len(entities)
    responses := make([]TTreatmentDiscountResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TTreatmentDiscount) {
            defer wg.Done()
            responses[index] = ToTTreatmentDiscountResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}