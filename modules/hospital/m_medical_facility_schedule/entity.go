package m_medical_facility_schedule

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MMedicalFacilitySchedule struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    MedicalFacilityId *uint         `form:"medicalFacilityId" json:"medicalFacilityId" xml:"medicalFacilityId" gorm:"type:bigint"`
    Day        *string       `form:"day" json:"day" xml:"day" gorm:"size:10;type:varchar(10);comment:"Senin" validate:"max=10"`
    TimeScheduleStart *string       `form:"timeScheduleStart" json:"timeScheduleStart" xml:"timeScheduleStart" gorm:"size:10;type:varchar(10);comment:Ex : 08:00" validate:"max=10"`
    TimeScheduleEnd *string       `form:"timeScheduleEnd" json:"timeScheduleEnd" xml:"timeScheduleEnd" gorm:"size:10;type:varchar(10);comment:Ex : 19:30" validate:"max=10"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MMedicalFacilitySchedule) TableName() string {
    return "m_medical_facility_schedule"
}

type MMedicalFacilityScheduleRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    MedicalFacilityId *uint         `form:"medicalFacilityId" json:"medicalFacilityId" xml:"medicalFacilityId" gorm:"type:bigint"`
    Day        *string       `form:"day" json:"day" xml:"day" gorm:"size:10;type:varchar(10)" validate:"max=10"`
    TimeScheduleStart *string       `form:"timeScheduleStart" json:"timeScheduleStart" xml:"timeScheduleStart" gorm:"size:10;type:varchar(10)" validate:"max=10"`
    TimeScheduleEnd *string       `form:"timeScheduleEnd" json:"timeScheduleEnd" xml:"timeScheduleEnd" gorm:"size:10;type:varchar(10)" validate:"max=10"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MMedicalFacilityScheduleResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    MedicalFacilityId *uint         `form:"medicalFacilityId" json:"medicalFacilityId" xml:"medicalFacilityId"`
    Day        *string       `form:"day" json:"day" xml:"day" validate:"max=10"`
    TimeScheduleStart *string       `form:"timeScheduleStart" json:"timeScheduleStart" xml:"timeScheduleStart" validate:"max=10"`
    TimeScheduleEnd *string       `form:"timeScheduleEnd" json:"timeScheduleEnd" xml:"timeScheduleEnd" validate:"max=10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMMedicalFacilityScheduleEntity(req MMedicalFacilityScheduleRequest, imageData []byte, userId uint) MMedicalFacilitySchedule {
    entity := MMedicalFacilitySchedule{
        MedicalFacilityId:  req.MedicalFacilityId,
        Day:           req.Day,
        TimeScheduleStart:  req.TimeScheduleStart,
        TimeScheduleEnd:  req.TimeScheduleEnd,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMMedicalFacilityScheduleResponse(entity MMedicalFacilitySchedule) MMedicalFacilityScheduleResponse {
    res := MMedicalFacilityScheduleResponse{
        Id: &entity.Id,
        MedicalFacilityId: entity.MedicalFacilityId,
        Day: entity.Day,
        TimeScheduleStart: entity.TimeScheduleStart,
        TimeScheduleEnd: entity.TimeScheduleEnd,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMMedicalFacilityScheduleResponsesParallel(entities []MMedicalFacilitySchedule) []MMedicalFacilityScheduleResponse {
    numEntities := len(entities)
    responses := make([]MMedicalFacilityScheduleResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MMedicalFacilitySchedule) {
            defer wg.Done()
            responses[index] = ToMMedicalFacilityScheduleResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}