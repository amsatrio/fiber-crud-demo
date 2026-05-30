package t_customer_chat_history

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type TCustomerChatHistory struct {
	Id             uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	CustomerChatId *uint         `form:"customerChatId" json:"customerChatId" xml:"customerChatId" gorm:"type:bigint"`
	ChatContent    *string       `form:"chatContent" json:"chatContent" xml:"chatContent" gorm:"type:text"`
	CreatedBy      uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn      dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy     *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn     *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy      *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn      *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete       bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCustomerChatHistory) TableName() string {
	return "t_customer_chat_history"
}

type TCustomerChatHistoryRequest struct {
	Id             *uint   `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	CustomerChatId *uint   `form:"customerChatId" json:"customerChatId" xml:"customerChatId" gorm:"type:bigint"`
	ChatContent    *string `form:"chatContent" json:"chatContent" xml:"chatContent" gorm:"type:text"`
	IsDelete       *bool   `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCustomerChatHistoryResponse struct {
	Id             *uint   `form:"id" json:"id" xml:"id"`
	CustomerChatId *uint   `form:"customerChatId" json:"customerChatId" xml:"customerChatId"`
	ChatContent    *string `form:"chatContent" json:"chatContent" xml:"chatContent"`
	IsDelete       bool    `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToTCustomerChatHistoryEntity(req TCustomerChatHistoryRequest, imageData []byte, userId uint) TCustomerChatHistory {
	entity := TCustomerChatHistory{
		CustomerChatId: req.CustomerChatId,
		ChatContent:    req.ChatContent,
		CreatedBy:      userId,
		CreatedOn:      dto.JSONTime{Time: time.Now()},
		IsDelete:       false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToTCustomerChatHistoryResponse(entity TCustomerChatHistory) TCustomerChatHistoryResponse {
	res := TCustomerChatHistoryResponse{
		Id:             &entity.Id,
		CustomerChatId: entity.CustomerChatId,
		ChatContent:    entity.ChatContent,
		IsDelete:       entity.IsDelete,
	}
	return res
}
func ToTCustomerChatHistoryResponsesParallel(entities []TCustomerChatHistory) []TCustomerChatHistoryResponse {
	numEntities := len(entities)
	responses := make([]TCustomerChatHistoryResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity TCustomerChatHistory) {
			defer wg.Done()
			responses[index] = ToTCustomerChatHistoryResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
