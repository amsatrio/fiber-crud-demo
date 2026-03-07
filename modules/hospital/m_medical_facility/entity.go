package m_medical_facility

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MMedicalFacility struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    MedicalFacilityCategoryId *uint         `form:"medicalFacilityCategoryId" json:"medicalFacilityCategoryId" xml:"medicalFacilityCategoryId" gorm:"type:bigint"`
    LocationId *uint         `form:"locationId" json:"locationId" xml:"locationId" gorm:"type:bigint"`
    FullAddress *string       `form:"fullAddress" json:"fullAddress" xml:"fullAddress" gorm:"type:text"`
    Email      *string       `form:"email" json:"email" xml:"email" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    PhoneCode  *string       `form:"phoneCode" json:"phoneCode" xml:"phoneCode" gorm:"size:10;type:varchar(10)" validate:"max=10"`
    Phone      *string       `form:"phone" json:"phone" xml:"phone" gorm:"size:15;type:varchar(15)" validate:"max=15"`
    Fax        *string       `form:"fax" json:"fax" xml:"fax" gorm:"size:15;type:varchar(15)" validate:"max=15"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MMedicalFacility) TableName() string {
    return "m_medical_facility"
}

type MMedicalFacilityRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    MedicalFacilityCategoryId *uint         `form:"medicalFacilityCategoryId" json:"medicalFacilityCategoryId" xml:"medicalFacilityCategoryId" gorm:"type:bigint"`
    LocationId *uint         `form:"locationId" json:"locationId" xml:"locationId" gorm:"type:bigint"`
    FullAddress *string       `form:"fullAddress" json:"fullAddress" xml:"fullAddress" gorm:"type:text"`
    Email      *string       `form:"email" json:"email" xml:"email" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    PhoneCode  *string       `form:"phoneCode" json:"phoneCode" xml:"phoneCode" gorm:"size:10;type:varchar(10)" validate:"max=10"`
    Phone      *string       `form:"phone" json:"phone" xml:"phone" gorm:"size:15;type:varchar(15)" validate:"max=15"`
    Fax        *string       `form:"fax" json:"fax" xml:"fax" gorm:"size:15;type:varchar(15)" validate:"max=15"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MMedicalFacilityResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Name       *string       `form:"name" json:"name" xml:"name" validate:"max=50"`
    MedicalFacilityCategoryId *uint         `form:"medicalFacilityCategoryId" json:"medicalFacilityCategoryId" xml:"medicalFacilityCategoryId"`
    LocationId *uint         `form:"locationId" json:"locationId" xml:"locationId"`
    FullAddress *string       `form:"fullAddress" json:"fullAddress" xml:"fullAddress"`
    Email      *string       `form:"email" json:"email" xml:"email" validate:"max=100"`
    PhoneCode  *string       `form:"phoneCode" json:"phoneCode" xml:"phoneCode" validate:"max=10"`
    Phone      *string       `form:"phone" json:"phone" xml:"phone" validate:"max=15"`
    Fax        *string       `form:"fax" json:"fax" xml:"fax" validate:"max=15"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMMedicalFacilityEntity(req MMedicalFacilityRequest, imageData []byte, userId uint) MMedicalFacility {
    entity := MMedicalFacility{
        Name:          req.Name,
        MedicalFacilityCategoryId:  req.MedicalFacilityCategoryId,
        LocationId:    req.LocationId,
        FullAddress:   req.FullAddress,
        Email:         req.Email,
        PhoneCode:     req.PhoneCode,
        Phone:         req.Phone,
        Fax:           req.Fax,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMMedicalFacilityResponse(entity MMedicalFacility) MMedicalFacilityResponse {
    res := MMedicalFacilityResponse{
        Id: &entity.Id,
        Name: entity.Name,
        MedicalFacilityCategoryId: entity.MedicalFacilityCategoryId,
        LocationId: entity.LocationId,
        FullAddress: entity.FullAddress,
        Email: entity.Email,
        PhoneCode: entity.PhoneCode,
        Phone: entity.Phone,
        Fax: entity.Fax,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMMedicalFacilityResponsesParallel(entities []MMedicalFacility) []MMedicalFacilityResponse {
    numEntities := len(entities)
    responses := make([]MMedicalFacilityResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MMedicalFacility) {
            defer wg.Done()
            responses[index] = ToMMedicalFacilityResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}