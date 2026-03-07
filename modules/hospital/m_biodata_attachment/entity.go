package m_biodata_attachment

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MBiodataAttachment struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
    FileName   *string       `form:"fileName" json:"fileName" xml:"fileName" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    FilePath   *string       `form:"filePath" json:"filePath" xml:"filePath" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    FileSize   *int          `form:"fileSize" json:"fileSize" xml:"fileSize" gorm:"type:int"`
    File       *[]byte       `form:"file" json:"file" xml:"file" gorm:"type:blob"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MBiodataAttachment) TableName() string {
    return "m_biodata_attachment"
}

type MBiodataAttachmentRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
    FileName   *string       `form:"fileName" json:"fileName" xml:"fileName" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    FilePath   *string       `form:"filePath" json:"filePath" xml:"filePath" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    FileSize   *int          `form:"fileSize" json:"fileSize" xml:"fileSize" gorm:"type:int"`
    File       *[]byte       `form:"file" json:"file" xml:"file" gorm:"type:blob"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MBiodataAttachmentResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId"`
    FileName   *string       `form:"fileName" json:"fileName" xml:"fileName" validate:"max=50"`
    FilePath   *string       `form:"filePath" json:"filePath" xml:"filePath" validate:"max=100"`
    FileSize   *int          `form:"fileSize" json:"fileSize" xml:"fileSize"`
    File       *[]byte       `form:"file" json:"file" xml:"file"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMBiodataAttachmentEntity(req MBiodataAttachmentRequest, imageData []byte, userId uint) MBiodataAttachment {
    entity := MBiodataAttachment{
        BiodataId:     req.BiodataId,
        FileName:      req.FileName,
        FilePath:      req.FilePath,
        FileSize:      req.FileSize,
        File:          req.File,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMBiodataAttachmentResponse(entity MBiodataAttachment) MBiodataAttachmentResponse {
    res := MBiodataAttachmentResponse{
        Id: &entity.Id,
        BiodataId: entity.BiodataId,
        FileName: entity.FileName,
        FilePath: entity.FilePath,
        FileSize: entity.FileSize,
        File: entity.File,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMBiodataAttachmentResponsesParallel(entities []MBiodataAttachment) []MBiodataAttachmentResponse {
    numEntities := len(entities)
    responses := make([]MBiodataAttachmentResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MBiodataAttachment) {
            defer wg.Done()
            responses[index] = ToMBiodataAttachmentResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}