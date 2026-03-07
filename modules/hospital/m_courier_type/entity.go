package m_courier_type

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MCourierType struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    CourierId  *uint         `form:"courierId" json:"courierId" xml:"courierId" gorm:"type:bigint"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:20;type:varchar(20)" validate:"max=20"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MCourierType) TableName() string {
    return "m_courier_type"
}

type MCourierTypeRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    CourierId  *uint         `form:"courierId" json:"courierId" xml:"courierId" gorm:"type:bigint"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:20;type:varchar(20)" validate:"max=20"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MCourierTypeResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    CourierId  *uint         `form:"courierId" json:"courierId" xml:"courierId"`
    Name       *string       `form:"name" json:"name" xml:"name" validate:"max=20"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMCourierTypeEntity(req MCourierTypeRequest, imageData []byte, userId uint) MCourierType {
    entity := MCourierType{
        CourierId:     req.CourierId,
        Name:          req.Name,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMCourierTypeResponse(entity MCourierType) MCourierTypeResponse {
    res := MCourierTypeResponse{
        Id: &entity.Id,
        CourierId: entity.CourierId,
        Name: entity.Name,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMCourierTypeResponsesParallel(entities []MCourierType) []MCourierTypeResponse {
    numEntities := len(entities)
    responses := make([]MCourierTypeResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MCourierType) {
            defer wg.Done()
            responses[index] = ToMCourierTypeResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}