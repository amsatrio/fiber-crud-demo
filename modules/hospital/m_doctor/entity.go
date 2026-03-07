package m_doctor

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MDoctor struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
    Str        *string       `form:"str" json:"str" xml:"str" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MDoctor) TableName() string {
    return "m_doctor"
}

type MDoctorRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
    Str        *string       `form:"str" json:"str" xml:"str" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MDoctorResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId"`
    Str        *string       `form:"str" json:"str" xml:"str" validate:"max=50"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMDoctorEntity(req MDoctorRequest, imageData []byte, userId uint) MDoctor {
    entity := MDoctor{
        BiodataId:     req.BiodataId,
        Str:           req.Str,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMDoctorResponse(entity MDoctor) MDoctorResponse {
    res := MDoctorResponse{
        Id: &entity.Id,
        BiodataId: entity.BiodataId,
        Str: entity.Str,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMDoctorResponsesParallel(entities []MDoctor) []MDoctorResponse {
    numEntities := len(entities)
    responses := make([]MDoctorResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MDoctor) {
            defer wg.Done()
            responses[index] = ToMDoctorResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}