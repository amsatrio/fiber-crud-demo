package schema

import (
	"fiber-crud-demo/dto"
	"time"
)

type MRole struct {
	Id         uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment" binding:"required"`
	Name       string       `form:"name" json:"name" xml:"name" gorm:"size:20;type:varchar(20)" binding:"max=20"`
	Code       string       `form:"code" json:"code" xml:"code" gorm:"size:20;type:varchar(20)" binding:"max=20"`
	Level      int          `form:"level" json:"level" xml:"level" gorm:"type:tinyint"`
	CreatedBy  uint         `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn  dto.JSONTime `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy  uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn  dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete   *bool        `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}

func (MRole) TableName() string {
	return "m_role"
}

type MRoleRequest struct {
	Id       uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	Name     string `form:"name" json:"name" xml:"name" gorm:"size:20;type:varchar(20)" binding:"max=20"`
	Code     string `form:"code" json:"code" xml:"code" gorm:"size:20;type:varchar(20)" binding:"max=20"`
	Level    int    `form:"level" json:"level" xml:"level" gorm:"type:tinyint"`
	IsDelete *bool  `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}

func (req *MRoleRequest) ToModelNew(mUserId uint) *MRole {
	bool_true := false
	return &MRole{
		Id:        req.Id,
		Name:      req.Name,
		Code:      req.Code,
		Level:     req.Level,
		CreatedOn: dto.JSONTime{Time: time.Now()},
		CreatedBy: mUserId,
		IsDelete:  &bool_true,
	}
}
