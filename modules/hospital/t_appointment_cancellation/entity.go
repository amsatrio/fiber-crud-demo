package t_appointment_cancellation

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TAppointmentCancellation struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    AppointmentId *uint         `form:"appointmentId" json:"appointmentId" xml:"appointmentId" gorm:"type:bigint"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TAppointmentCancellation) TableName() string {
    return "t_appointment_cancellation"
}

type TAppointmentCancellationRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    AppointmentId *uint         `form:"appointmentId" json:"appointmentId" xml:"appointmentId" gorm:"type:bigint"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TAppointmentCancellationResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    AppointmentId *uint         `form:"appointmentId" json:"appointmentId" xml:"appointmentId"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTAppointmentCancellationEntity(req TAppointmentCancellationRequest, imageData []byte, userId uint) TAppointmentCancellation {
    entity := TAppointmentCancellation{
        AppointmentId:  req.AppointmentId,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTAppointmentCancellationResponse(entity TAppointmentCancellation) TAppointmentCancellationResponse {
    res := TAppointmentCancellationResponse{
        Id: &entity.Id,
        AppointmentId: entity.AppointmentId,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTAppointmentCancellationResponsesParallel(entities []TAppointmentCancellation) []TAppointmentCancellationResponse {
    numEntities := len(entities)
    responses := make([]TAppointmentCancellationResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TAppointmentCancellation) {
            defer wg.Done()
            responses[index] = ToTAppointmentCancellationResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}