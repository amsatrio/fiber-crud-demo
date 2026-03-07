package m_user

import (
    "fiber-crud-demo/dto"
    "sync"
    "time"
)

type MUser struct {
    Id         uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
    BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
    RoleId     *uint         `form:"roleId" json:"roleId" xml:"roleId" gorm:"type:bigint"`
    Email      *string       `form:"email" json:"email" xml:"email" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    Password   *string       `form:"password" json:"password" xml:"password" gorm:"size:255;type:varchar(255)" validate:"max=255"`
    LoginAttempt *int          `form:"loginAttempt" json:"loginAttempt" xml:"loginAttempt" gorm:"type:int"`
    IsLocked   *bool         `form:"isLocked" json:"isLocked" xml:"isLocked" gorm:"type:boolean"`
    LastLogin  *dto.JSONTime `form:"lastLogin" json:"lastLogin" xml:"lastLogin" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    CreatedBy  uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
    CreatedOn  dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    ModifiedBy *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
    ModifiedOn *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    DeletedBy  *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
    DeletedOn  *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MUser) TableName() string {
    return "m_user"
}

type MUserRequest struct {
    Id         *uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
    BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId" gorm:"type:bigint"`
    RoleId     *uint         `form:"roleId" json:"roleId" xml:"roleId" gorm:"type:bigint"`
    Email      *string       `form:"email" json:"email" xml:"email" gorm:"size:100;type:varchar(100)" validate:"max=100"`
    Password   *string       `form:"password" json:"password" xml:"password" gorm:"size:255;type:varchar(255)" validate:"max=255"`
    LoginAttempt *int          `form:"loginAttempt" json:"loginAttempt" xml:"loginAttempt" gorm:"type:int"`
    IsLocked   *bool         `form:"isLocked" json:"isLocked" xml:"isLocked" gorm:"type:boolean"`
    LastLogin  *dto.JSONTime `form:"lastLogin" json:"lastLogin" xml:"lastLogin" gorm:"type:datetime"`
    IsDelete   *bool         `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MUserResponse struct {
    Id         *uint         `form:"id" json:"id" xml:"id"`
    BiodataId  *uint         `form:"biodataId" json:"biodataId" xml:"biodataId"`
    RoleId     *uint         `form:"roleId" json:"roleId" xml:"roleId"`
    Email      *string       `form:"email" json:"email" xml:"email" validate:"max=100"`
    Password   *string       `form:"password" json:"password" xml:"password" validate:"max=255"`
    LoginAttempt *int          `form:"loginAttempt" json:"loginAttempt" xml:"loginAttempt"`
    IsLocked   *bool         `form:"isLocked" json:"isLocked" xml:"isLocked"`
    LastLogin  *dto.JSONTime `form:"lastLogin" json:"lastLogin" xml:"lastLogin"`
    IsDelete   bool          `form:"isDelete" json:"isDelete" xml:"isDelete"`
}
func ToMUserEntity(req MUserRequest, imageData []byte, userId uint) MUser {
    entity := MUser{
        BiodataId:     req.BiodataId,
        RoleId:        req.RoleId,
        Email:         req.Email,
        Password:      req.Password,
        LoginAttempt:  req.LoginAttempt,
        IsLocked:      req.IsLocked,
        LastLogin:     req.LastLogin,
        CreatedBy:   userId,
        CreatedOn:   dto.JSONTime{Time: time.Now()},
        IsDelete:    false,
    }
    if req.Id != nil { entity.Id = *req.Id }
    return entity
}

func ToMUserResponse(entity MUser) MUserResponse {
    res := MUserResponse{
        Id: &entity.Id,
        BiodataId: entity.BiodataId,
        RoleId: entity.RoleId,
        Email: entity.Email,
        Password: entity.Password,
        LoginAttempt: entity.LoginAttempt,
        IsLocked: entity.IsLocked,
        LastLogin: entity.LastLogin,
        IsDelete: entity.IsDelete,
    }
    return res
}
func ToMUserResponsesParallel(entities []MUser) []MUserResponse {
    numEntities := len(entities)
    responses := make([]MUserResponse, numEntities)

    var wg sync.WaitGroup
    wg.Add(numEntities)

    for i, e := range entities {
        go func(index int, entity MUser) {
            defer wg.Done()
            responses[index] = ToMUserResponse(entity)
        }(i, e)
    }

    wg.Wait()
    return responses
}