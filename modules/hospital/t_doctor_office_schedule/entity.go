package t_doctor_office_schedule

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TDoctorOfficeSchedule struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    MedicalFacilityScheduleId *uint         `form:"medicalFacilityScheduleId" json:"medicalFacilityScheduleId" xml:"medicalFacilityScheduleId" gorm:"type:bigint"`
    Slot       *int          `form:"slot" json:"slot" xml:"slot" gorm:"type:int"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TDoctorOfficeSchedule) TableName() string {
    return "t_doctor_office_schedule"
}

type TDoctorOfficeScheduleRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    MedicalFacilityScheduleId *uint         `form:"medicalFacilityScheduleId" json:"medicalFacilityScheduleId" xml:"medicalFacilityScheduleId" gorm:"type:bigint"`
    Slot       *int          `form:"slot" json:"slot" xml:"slot" gorm:"type:int"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TDoctorOfficeScheduleResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId"`
    MedicalFacilityScheduleId *uint         `form:"medicalFacilityScheduleId" json:"medicalFacilityScheduleId" xml:"medicalFacilityScheduleId"`
    Slot       *int          `form:"slot" json:"slot" xml:"slot"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTDoctorOfficeScheduleEntity(req TDoctorOfficeScheduleRequest, imageData []byte, userId uint) TDoctorOfficeSchedule {
    entity := TDoctorOfficeSchedule{
        DoctorId:      req.DoctorId,
        MedicalFacilityScheduleId:  req.MedicalFacilityScheduleId,
        Slot:          req.Slot,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTDoctorOfficeScheduleResponse(entity TDoctorOfficeSchedule) TDoctorOfficeScheduleResponse {
    res := TDoctorOfficeScheduleResponse{
        Id: &entity.Id,
        DoctorId: entity.DoctorId,
        MedicalFacilityScheduleId: entity.MedicalFacilityScheduleId,
        Slot: entity.Slot,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTDoctorOfficeScheduleResponsesParallel(entities []TDoctorOfficeSchedule) []TDoctorOfficeScheduleResponse {
    numEntities := len(entities)
    responses := make([]TDoctorOfficeScheduleResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TDoctorOfficeSchedule) {
            defer wg.Done()
            responses[index] = ToTDoctorOfficeScheduleResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}