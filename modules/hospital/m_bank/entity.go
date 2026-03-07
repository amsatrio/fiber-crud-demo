package m_bank

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MBank struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    VaCode     *string       `form:"vaCode" json:"vaCode" xml:"vaCode" gorm:"size:10;type:varchar(10)" validate:"max=10"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MBank) TableName() string {
    return "m_bank"
}

type MBankRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    VaCode     *string       `form:"vaCode" json:"vaCode" xml:"vaCode" gorm:"size:10;type:varchar(10)" validate:"max=10"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MBankResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Name       *string       `form:"name" json:"name" xml:"name" validate:"max=100"`
    VaCode     *string       `form:"vaCode" json:"vaCode" xml:"vaCode" validate:"max=10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMBankEntity(req MBankRequest, imageData []byte, userId uint) MBank {
    entity := MBank{
        Name:          req.Name,
        VaCode:        req.VaCode,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMBankResponse(entity MBank) MBankResponse {
    res := MBankResponse{
        Id: &entity.Id,
        Name: entity.Name,
        VaCode: entity.VaCode,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMBankResponsesParallel(entities []MBank) []MBankResponse {
    numEntities := len(entities)
    responses := make([]MBankResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MBank) {
            defer wg.Done()
            responses[index] = ToMBankResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}