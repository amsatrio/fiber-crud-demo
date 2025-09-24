package domain

import (
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
)

type MBiodataRepository interface {
	GetMBiodata(id uint) (*MBiodata, error)
	CreateMBiodata(mBiodata *MBiodataRequest, mUserId uint) error
	UpdateMBiodata(mBiodata *MBiodataRequest, mUserId uint) error
	DeleteMBiodata(id uint, mUserId uint) error
	GetPageMBiodata(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}
