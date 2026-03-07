package t_customer_chat

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TCustomerChat struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCustomerChat) TableName() string {
    return "t_customer_chat"
}

type TCustomerChatRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCustomerChatResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTCustomerChatEntity(req TCustomerChatRequest, imageData []byte, userId uint) TCustomerChat {
    entity := TCustomerChat{
        CustomerId:    req.CustomerId,
        DoctorId:      req.DoctorId,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTCustomerChatResponse(entity TCustomerChat) TCustomerChatResponse {
    res := TCustomerChatResponse{
        Id: &entity.Id,
        CustomerId: entity.CustomerId,
        DoctorId: entity.DoctorId,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTCustomerChatResponsesParallel(entities []TCustomerChat) []TCustomerChatResponse {
    numEntities := len(entities)
    responses := make([]TCustomerChatResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TCustomerChat) {
            defer wg.Done()
            responses[index] = ToTCustomerChatResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}