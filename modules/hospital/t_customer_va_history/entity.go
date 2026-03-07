package t_customer_va_history

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TCustomerVaHistory struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    CustomerVaId *uint         `form:"customerVaId" json:"customerVaId" xml:"customerVaId" gorm:"type:bigint"`
    Amount     *string       `form:"amount" json:"amount" xml:"amount" gorm:"type:decimal"`
    ExpiredOn  *dto.JSONTime `form:"expiredOn" json:"expiredOn" xml:"expiredOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCustomerVaHistory) TableName() string {
    return "t_customer_va_history"
}

type TCustomerVaHistoryRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    CustomerVaId *uint         `form:"customerVaId" json:"customerVaId" xml:"customerVaId" gorm:"type:bigint"`
    Amount     *string       `form:"amount" json:"amount" xml:"amount" gorm:"type:decimal"`
    ExpiredOn  *dto.JSONTime `form:"expiredOn" json:"expiredOn" xml:"expiredOn" gorm:"type:datetime"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCustomerVaHistoryResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    CustomerVaId *uint         `form:"customerVaId" json:"customerVaId" xml:"customerVaId"`
    Amount     *string       `form:"amount" json:"amount" xml:"amount"`
    ExpiredOn  *dto.JSONTime `form:"expiredOn" json:"expiredOn" xml:"expiredOn"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTCustomerVaHistoryEntity(req TCustomerVaHistoryRequest, imageData []byte, userId uint) TCustomerVaHistory {
    entity := TCustomerVaHistory{
        CustomerVaId:  req.CustomerVaId,
        Amount:        req.Amount,
        ExpiredOn:     req.ExpiredOn,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTCustomerVaHistoryResponse(entity TCustomerVaHistory) TCustomerVaHistoryResponse {
    res := TCustomerVaHistoryResponse{
        Id: &entity.Id,
        CustomerVaId: entity.CustomerVaId,
        Amount: entity.Amount,
        ExpiredOn: entity.ExpiredOn,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTCustomerVaHistoryResponsesParallel(entities []TCustomerVaHistory) []TCustomerVaHistoryResponse {
    numEntities := len(entities)
    responses := make([]TCustomerVaHistoryResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TCustomerVaHistory) {
            defer wg.Done()
            responses[index] = ToTCustomerVaHistoryResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}