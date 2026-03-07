package m_medical_item

import (
    "fiber-crud-demo/dto"
    "encoding/base64"
    "mime/multipart"
    "sync"
    "time"
)

type MMedicalItem struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    MedicalItemCategoryId *uint         `form:"medicalItemCategoryId" json:"medicalItemCategoryId" xml:"medicalItemCategoryId" gorm:"type:bigint"`
    Composition *string       `form:"composition" json:"composition" xml:"composition" gorm:"type:text"`
    MedicalItemSegmentationId *uint         `form:"medicalItemSegmentationId" json:"medicalItemSegmentationId" xml:"medicalItemSegmentationId" gorm:"type:bigint"`
    Manufacturer *string       `form:"manufacturer" json:"manufacturer" xml:"manufacturer" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    Indication *string       `form:"indication" json:"indication" xml:"indication" gorm:"type:text"`
    Dosage     *string       `form:"dosage" json:"dosage" xml:"dosage" gorm:"type:text"`
    Directions *string       `form:"directions" json:"directions" xml:"directions" gorm:"type:text;comment:aturan pakai"`
    Contraindication *string       `form:"contraindication" json:"contraindication" xml:"contraindication" gorm:"type:text"`
    Caution    *string       `form:"caution" json:"caution" xml:"caution" gorm:"type:text"`
    Packaging  *string       `form:"packaging" json:"packaging" xml:"packaging" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    PriceMax   *uint         `form:"priceMax" json:"priceMax" xml:"priceMax" gorm:"type:bigint"`
    PriceMin   *uint         `form:"priceMin" json:"priceMin" xml:"priceMin" gorm:"type:bigint"`
    Image      *[]byte       `form:"image" json:"image" xml:"image" gorm:"type:blob"`
    ImagePath  *string       `form:"imagePath" json:"imagePath" xml:"imagePath" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MMedicalItem) TableName() string {
    return "m_medical_item"
}

type MMedicalItemRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    Name       *string       `form:"name" json:"name" xml:"name" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    MedicalItemCategoryId *uint         `form:"medicalItemCategoryId" json:"medicalItemCategoryId" xml:"medicalItemCategoryId" gorm:"type:bigint"`
    Composition *string       `form:"composition" json:"composition" xml:"composition" gorm:"type:text"`
    MedicalItemSegmentationId *uint         `form:"medicalItemSegmentationId" json:"medicalItemSegmentationId" xml:"medicalItemSegmentationId" gorm:"type:bigint"`
    Manufacturer *string       `form:"manufacturer" json:"manufacturer" xml:"manufacturer" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    Indication *string       `form:"indication" json:"indication" xml:"indication" gorm:"type:text"`
    Dosage     *string       `form:"dosage" json:"dosage" xml:"dosage" gorm:"type:text"`
    Directions *string       `form:"directions" json:"directions" xml:"directions" gorm:"type:text"`
    Contraindication *string       `form:"contraindication" json:"contraindication" xml:"contraindication" gorm:"type:text"`
    Caution    *string       `form:"caution" json:"caution" xml:"caution" gorm:"type:text"`
    Packaging  *string       `form:"packaging" json:"packaging" xml:"packaging" gorm:"size:50;type:varchar(50)" validate:"max=50"`
    PriceMax   *uint         `form:"priceMax" json:"priceMax" xml:"priceMax" gorm:"type:bigint"`
    PriceMin   *uint         `form:"priceMin" json:"priceMin" xml:"priceMin" gorm:"type:bigint"`
    Image      *multipart.FileHeader `form:"image" json:"image" xml:"image" gorm:"type:blob"`
    ImagePath  *string       `form:"imagePath" json:"imagePath" xml:"imagePath" gorm:""`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MMedicalItemResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    Name       *string       `form:"name" json:"name" xml:"name" validate:"max=50"`
    MedicalItemCategoryId *uint         `form:"medicalItemCategoryId" json:"medicalItemCategoryId" xml:"medicalItemCategoryId"`
    Composition *string       `form:"composition" json:"composition" xml:"composition"`
    MedicalItemSegmentationId *uint         `form:"medicalItemSegmentationId" json:"medicalItemSegmentationId" xml:"medicalItemSegmentationId"`
    Manufacturer *string       `form:"manufacturer" json:"manufacturer" xml:"manufacturer" validate:"max=100"`
    Indication *string       `form:"indication" json:"indication" xml:"indication"`
    Dosage     *string       `form:"dosage" json:"dosage" xml:"dosage"`
    Directions *string       `form:"directions" json:"directions" xml:"directions"`
    Contraindication *string       `form:"contraindication" json:"contraindication" xml:"contraindication"`
    Caution    *string       `form:"caution" json:"caution" xml:"caution"`
    Packaging  *string       `form:"packaging" json:"packaging" xml:"packaging" validate:"max=50"`
    PriceMax   *uint         `form:"priceMax" json:"priceMax" xml:"priceMax"`
    PriceMin   *uint         `form:"priceMin" json:"priceMin" xml:"priceMin"`
    Image      *string       `form:"image" json:"image" xml:"image"`
    ImagePath  *string       `form:"imagePath" json:"imagePath" xml:"imagePath" validate:"max=100"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMMedicalItemEntity(req MMedicalItemRequest, imageData []byte, userId uint) MMedicalItem {
    entity := MMedicalItem{
        Name:          req.Name,
        MedicalItemCategoryId:  req.MedicalItemCategoryId,
        Composition:   req.Composition,
        MedicalItemSegmentationId:  req.MedicalItemSegmentationId,
        Manufacturer:  req.Manufacturer,
        Indication:    req.Indication,
        Dosage:        req.Dosage,
        Directions:    req.Directions,
        Contraindication:  req.Contraindication,
        Caution:       req.Caution,
        Packaging:     req.Packaging,
        PriceMax:      req.PriceMax,
        PriceMin:      req.PriceMin,
        Image:       &imageData,
        ImagePath:     req.ImagePath,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMMedicalItemResponse(entity MMedicalItem) MMedicalItemResponse {
    res := MMedicalItemResponse{
        Id: &entity.Id,
        Name: entity.Name,
        MedicalItemCategoryId: entity.MedicalItemCategoryId,
        Composition: entity.Composition,
        MedicalItemSegmentationId: entity.MedicalItemSegmentationId,
        Manufacturer: entity.Manufacturer,
        Indication: entity.Indication,
        Dosage: entity.Dosage,
        Directions: entity.Directions,
        Contraindication: entity.Contraindication,
        Caution: entity.Caution,
        Packaging: entity.Packaging,
        PriceMax: entity.PriceMax,
        PriceMin: entity.PriceMin,
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
func ToMMedicalItemResponsesParallel(entities []MMedicalItem) []MMedicalItemResponse {
    numEntities := len(entities)
    responses := make([]MMedicalItemResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MMedicalItem) {
            defer wg.Done()
            responses[index] = ToMMedicalItemResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}