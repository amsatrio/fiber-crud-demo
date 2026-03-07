package m_blood_group

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MBloodGroup struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Code       *string       `form:"code" json:"code" xml:"code" gorm:"size:5;type:varchar(5)" validate:"max=5"`
    Descrtiption *string       `form:"descrtiption" json:"descrtiption" xml:"descrtiption" gorm:"size:255;type:varchar(255)" validate:"max=255"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MBloodGroup) TableName() string {
    return "m_blood_group"
}

type MBloodGroupRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Code       *string       `form:"code" json:"code" xml:"code" gorm:"size:5;type:varchar(5)" validate:"max=5"`
    Descrtiption *string       `form:"descrtiption" json:"descrtiption" xml:"descrtiption" gorm:"size:255;type:varchar(255)" validate:"max=255"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MBloodGroupResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Code       *string       `form:"code" json:"code" xml:"code" validate:"max=5"`
    Descrtiption *string       `form:"descrtiption" json:"descrtiption" xml:"descrtiption" validate:"max=255"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMBloodGroupEntity(req MBloodGroupRequest, imageData []byte, userId uint) MBloodGroup {
    entity := MBloodGroup{
        Code:          req.Code,
        Descrtiption:  req.Descrtiption,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMBloodGroupResponse(entity MBloodGroup) MBloodGroupResponse {
    res := MBloodGroupResponse{
        Id: &entity.Id,
        Code: entity.Code,
        Descrtiption: entity.Descrtiption,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMBloodGroupResponsesParallel(entities []MBloodGroup) []MBloodGroupResponse {
    numEntities := len(entities)
    responses := make([]MBloodGroupResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MBloodGroup) {
            defer wg.Done()
            responses[index] = ToMBloodGroupResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}