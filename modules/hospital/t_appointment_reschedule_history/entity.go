package t_appointment_reschedule_history

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TAppointmentRescheduleHistory struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    AppointmentId *uint         `form:"appointmentId" json:"appointmentId" xml:"appointmentId" gorm:"type:bigint"`
    DoctorOfficeScheduleId *uint         `form:"doctorOfficeScheduleId" json:"doctorOfficeScheduleId" xml:"doctorOfficeScheduleId" gorm:"type:bigint"`
    DoctorOfficeTreatmentId *uint         `form:"doctorOfficeTreatmentId" json:"doctorOfficeTreatmentId" xml:"doctorOfficeTreatmentId" gorm:"type:bigint"`
    AppointmentDate *string       `form:"appointmentDate" json:"appointmentDate" xml:"appointmentDate" gorm:"type:date"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TAppointmentRescheduleHistory) TableName() string {
    return "t_appointment_reschedule_history"
}

type TAppointmentRescheduleHistoryRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    AppointmentId *uint         `form:"appointmentId" json:"appointmentId" xml:"appointmentId" gorm:"type:bigint"`
    DoctorOfficeScheduleId *uint         `form:"doctorOfficeScheduleId" json:"doctorOfficeScheduleId" xml:"doctorOfficeScheduleId" gorm:"type:bigint"`
    DoctorOfficeTreatmentId *uint         `form:"doctorOfficeTreatmentId" json:"doctorOfficeTreatmentId" xml:"doctorOfficeTreatmentId" gorm:"type:bigint"`
    AppointmentDate *string       `form:"appointmentDate" json:"appointmentDate" xml:"appointmentDate" gorm:"type:date"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TAppointmentRescheduleHistoryResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    AppointmentId *uint         `form:"appointmentId" json:"appointmentId" xml:"appointmentId"`
    DoctorOfficeScheduleId *uint         `form:"doctorOfficeScheduleId" json:"doctorOfficeScheduleId" xml:"doctorOfficeScheduleId"`
    DoctorOfficeTreatmentId *uint         `form:"doctorOfficeTreatmentId" json:"doctorOfficeTreatmentId" xml:"doctorOfficeTreatmentId"`
    AppointmentDate *string       `form:"appointmentDate" json:"appointmentDate" xml:"appointmentDate"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTAppointmentRescheduleHistoryEntity(req TAppointmentRescheduleHistoryRequest, imageData []byte, userId uint) TAppointmentRescheduleHistory {
    entity := TAppointmentRescheduleHistory{
        AppointmentId:  req.AppointmentId,
        DoctorOfficeScheduleId:  req.DoctorOfficeScheduleId,
        DoctorOfficeTreatmentId:  req.DoctorOfficeTreatmentId,
        AppointmentDate:  req.AppointmentDate,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTAppointmentRescheduleHistoryResponse(entity TAppointmentRescheduleHistory) TAppointmentRescheduleHistoryResponse {
    res := TAppointmentRescheduleHistoryResponse{
        Id: &entity.Id,
        AppointmentId: entity.AppointmentId,
        DoctorOfficeScheduleId: entity.DoctorOfficeScheduleId,
        DoctorOfficeTreatmentId: entity.DoctorOfficeTreatmentId,
        AppointmentDate: entity.AppointmentDate,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTAppointmentRescheduleHistoryResponsesParallel(entities []TAppointmentRescheduleHistory) []TAppointmentRescheduleHistoryResponse {
    numEntities := len(entities)
    responses := make([]TAppointmentRescheduleHistoryResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TAppointmentRescheduleHistory) {
            defer wg.Done()
            responses[index] = ToTAppointmentRescheduleHistoryResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}