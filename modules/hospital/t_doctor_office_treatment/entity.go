package t_doctor_office_treatment

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TDoctorOfficeTreatment struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    DoctorTreatmentId *uint         `form:"doctorTreatmentId" json:"doctorTreatmentId" xml:"doctorTreatmentId" gorm:"type:bigint"`
    DoctorOfficeId *uint         `form:"doctorOfficeId" json:"doctorOfficeId" xml:"doctorOfficeId" gorm:"type:bigint"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TDoctorOfficeTreatment) TableName() string {
    return "t_doctor_office_treatment"
}

type TDoctorOfficeTreatmentRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    DoctorTreatmentId *uint         `form:"doctorTreatmentId" json:"doctorTreatmentId" xml:"doctorTreatmentId" gorm:"type:bigint"`
    DoctorOfficeId *uint         `form:"doctorOfficeId" json:"doctorOfficeId" xml:"doctorOfficeId" gorm:"type:bigint"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TDoctorOfficeTreatmentResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    DoctorTreatmentId *uint         `form:"doctorTreatmentId" json:"doctorTreatmentId" xml:"doctorTreatmentId"`
    DoctorOfficeId *uint         `form:"doctorOfficeId" json:"doctorOfficeId" xml:"doctorOfficeId"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTDoctorOfficeTreatmentEntity(req TDoctorOfficeTreatmentRequest, imageData []byte, userId uint) TDoctorOfficeTreatment {
    entity := TDoctorOfficeTreatment{
        DoctorTreatmentId:  req.DoctorTreatmentId,
        DoctorOfficeId:  req.DoctorOfficeId,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTDoctorOfficeTreatmentResponse(entity TDoctorOfficeTreatment) TDoctorOfficeTreatmentResponse {
    res := TDoctorOfficeTreatmentResponse{
        Id: &entity.Id,
        DoctorTreatmentId: entity.DoctorTreatmentId,
        DoctorOfficeId: entity.DoctorOfficeId,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTDoctorOfficeTreatmentResponsesParallel(entities []TDoctorOfficeTreatment) []TDoctorOfficeTreatmentResponse {
    numEntities := len(entities)
    responses := make([]TDoctorOfficeTreatmentResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TDoctorOfficeTreatment) {
            defer wg.Done()
            responses[index] = ToTDoctorOfficeTreatmentResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}