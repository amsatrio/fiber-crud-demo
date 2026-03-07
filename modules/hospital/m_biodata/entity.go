package m_biodata

import (
    "fiber-crud-demo/dto"
    "encoding/base64"
    "mime/multipart"
    "sync"
    "time"
)

type MBiodata struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Fullname   *string       `form:"fullname" json:"fullname" xml:"fullname" gorm:"size:255;type:varchar(255)" validate:"max=255"`
    MobilePhone *string       `form:"mobilePhone" json:"mobilePhone" xml:"mobilePhone" gorm:"size:15;type:varchar(15)" validate:"max=15"`
    Image      *[]byte       `form:"image" json:"image" xml:"image" gorm:"type:blob"`
    ImagePath  *string       `form:"imagePath" json:"imagePath" xml:"imagePath" gorm:"size:255;type:varchar(255)" validate:"max=255"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MBiodata) TableName() string {
    return "m_biodata"
}

type MBiodataRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Fullname   *string       `form:"fullname" json:"fullname" xml:"fullname" gorm:"size:255;type:varchar(255)" validate:"max=255"`
    MobilePhone *string       `form:"mobilePhone" json:"mobilePhone" xml:"mobilePhone" gorm:"size:15;type:varchar(15)" validate:"max=15"`
    Image      *multipart.FileHeader `form:"image" json:"image" xml:"image" gorm:"type:blob"`
    ImagePath  *string       `form:"imagePath" json:"imagePath" xml:"imagePath" gorm:""`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MBiodataResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Fullname   *string       `form:"fullname" json:"fullname" xml:"fullname" validate:"max=255"`
    MobilePhone *string       `form:"mobilePhone" json:"mobilePhone" xml:"mobilePhone" validate:"max=15"`
    Image      *string       `form:"image" json:"image" xml:"image"`
    ImagePath  *string       `form:"imagePath" json:"imagePath" xml:"imagePath" validate:"max=255"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMBiodataEntity(req MBiodataRequest, imageData []byte, userId uint) MBiodata {
    entity := MBiodata{
        Fullname:      req.Fullname,
        MobilePhone:   req.MobilePhone,
        Image:       &imageData,
        ImagePath:     req.ImagePath,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMBiodataResponse(entity MBiodata) MBiodataResponse {
    res := MBiodataResponse{
        Id: &entity.Id,
        Fullname: entity.Fullname,
        MobilePhone: entity.MobilePhone,
        Image: nil,
        ImagePath: entity.ImagePath,
        IsDelete: entity.IsDelete,
    }
    if entity.Image != nil && len(*entity.Image) > 0 {
        encoded := base64.StdEncoding.EncodeToString(*entity.Image)
        res.Image = &encoded
    }
    return res
}
func ToMBiodataResponsesParallel(entities []MBiodata) []MBiodataResponse {
    numEntities := len(entities)
    responses := make([]MBiodataResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MBiodata) {
            defer wg.Done()
            responses[index] = ToMBiodataResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}