package m_doctor_education

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MDoctorEducation struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    EducationLevelId *uint         `form:"educationLevelId" json:"educationLevelId" xml:"educationLevelId" gorm:"type:bigint"`
    InstitutionName *string       `form:"institutionName" json:"institutionName" xml:"institutionName" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    Major      *string       `form:"major" json:"major" xml:"major" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    StartYear  *string       `form:"startYear" json:"startYear" xml:"startYear" gorm:"size:4;type:varchar(4)" validate:"max=4"`
    EndYear    *string       `form:"endYear" json:"endYear" xml:"endYear" gorm:"size:4;type:varchar(4)" validate:"max=4"`
    IsLastEducation *bool         `form:"isLastEducation" json:"isLastEducation" xml:"isLastEducation" gorm:"type:boolean;comment:default FALSE"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MDoctorEducation) TableName() string {
    return "m_doctor_education"
}

type MDoctorEducationRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId" gorm:"type:bigint"`
    EducationLevelId *uint         `form:"educationLevelId" json:"educationLevelId" xml:"educationLevelId" gorm:"type:bigint"`
    InstitutionName *string       `form:"institutionName" json:"institutionName" xml:"institutionName" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    Major      *string       `form:"major" json:"major" xml:"major" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    StartYear  *string       `form:"startYear" json:"startYear" xml:"startYear" gorm:"size:4;type:varchar(4)" validate:"max=4"`
    EndYear    *string       `form:"endYear" json:"endYear" xml:"endYear" gorm:"size:4;type:varchar(4)" validate:"max=4"`
    IsLastEducation *bool         `form:"isLastEducation" json:"isLastEducation" xml:"isLastEducation" gorm:"type:boolean"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MDoctorEducationResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    DoctorId   *uint         `form:"doctorId" json:"doctorId" xml:"doctorId"`
    EducationLevelId *uint         `form:"educationLevelId" json:"educationLevelId" xml:"educationLevelId"`
    InstitutionName *string       `form:"institutionName" json:"institutionName" xml:"institutionName" validate:"max=100"`
    Major      *string       `form:"major" json:"major" xml:"major" validate:"max=100"`
    StartYear  *string       `form:"startYear" json:"startYear" xml:"startYear" validate:"max=4"`
    EndYear    *string       `form:"endYear" json:"endYear" xml:"endYear" validate:"max=4"`
    IsLastEducation *bool         `form:"isLastEducation" json:"isLastEducation" xml:"isLastEducation"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMDoctorEducationEntity(req MDoctorEducationRequest, imageData []byte, userId uint) MDoctorEducation {
    entity := MDoctorEducation{
        DoctorId:      req.DoctorId,
        EducationLevelId:  req.EducationLevelId,
        InstitutionName:  req.InstitutionName,
        Major:         req.Major,
        StartYear:     req.StartYear,
        EndYear:       req.EndYear,
        IsLastEducation:  req.IsLastEducation,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMDoctorEducationResponse(entity MDoctorEducation) MDoctorEducationResponse {
    res := MDoctorEducationResponse{
        Id: &entity.Id,
        DoctorId: entity.DoctorId,
        EducationLevelId: entity.EducationLevelId,
        InstitutionName: entity.InstitutionName,
        Major: entity.Major,
        StartYear: entity.StartYear,
        EndYear: entity.EndYear,
        IsLastEducation: entity.IsLastEducation,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMDoctorEducationResponsesParallel(entities []MDoctorEducation) []MDoctorEducationResponse {
    numEntities := len(entities)
    responses := make([]MDoctorEducationResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MDoctorEducation) {
            defer wg.Done()
            responses[index] = ToMDoctorEducationResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}