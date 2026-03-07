package t_doctor_office

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TDoctorOffice struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    MedicalFacilityId *uint         `form:"medicalFacilityId" json:"medicalFacilityId" xml:"medicalFacilityId" gorm:"type:bigint"`
    Specialization *string       `form:"specialization" json:"specialization" xml:"specialization" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TDoctorOffice) TableName() string {
    return "t_doctor_office"
}

type TDoctorOfficeRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    MedicalFacilityId *uint         `form:"medicalFacilityId" json:"medicalFacilityId" xml:"medicalFacilityId" gorm:"type:bigint"`
    Specialization *string       `form:"specialization" json:"specialization" xml:"specialization" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TDoctorOfficeResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId"`
    MedicalFacilityId *uint         `form:"medicalFacilityId" json:"medicalFacilityId" xml:"medicalFacilityId"`
    Specialization *string       `form:"specialization" json:"specialization" xml:"specialization" validate:"max=100"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTDoctorOfficeEntity(req TDoctorOfficeRequest, imageData []byte, userId uint) TDoctorOffice {
    entity := TDoctorOffice{
        DoctorId:      req.DoctorId,
        MedicalFacilityId:  req.MedicalFacilityId,
        Specialization:  req.Specialization,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTDoctorOfficeResponse(entity TDoctorOffice) TDoctorOfficeResponse {
    res := TDoctorOfficeResponse{
        Id: &entity.Id,
        DoctorId: entity.DoctorId,
        MedicalFacilityId: entity.MedicalFacilityId,
        Specialization: entity.Specialization,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTDoctorOfficeResponsesParallel(entities []TDoctorOffice) []TDoctorOfficeResponse {
    numEntities := len(entities)
    responses := make([]TDoctorOfficeResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TDoctorOffice) {
            defer wg.Done()
            responses[index] = ToTDoctorOfficeResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}