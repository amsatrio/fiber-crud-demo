package m_medical_item_category

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MMedicalItemCategory struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MMedicalItemCategory) TableName() string {
    return "m_medical_item_category"
}

type MMedicalItemCategoryRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MMedicalItemCategoryResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Name       *string       `form:"name" json:"name" xml:"name" validate:"max=50"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMMedicalItemCategoryEntity(req MMedicalItemCategoryRequest, imageData []byte, userId uint) MMedicalItemCategory {
    entity := MMedicalItemCategory{
        Name:          req.Name,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMMedicalItemCategoryResponse(entity MMedicalItemCategory) MMedicalItemCategoryResponse {
    res := MMedicalItemCategoryResponse{
        Id: &entity.Id,
        Name: entity.Name,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMMedicalItemCategoryResponsesParallel(entities []MMedicalItemCategory) []MMedicalItemCategoryResponse {
    numEntities := len(entities)
    responses := make([]MMedicalItemCategoryResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MMedicalItemCategory) {
            defer wg.Done()
            responses[index] = ToMMedicalItemCategoryResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}