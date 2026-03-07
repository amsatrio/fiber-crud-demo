package t_current_doctor_specialization

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TCurrentDoctorSpecialization struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    SpecializationId *uint         `form:"specializationId" json:"specializationId" xml:"specializationId" gorm:"type:bigint"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCurrentDoctorSpecialization) TableName() string {
    return "t_current_doctor_specialization"
}

type TCurrentDoctorSpecializationRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    SpecializationId *uint         `form:"specializationId" json:"specializationId" xml:"specializationId" gorm:"type:bigint"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCurrentDoctorSpecializationResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId"`
    SpecializationId *uint         `form:"specializationId" json:"specializationId" xml:"specializationId"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTCurrentDoctorSpecializationEntity(req TCurrentDoctorSpecializationRequest, imageData []byte, userId uint) TCurrentDoctorSpecialization {
    entity := TCurrentDoctorSpecialization{
        DoctorId:      req.DoctorId,
        SpecializationId:  req.SpecializationId,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTCurrentDoctorSpecializationResponse(entity TCurrentDoctorSpecialization) TCurrentDoctorSpecializationResponse {
    res := TCurrentDoctorSpecializationResponse{
        Id: &entity.Id,
        DoctorId: entity.DoctorId,
        SpecializationId: entity.SpecializationId,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTCurrentDoctorSpecializationResponsesParallel(entities []TCurrentDoctorSpecialization) []TCurrentDoctorSpecializationResponse {
    numEntities := len(entities)
    responses := make([]TCurrentDoctorSpecializationResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TCurrentDoctorSpecialization) {
            defer wg.Done()
            responses[index] = ToTCurrentDoctorSpecializationResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}