package m_menu

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MMenu struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:20;type:varchar(20)" validate:"max=20"`
    Url        *string       `form:"url" json:"url" xml:"url" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    ParentId   *uint         `form:"parentId" json:"parentId" xml:"parentId" gorm:"type:bigint"`
    BigIcon    *string       `form:"bigIcon" json:"bigIcon" xml:"bigIcon" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    SmallIcon  *string       `form:"smallIcon" json:"smallIcon" xml:"smallIcon" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MMenu) TableName() string {
    return "m_menu"
}

type MMenuRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:20;type:varchar(20)" validate:"max=20"`
    Url        *string       `form:"url" json:"url" xml:"url" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    ParentId   *uint         `form:"parentId" json:"parentId" xml:"parentId" gorm:"type:bigint"`
    BigIcon    *string       `form:"bigIcon" json:"bigIcon" xml:"bigIcon" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    SmallIcon  *string       `form:"smallIcon" json:"smallIcon" xml:"smallIcon" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MMenuResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Name       *string       `form:"name" json:"name" xml:"name" validate:"max=20"`
    Url        *string       `form:"url" json:"url" xml:"url" validate:"max=50"`
    ParentId   *uint         `form:"parentId" json:"parentId" xml:"parentId"`
    BigIcon    *string       `form:"bigIcon" json:"bigIcon" xml:"bigIcon" validate:"max=100"`
    SmallIcon  *string       `form:"smallIcon" json:"smallIcon" xml:"smallIcon" validate:"max=100"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMMenuEntity(req MMenuRequest, imageData []byte, userId uint) MMenu {
    entity := MMenu{
        Name:          req.Name,
        Url:           req.Url,
        ParentId:      req.ParentId,
        BigIcon:       req.BigIcon,
        SmallIcon:     req.SmallIcon,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMMenuResponse(entity MMenu) MMenuResponse {
    res := MMenuResponse{
        Id: &entity.Id,
        Name: entity.Name,
        Url: entity.Url,
        ParentId: entity.ParentId,
        BigIcon: entity.BigIcon,
        SmallIcon: entity.SmallIcon,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMMenuResponsesParallel(entities []MMenu) []MMenuResponse {
    numEntities := len(entities)
    responses := make([]MMenuResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MMenu) {
            defer wg.Done()
            responses[index] = ToMMenuResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}