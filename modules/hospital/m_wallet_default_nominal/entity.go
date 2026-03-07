package m_wallet_default_nominal

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MWalletDefaultNominal struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Nominal    *int          `form:"nominal" json:"nominal" xml:"nominal" gorm:"type:int"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MWalletDefaultNominal) TableName() string {
    return "m_wallet_default_nominal"
}

type MWalletDefaultNominalRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Nominal    *int          `form:"nominal" json:"nominal" xml:"nominal" gorm:"type:int"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MWalletDefaultNominalResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Nominal    *int          `form:"nominal" json:"nominal" xml:"nominal"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMWalletDefaultNominalEntity(req MWalletDefaultNominalRequest, imageData []byte, userId uint) MWalletDefaultNominal {
    entity := MWalletDefaultNominal{
        Nominal:       req.Nominal,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMWalletDefaultNominalResponse(entity MWalletDefaultNominal) MWalletDefaultNominalResponse {
    res := MWalletDefaultNominalResponse{
        Id: &entity.Id,
        Nominal: entity.Nominal,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMWalletDefaultNominalResponsesParallel(entities []MWalletDefaultNominal) []MWalletDefaultNominalResponse {
    numEntities := len(entities)
    responses := make([]MWalletDefaultNominalResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MWalletDefaultNominal) {
            defer wg.Done()
            responses[index] = ToMWalletDefaultNominalResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}