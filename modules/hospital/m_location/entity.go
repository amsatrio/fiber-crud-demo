package m_location

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MLocation struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    ParentId   *uint         `form:"parentId" json:"parentId" xml:"parentId" gorm:"type:bigint"`
    LocationLevelId *uint         `form:"locationLevelId" json:"locationLevelId" xml:"locationLevelId" gorm:"type:bigint"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MLocation) TableName() string {
    return "m_location"
}

type MLocationRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    ParentId   *uint         `form:"parentId" json:"parentId" xml:"parentId" gorm:"type:bigint"`
    LocationLevelId *uint         `form:"locationLevelId" json:"locationLevelId" xml:"locationLevelId" gorm:"type:bigint"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MLocationResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Name       *string       `form:"name" json:"name" xml:"name" validate:"max=100"`
    ParentId   *uint         `form:"parentId" json:"parentId" xml:"parentId"`
    LocationLevelId *uint         `form:"locationLevelId" json:"locationLevelId" xml:"locationLevelId"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMLocationEntity(req MLocationRequest, imageData []byte, userId uint) MLocation {
    entity := MLocation{
        Name:          req.Name,
        ParentId:      req.ParentId,
        LocationLevelId:  req.LocationLevelId,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMLocationResponse(entity MLocation) MLocationResponse {
    res := MLocationResponse{
        Id: &entity.Id,
        Name: entity.Name,
        ParentId: entity.ParentId,
        LocationLevelId: entity.LocationLevelId,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMLocationResponsesParallel(entities []MLocation) []MLocationResponse {
    numEntities := len(entities)
    responses := make([]MLocationResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MLocation) {
            defer wg.Done()
            responses[index] = ToMLocationResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}