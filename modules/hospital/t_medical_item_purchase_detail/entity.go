package t_medical_item_purchase_detail

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type TMedicalItemPurchaseDetail struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    MedicalItemPurchaseId *uint         `form:"medicalItemPurchaseId" json:"medicalItemPurchaseId" xml:"medicalItemPurchaseId" gorm:"type:bigint"`
    MedicalItemId *uint         `form:"medicalItemId" json:"medicalItemId" xml:"medicalItemId" gorm:"type:bigint"`
    Qty        *int          `form:"qty" json:"qty" xml:"qty" gorm:"type:int"`
    MedicalFacilityId *uint         `form:"medicalFacilityId" json:"medicalFacilityId" xml:"medicalFacilityId" gorm:"type:bigint"`
    CourirId   *uint         `form:"courirId" json:"courirId" xml:"courirId" gorm:"type:bigint"`
    SubTotal   *string       `form:"subTotal" json:"subTotal" xml:"subTotal" gorm:"type:decimal"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TMedicalItemPurchaseDetail) TableName() string {
    return "t_medical_item_purchase_detail"
}

type TMedicalItemPurchaseDetailRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    MedicalItemPurchaseId *uint         `form:"medicalItemPurchaseId" json:"medicalItemPurchaseId" xml:"medicalItemPurchaseId" gorm:"type:bigint"`
    MedicalItemId *uint         `form:"medicalItemId" json:"medicalItemId" xml:"medicalItemId" gorm:"type:bigint"`
    Qty        *int          `form:"qty" json:"qty" xml:"qty" gorm:"type:int"`
    MedicalFacilityId *uint         `form:"medicalFacilityId" json:"medicalFacilityId" xml:"medicalFacilityId" gorm:"type:bigint"`
    CourirId   *uint         `form:"courirId" json:"courirId" xml:"courirId" gorm:"type:bigint"`
    SubTotal   *string       `form:"subTotal" json:"subTotal" xml:"subTotal" gorm:"type:decimal"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TMedicalItemPurchaseDetailResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    MedicalItemPurchaseId *uint         `form:"medicalItemPurchaseId" json:"medicalItemPurchaseId" xml:"medicalItemPurchaseId"`
    MedicalItemId *uint         `form:"medicalItemId" json:"medicalItemId" xml:"medicalItemId"`
    Qty        *int          `form:"qty" json:"qty" xml:"qty"`
    MedicalFacilityId *uint         `form:"medicalFacilityId" json:"medicalFacilityId" xml:"medicalFacilityId"`
    CourirId   *uint         `form:"courirId" json:"courirId" xml:"courirId"`
    SubTotal   *string       `form:"subTotal" json:"subTotal" xml:"subTotal"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToTMedicalItemPurchaseDetailEntity(req TMedicalItemPurchaseDetailRequest, imageData []byte, userId uint) TMedicalItemPurchaseDetail {
    entity := TMedicalItemPurchaseDetail{
        MedicalItemPurchaseId:  req.MedicalItemPurchaseId,
        MedicalItemId:  req.MedicalItemId,
        Qty:           req.Qty,
        MedicalFacilityId:  req.MedicalFacilityId,
        CourirId:      req.CourirId,
        SubTotal:      req.SubTotal,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToTMedicalItemPurchaseDetailResponse(entity TMedicalItemPurchaseDetail) TMedicalItemPurchaseDetailResponse {
    res := TMedicalItemPurchaseDetailResponse{
        Id: &entity.Id,
        MedicalItemPurchaseId: entity.MedicalItemPurchaseId,
        MedicalItemId: entity.MedicalItemId,
        Qty: entity.Qty,
        MedicalFacilityId: entity.MedicalFacilityId,
        CourirId: entity.CourirId,
        SubTotal: entity.SubTotal,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToTMedicalItemPurchaseDetailResponsesParallel(entities []TMedicalItemPurchaseDetail) []TMedicalItemPurchaseDetailResponse {
    numEntities := len(entities)
    responses := make([]TMedicalItemPurchaseDetailResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity TMedicalItemPurchaseDetail) {
            defer wg.Done()
            responses[index] = ToTMedicalItemPurchaseDetailResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}