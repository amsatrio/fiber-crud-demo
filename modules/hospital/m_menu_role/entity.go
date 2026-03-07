package m_menu_role

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MMenuRole struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    MenuId     *uint         `form:"menuId" json:"menuId" xml:"menuId" gorm:"type:bigint"`
    RoleId     *uint         `form:"roleId" json:"roleId" xml:"roleId" gorm:"type:bigint"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MMenuRole) TableName() string {
    return "m_menu_role"
}

type MMenuRoleRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    MenuId     *uint         `form:"menuId" json:"menuId" xml:"menuId" gorm:"type:bigint"`
    RoleId     *uint         `form:"roleId" json:"roleId" xml:"roleId" gorm:"type:bigint"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MMenuRoleResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    MenuId     *uint         `form:"menuId" json:"menuId" xml:"menuId"`
    RoleId     *uint         `form:"roleId" json:"roleId" xml:"roleId"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMMenuRoleEntity(req MMenuRoleRequest, imageData []byte, userId uint) MMenuRole {
    entity := MMenuRole{
        MenuId:        req.MenuId,
        RoleId:        req.RoleId,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMMenuRoleResponse(entity MMenuRole) MMenuRoleResponse {
    res := MMenuRoleResponse{
        Id: &entity.Id,
        MenuId: entity.MenuId,
        RoleId: entity.RoleId,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMMenuRoleResponsesParallel(entities []MMenuRole) []MMenuRoleResponse {
    numEntities := len(entities)
    responses := make([]MMenuRoleResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MMenuRole) {
            defer wg.Done()
            responses[index] = ToMMenuRoleResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}