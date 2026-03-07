package t_doctor_treatment

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TDoctorTreatment struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TDoctorTreatment) TableName() string {
    return "t_doctor_treatment"
}

type TDoctorTreatmentRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TDoctorTreatmentResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId"`
    Name       *string       `form:"name" json:"name" xml:"name" validate:"max=50"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTDoctorTreatmentEntity(req TDoctorTreatmentRequest, imageData []byte, userId uint) TDoctorTreatment {
    entity := TDoctorTreatment{
        DoctorId:      req.DoctorId,
        Name:          req.Name,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTDoctorTreatmentResponse(entity TDoctorTreatment) TDoctorTreatmentResponse {
    res := TDoctorTreatmentResponse{
        Id: &entity.Id,
        DoctorId: entity.DoctorId,
        Name: entity.Name,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTDoctorTreatmentResponsesParallel(entities []TDoctorTreatment) []TDoctorTreatmentResponse {
    numEntities := len(entities)
    responses := make([]TDoctorTreatmentResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TDoctorTreatment) {
            defer wg.Done()
            responses[index] = ToTDoctorTreatmentResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}