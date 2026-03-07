package t_medical_item_purchase

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TMedicalItemPurchase struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
    PaymentMethodId *uint         `form:"paymentMethodId" json:"paymentMethodId" xml:"paymentMethodId" gorm:"type:bigint"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TMedicalItemPurchase) TableName() string {
    return "t_medical_item_purchase"
}

type TMedicalItemPurchaseRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
    PaymentMethodId *uint         `form:"paymentMethodId" json:"paymentMethodId" xml:"paymentMethodId" gorm:"type:bigint"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TMedicalItemPurchaseResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    CustomerId *uint         `form:"customerId" json:"customerId" xml:"customerId"`
    PaymentMethodId *uint         `form:"paymentMethodId" json:"paymentMethodId" xml:"paymentMethodId"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTMedicalItemPurchaseEntity(req TMedicalItemPurchaseRequest, imageData []byte, userId uint) TMedicalItemPurchase {
    entity := TMedicalItemPurchase{
        CustomerId:    req.CustomerId,
        PaymentMethodId:  req.PaymentMethodId,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTMedicalItemPurchaseResponse(entity TMedicalItemPurchase) TMedicalItemPurchaseResponse {
    res := TMedicalItemPurchaseResponse{
        Id: &entity.Id,
        CustomerId: entity.CustomerId,
        PaymentMethodId: entity.PaymentMethodId,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTMedicalItemPurchaseResponsesParallel(entities []TMedicalItemPurchase) []TMedicalItemPurchaseResponse {
    numEntities := len(entities)
    responses := make([]TMedicalItemPurchaseResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TMedicalItemPurchase) {
            defer wg.Done()
            responses[index] = ToTMedicalItemPurchaseResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}