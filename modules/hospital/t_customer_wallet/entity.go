package t_customer_wallet

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TCustomerWallet struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
    Pin        *string       `form:"pin" json:"pin" xml:"pin" gorm:"size:6;type:varchar(6)" validate:"max=6"`
    Balance    *string       `form:"balance" json:"balance" xml:"balance" gorm:"type:decimal"`
    Barcode    *string       `form:"barcode" json:"barcode" xml:"barcode" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    Points     *string       `form:"points" json:"points" xml:"points" gorm:"type:decimal"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCustomerWallet) TableName() string {
    return "t_customer_wallet"
}

type TCustomerWalletRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
    Pin        *string       `form:"pin" json:"pin" xml:"pin" gorm:"size:6;type:varchar(6)" validate:"max=6"`
    Balance    *string       `form:"balance" json:"balance" xml:"balance" gorm:"type:decimal"`
    Barcode    *string       `form:"barcode" json:"barcode" xml:"barcode" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    Points     *string       `form:"points" json:"points" xml:"points" gorm:"type:decimal"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCustomerWalletResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId"`
    Pin        *string       `form:"pin" json:"pin" xml:"pin" validate:"max=6"`
    Balance    *string       `form:"balance" json:"balance" xml:"balance"`
    Barcode    *string       `form:"barcode" json:"barcode" xml:"barcode" validate:"max=50"`
    Points     *string       `form:"points" json:"points" xml:"points"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTCustomerWalletEntity(req TCustomerWalletRequest, imageData []byte, userId uint) TCustomerWallet {
    entity := TCustomerWallet{
        CustomerId:    req.CustomerId,
        Pin:           req.Pin,
        Balance:       req.Balance,
        Barcode:       req.Barcode,
        Points:        req.Points,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTCustomerWalletResponse(entity TCustomerWallet) TCustomerWalletResponse {
    res := TCustomerWalletResponse{
        Id: &entity.Id,
        CustomerId: entity.CustomerId,
        Pin: entity.Pin,
        Balance: entity.Balance,
        Barcode: entity.Barcode,
        Points: entity.Points,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTCustomerWalletResponsesParallel(entities []TCustomerWallet) []TCustomerWalletResponse {
    numEntities := len(entities)
    responses := make([]TCustomerWalletResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TCustomerWallet) {
            defer wg.Done()
            responses[index] = ToTCustomerWalletResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}