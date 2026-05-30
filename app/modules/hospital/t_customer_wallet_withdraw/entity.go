package t_customer_wallet_withdraw

import (
	"sync"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
)

type TCustomerWalletWithdraw struct {
	Id                     uint          `form:"id" json:"id" xml:"id" gorm:"primary_key;autoIncrement;not null;type:bigint" validate:"required"`
	CustomerId             uint          `form:"customerId" json:"customerId" xml:"customerId" gorm:"not null;type:bigint"`
	WalletDefaultNominalId *uint         `form:"walletDefaultNominalId" json:"walletDefaultNominalId" xml:"walletDefaultNominalId" gorm:"type:bigint"`
	Amount                 int           `form:"amount" json:"amount" xml:"amount" gorm:"not null;type:int"`
	BankName               string        `form:"bankName" json:"bankName" xml:"bankName" gorm:"not null;size:50;type:varchar(50)" validate:"max=50"`
	AccountNumber          string        `form:"accountNumber" json:"accountNumber" xml:"accountNumber" gorm:"not null;size:50;type:varchar(50)" validate:"max=50"`
	AccountName            string        `form:"accountName" json:"accountName" xml:"accountName" gorm:"not null;size:255;type:varchar(255)" validate:"max=255"`
	Otp                    int           `form:"otp" json:"otp" xml:"otp" gorm:"not null;type:int"`
	CreatedBy              uint          `form:"createdBy" json:"createdBy" xml:"createdBy" gorm:"not null;type:bigint"`
	CreatedOn              dto.JSONTime  `form:"createdOn" json:"createdOn" xml:"createdOn" gorm:"not null;type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	ModifiedBy             *uint         `form:"modifiedBy" json:"modifiedBy" xml:"modifiedBy" gorm:"type:bigint"`
	ModifiedOn             *dto.JSONTime `form:"modifiedOn" json:"modifiedOn" xml:"modifiedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	DeletedBy              *uint         `form:"deletedBy" json:"deletedBy" xml:"deletedBy" gorm:"type:bigint"`
	DeletedOn              *dto.JSONTime `form:"deletedOn" json:"deletedOn" xml:"deletedOn" gorm:"type:datetime" swaggertype:"string" example:"2024-02-16 10:33:10"`
	IsDelete               bool          `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"not null;type:boolean;comment:default FALSE"`
}

func (TCustomerWalletWithdraw) TableName() string {
	return "t_customer_wallet_withdraw"
}

type TCustomerWalletWithdrawRequest struct {
	Id                     *uint  `form:"id" json:"id" xml:"id" gorm:"primary_key;not null;type:bigint;comment:Auto increment"`
	CustomerId             uint   `form:"customerId" json:"customerId" xml:"customerId" gorm:"type:bigint"`
	WalletDefaultNominalId *uint  `form:"walletDefaultNominalId" json:"walletDefaultNominalId" xml:"walletDefaultNominalId" gorm:"type:bigint"`
	Amount                 int    `form:"amount" json:"amount" xml:"amount" gorm:"type:int"`
	BankName               string `form:"bankName" json:"bankName" xml:"bankName" gorm:"size:50;type:varchar(50)" validate:"max=50"`
	AccountNumber          string `form:"accountNumber" json:"accountNumber" xml:"accountNumber" gorm:"size:50;type:varchar(50)" validate:"max=50"`
	AccountName            string `form:"accountName" json:"accountName" xml:"accountName" gorm:"size:255;type:varchar(255)" validate:"max=255"`
	Otp                    int    `form:"otp" json:"otp" xml:"otp" gorm:"type:int"`
	IsDelete               *bool  `form:"isDelete" json:"isDelete" xml:"isDelete" gorm:"type:boolean;comment:default FALSE"`
}
type TCustomerWalletWithdrawResponse struct {
	Id                     *uint  `form:"id" json:"id" xml:"id"`
	CustomerId             uint   `form:"customerId" json:"customerId" xml:"customerId"`
	WalletDefaultNominalId *uint  `form:"walletDefaultNominalId" json:"walletDefaultNominalId" xml:"walletDefaultNominalId"`
	Amount                 int    `form:"amount" json:"amount" xml:"amount"`
	BankName               string `form:"bankName" json:"bankName" xml:"bankName" validate:"max=50"`
	AccountNumber          string `form:"accountNumber" json:"accountNumber" xml:"accountNumber" validate:"max=50"`
	AccountName            string `form:"accountName" json:"accountName" xml:"accountName" validate:"max=255"`
	Otp                    int    `form:"otp" json:"otp" xml:"otp"`
	IsDelete               bool   `form:"isDelete" json:"isDelete" xml:"isDelete"`
}

func ToTCustomerWalletWithdrawEntity(req TCustomerWalletWithdrawRequest, imageData []byte, userId uint) TCustomerWalletWithdraw {
	entity := TCustomerWalletWithdraw{
		CustomerId:             req.CustomerId,
		WalletDefaultNominalId: req.WalletDefaultNominalId,
		Amount:                 req.Amount,
		BankName:               req.BankName,
		AccountNumber:          req.AccountNumber,
		AccountName:            req.AccountName,
		Otp:                    req.Otp,
		CreatedBy:              userId,
		CreatedOn:              dto.JSONTime{Time: time.Now()},
		IsDelete:               false,
	}
	if req.Id != nil {
		entity.Id = *req.Id
	}
	return entity
}

func ToTCustomerWalletWithdrawResponse(entity TCustomerWalletWithdraw) TCustomerWalletWithdrawResponse {
	res := TCustomerWalletWithdrawResponse{
		Id:                     &entity.Id,
		CustomerId:             entity.CustomerId,
		WalletDefaultNominalId: entity.WalletDefaultNominalId,
		Amount:                 entity.Amount,
		BankName:               entity.BankName,
		AccountNumber:          entity.AccountNumber,
		AccountName:            entity.AccountName,
		Otp:                    entity.Otp,
		IsDelete:               entity.IsDelete,
	}
	return res
}
func ToTCustomerWalletWithdrawResponsesParallel(entities []TCustomerWalletWithdraw) []TCustomerWalletWithdrawResponse {
	numEntities := len(entities)
	responses := make([]TCustomerWalletWithdrawResponse, numEntities)

	var wg sync.WaitGroup
	wg.Add(numEntities)

	for i, e := range entities {
		go func(index int, entity TCustomerWalletWithdraw) {
			defer wg.Done()
			responses[index] = ToTCustomerWalletWithdrawResponse(entity)
		}(i, e)
	}

	wg.Wait()
	return responses
}
