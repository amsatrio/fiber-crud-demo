package t_reset_password

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type TResetPassword struct {
	Id          uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	OldPassword *string       `form:"oldPassword" json:"oldPassword" xml:"oldPassword" gorm:"size:255;type:varchar(255)" validate:"max=255"`
	NewPassword *string       `form:"newPassword" json:"newPassword" xml:"newPassword" gorm:"size:255;type:varchar(255)" validate:"max=255"`
	ResetFor    *string       `form:"resetFor" json:"resetFor" xml:"resetFor" gorm:"size:20;type:varchar(20)" validate:"max=20"`
	CreatedBy   uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn   dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy  *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn  *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy   *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn   *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete    bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TResetPassword) TableName() string {
	return "t_reset_password"
}

type TResetPasswordRequest struct {
	Id          *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	OldPassword *string `form:"oldPassword" json:"oldPassword" xml:"oldPassword" gorm:"size:255;type:varchar(255)" validate:"max=255"`
	NewPassword *string `form:"newPassword" json:"newPassword" xml:"newPassword" gorm:"size:255;type:varchar(255)" validate:"max=255"`
	ResetFor    *string `form:"resetFor" json:"resetFor" xml:"resetFor" gorm:"size:20;type:varchar(20)" validate:"max=20"`
	IsDelete    *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TResetPasswordResponse struct {
	Id          *uint   `form:"id" json:"id" xml:"id"`
	OldPassword *string `form:"oldPassword" json:"oldPassword" xml:"oldPassword" validate:"max=255"`
	NewPassword *string `form:"newPassword" json:"newPassword" xml:"newPassword" validate:"max=255"`
	ResetFor    *string `form:"resetFor" json:"resetFor" xml:"resetFor" validate:"max=20"`
	IsDelete    bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToTResetPasswordEntity(req TResetPasswordRequest, imageData []byte, userId uint) TResetPassword {
	entity := TResetPassword{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
		ResetFor:    req.ResetFor,
		CreatedBy:   userId,
		CreatedOn:   dto.JSONTime{Time: time.Now()},
		IsDelete:    false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToTResetPasswordResponse(entity TResetPassword) TResetPasswordResponse {
	res := TResetPasswordResponse{
		Id:          &entity.Id,
		OldPassword: entity.OldPassword,
		NewPassword: entity.NewPassword,
		ResetFor:    entity.ResetFor,
		IsDelete:    entity.IsDelete,
	}
	return res
}
func ToTResetPasswordResponsesParallel(entities []TResetPassword) []TResetPasswordResponse {
	numEntities := len(entities)
	responses := make([]TResetPasswordResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity TResetPassword) {
			defer wg.Done()
			responses[index] = ToTResetPasswordResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
