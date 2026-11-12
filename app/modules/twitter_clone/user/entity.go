package user

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type User struct {
	Id              uint         `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	Username        *string      `form:"username" json:"username" xml:"username" gorm:"size:30;type:varchar(30);uniqueIndex" validate:"max=30"`
	Email           *string      `form:"email" json:"email" xml:"email" gorm:"size:255;type:varchar(255);uniqueIndex" validate:"max=255"`
	DisplayName     *string      `form:"displayName" json:"displayName" xml:"displayName" gorm:"size:50;type:varchar(50)" validate:"max=50"`
	PasswordHash    *string      `form:"passwordHash" json:"passwordHash" xml:"passwordHash" gorm:"size:255;type:varchar(255)" validate:"max=255"`
	Bio             *string      `form:"bio" json:"bio" xml:"bio" gorm:"size:160;type:varchar(160)" validate:"max=160"`
	ProfileImageUrl *string      `form:"profileImageUrl" json:"profileImageUrl" xml:"profileImageUrl" gorm:"size:512;type:varchar(512)" validate:"max=512"`
	CreatedAt       dto.JSONTime `form:"createdAt" json:"createdAt" xml:"createdAt" gorm:"not null;type:timestamp" swaggertype:"string" example:"2024-02-16 10:33:10"`
}

func (User) TableName() string {
	return "users"
}

type UserRequest struct {
	Id              *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	Username        *string `form:"username" json:"username" xml:"username" gorm:"size:30;type:varchar(30)" validate:"required,max=30"`
	Email           *string `form:"email" json:"email" xml:"email" gorm:"size:255;type:varchar(255)" validate:"required,max=255,email"`
	DisplayName     *string `form:"displayName" json:"displayName" xml:"displayName" gorm:"size:50;type:varchar(50)" validate:"required,max=50"`
	PasswordHash    *string `form:"passwordHash" json:"passwordHash" xml:"passwordHash" gorm:"size:255;type:varchar(255)" validate:"required,max=255"`
	Bio             *string `form:"bio" json:"bio" xml:"bio" gorm:"size:160;type:varchar(160)" validate:"max=160"`
	ProfileImageUrl *string `form:"profileImageUrl" json:"profileImageUrl" xml:"profileImageUrl" gorm:"size:512;type:varchar(512)" validate:"max=512"`
}
type UserResponse struct {
	Id              *uint        `form:"id" json:"id" xml:"id"`
	Username        *string      `form:"username" json:"username" xml:"username" validate:"max=30"`
	Email           *string      `form:"email" json:"email" xml:"email" validate:"max=255"`
	DisplayName     *string      `form:"displayName" json:"displayName" xml:"displayName" validate:"max=50"`
	Bio             *string      `form:"bio" json:"bio" xml:"bio" validate:"max=160"`
	ProfileImageUrl *string      `form:"profileImageUrl" json:"profileImageUrl" xml:"profileImageUrl" validate:"max=512"`
	CreatedAt       dto.JSONTime `form:"createdAt" json:"createdAt" xml:"createdAt"`
}

func ToUserEntity(req UserRequest, userId uint) User {
	entity := User{
		Username:        req.Username,
		Email:           req.Email,
		DisplayName:     req.DisplayName,
		PasswordHash:    req.PasswordHash,
		Bio:             req.Bio,
		ProfileImageUrl: req.ProfileImageUrl,
		CreatedAt:       dto.JSONTime{Time: time.Now()},
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToUserResponse(entity User) UserResponse {
	res := UserResponse{
		Id:              &entity.Id,
		Username:        entity.Username,
		Email:           entity.Email,
		DisplayName:     entity.DisplayName,
		Bio:             entity.Bio,
		ProfileImageUrl: entity.ProfileImageUrl,
		CreatedAt:       entity.CreatedAt,
	}
	return res
}
func ToUserResponsesParallel(entities []User) []UserResponse {
	numEntities := len(entities)
	responses := make([]UserResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity User) {
			defer wg.Done()
			responses[index] = ToUserResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
