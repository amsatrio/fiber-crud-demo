package m_customer_member

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type MCustomerMember struct {
	Id                 uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	ParentBiodataId    *uint         `form:"parentBiodataId" json:"parentBiodataId" xml:"parentBiodataId" gorm:"type:bigint;comment:Biodata/pasien/orang yg mendaftarkan"`
	CustomerId         *uint         `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
	CustomerRelationId *uint         `form:"customerRelationId" json:"customerRelationId" xml:"customerRelationId" gorm:"type:bigint"`
	CreatedBy          uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn          dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy         *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn         *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy          *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn          *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete           bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (MCustomerMember) TableName() string {
	return "m_customer_member"
}

type MCustomerMemberRequest struct {
	Id                 *uint `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	ParentBiodataId    *uint `form:"parentBiodataId" json:"parentBiodataId" xml:"parentBiodataId" gorm:"type:bigint"`
	CustomerId         *uint `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
	CustomerRelationId *uint `form:"customerRelationId" json:"customerRelationId" xml:"customerRelationId" gorm:"type:bigint"`
	IsDelete           *bool `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type MCustomerMemberResponse struct {
	Id                 *uint `form:"id" json:"id" xml:"id"`
	ParentBiodataId    *uint `form:"parentBiodataId" json:"parentBiodataId" xml:"parentBiodataId"`
	CustomerId         *uint `form:"customerId" json:"customerId" xml:"customerId"`
	CustomerRelationId *uint `form:"customerRelationId" json:"customerRelationId" xml:"customerRelationId"`
	IsDelete           bool  `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToMCustomerMemberEntity(req MCustomerMemberRequest, imageData []byte, userId uint) MCustomerMember {
	entity := MCustomerMember{
		ParentBiodataId:    req.ParentBiodataId,
		CustomerId:         req.CustomerId,
		CustomerRelationId: req.CustomerRelationId,
		CreatedBy:          userId,
		CreatedOn:          dto.JSONTime{Time: time.Now()},
		IsDelete:           false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToMCustomerMemberResponse(entity MCustomerMember) MCustomerMemberResponse {
	res := MCustomerMemberResponse{
		Id:                 &entity.Id,
		ParentBiodataId:    entity.ParentBiodataId,
		CustomerId:         entity.CustomerId,
		CustomerRelationId: entity.CustomerRelationId,
		IsDelete:           entity.IsDelete,
	}
	return res
}
func ToMCustomerMemberResponsesParallel(entities []MCustomerMember) []MCustomerMemberResponse {
	numEntities := len(entities)
	responses := make([]MCustomerMemberResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity MCustomerMember) {
			defer wg.Done()
			responses[index] = ToMCustomerMemberResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
