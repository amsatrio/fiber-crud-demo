package t_appointment

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TAppointment struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
    DoctorOfficeId *uint         `form:"doctorOfficeId" json:"doctorOfficeId" xml:"doctorOfficeId" gorm:"type:bigint"`
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

func (TAppointment) TableName() string {
    return "t_appointment"
}

type TAppointmentRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
    DoctorOfficeId *uint         `form:"doctorOfficeId" json:"doctorOfficeId" xml:"doctorOfficeId" gorm:"type:bigint"`
    DoctorOfficeScheduleId *uint         `form:"doctorOfficeScheduleId" json:"doctorOfficeScheduleId" xml:"doctorOfficeScheduleId" gorm:"type:bigint"`
    DoctorOfficeTreatmentId *uint         `form:"doctorOfficeTreatmentId" json:"doctorOfficeTreatmentId" xml:"doctorOfficeTreatmentId" gorm:"type:bigint"`
    AppointmentDate *string       `form:"appointmentDate" json:"appointmentDate" xml:"appointmentDate" gorm:"type:date"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TAppointmentResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId"`
    DoctorOfficeId *uint         `form:"doctorOfficeId" json:"doctorOfficeId" xml:"doctorOfficeId"`
    DoctorOfficeScheduleId *uint         `form:"doctorOfficeScheduleId" json:"doctorOfficeScheduleId" xml:"doctorOfficeScheduleId"`
    DoctorOfficeTreatmentId *uint         `form:"doctorOfficeTreatmentId" json:"doctorOfficeTreatmentId" xml:"doctorOfficeTreatmentId"`
    AppointmentDate *string       `form:"appointmentDate" json:"appointmentDate" xml:"appointmentDate"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTAppointmentEntity(req TAppointmentRequest, imageData []byte, userId uint) TAppointment {
    entity := TAppointment{
        CustomerId:    req.CustomerId,
        DoctorOfficeId:  req.DoctorOfficeId,
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

func ToTAppointmentResponse(entity TAppointment) TAppointmentResponse {
    res := TAppointmentResponse{
        Id: &entity.Id,
        CustomerId: entity.CustomerId,
        DoctorOfficeId: entity.DoctorOfficeId,
        DoctorOfficeScheduleId: entity.DoctorOfficeScheduleId,
        DoctorOfficeTreatmentId: entity.DoctorOfficeTreatmentId,
        AppointmentDate: entity.AppointmentDate,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTAppointmentResponsesParallel(entities []TAppointment) []TAppointmentResponse {
    numEntities := len(entities)
    responses := make([]TAppointmentResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TAppointment) {
            defer wg.Done()
            responses[index] = ToTAppointmentResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}