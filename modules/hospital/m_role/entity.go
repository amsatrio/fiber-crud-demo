package m_role

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MRole struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:20;type:varchar(20)" validate:"max=20"`
    Code       *string       `form:"code" json:"code" xml:"code" gorm:"size:20;type:varchar(20)" validate:"max=20"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MRole) TableName() string {
    return "m_role"
}

type MRoleRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:20;type:varchar(20)" validate:"max=20"`
    Code       *string       `form:"code" json:"code" xml:"code" gorm:"size:20;type:varchar(20)" validate:"max=20"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MRoleResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Name       *string       `form:"name" json:"name" xml:"name" validate:"max=20"`
    Code       *string       `form:"code" json:"code" xml:"code" validate:"max=20"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMRoleEntity(req MRoleRequest, imageData []byte, userId uint) MRole {
    entity := MRole{
        Name:          req.Name,
        Code:          req.Code,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMRoleResponse(entity MRole) MRoleResponse {
    res := MRoleResponse{
        Id: &entity.Id,
        Name: entity.Name,
        Code: entity.Code,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMRoleResponsesParallel(entities []MRole) []MRoleResponse {
    numEntities := len(entities)
    responses := make([]MRoleResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MRole) {
            defer wg.Done()
            responses[index] = ToMRoleResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}