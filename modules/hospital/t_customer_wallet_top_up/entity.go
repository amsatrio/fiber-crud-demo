package t_customer_wallet_top_up

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TCustomerWalletTopUp struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    CustomerWalletId *uint         `form:"customerWalletId" json:"customerWalletId" xml:"customerWalletId" gorm:"type:bigint"`
    Amount     *string       `form:"amount" json:"amount" xml:"amount" gorm:"type:decimal"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCustomerWalletTopUp) TableName() string {
    return "t_customer_wallet_top_up"
}

type TCustomerWalletTopUpRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    CustomerWalletId *uint         `form:"customerWalletId" json:"customerWalletId" xml:"customerWalletId" gorm:"type:bigint"`
    Amount     *string       `form:"amount" json:"amount" xml:"amount" gorm:"type:decimal"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCustomerWalletTopUpResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    CustomerWalletId *uint         `form:"customerWalletId" json:"customerWalletId" xml:"customerWalletId"`
    Amount     *string       `form:"amount" json:"amount" xml:"amount"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTCustomerWalletTopUpEntity(req TCustomerWalletTopUpRequest, imageData []byte, userId uint) TCustomerWalletTopUp {
    entity := TCustomerWalletTopUp{
        CustomerWalletId:  req.CustomerWalletId,
        Amount:        req.Amount,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTCustomerWalletTopUpResponse(entity TCustomerWalletTopUp) TCustomerWalletTopUpResponse {
    res := TCustomerWalletTopUpResponse{
        Id: &entity.Id,
        CustomerWalletId: entity.CustomerWalletId,
        Amount: entity.Amount,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTCustomerWalletTopUpResponsesParallel(entities []TCustomerWalletTopUp) []TCustomerWalletTopUpResponse {
    numEntities := len(entities)
    responses := make([]TCustomerWalletTopUpResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TCustomerWalletTopUp) {
            defer wg.Done()
            responses[index] = ToTCustomerWalletTopUpResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}