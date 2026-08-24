package m_biodata_attachment

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type MBiodataAttachmentService interface {
	Get(id uint) (*MBiodataAttachment, error)
	Create(payload *MBiodataAttachmentRequest, mUserId uint) error
	Update(payload *MBiodataAttachmentRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MBiodataAttachmentServiceImpl struct {
	repo MBiodataAttachmentRepository
}

func NewMBiodataAttachmentService(repo MBiodataAttachmentRepository) MBiodataAttachmentService {
	return &MBiodataAttachmentServiceImpl{
		repo: repo,
	}
}

func (s *MBiodataAttachmentServiceImpl) Get(id uint) (*MBiodataAttachment, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *MBiodataAttachmentServiceImpl) Create(payload *MBiodataAttachmentRequest, mUserId uint) error {

	data := ToMBiodataAttachmentEntity(*payload, nil, mUserId)

	if payload.Id == nil {
		return s.repo.Create(&data)
	}

	_, err := s.repo.Get(*payload.Id)
	if err == nil {
		return errors.New("data exists")
	}

	data.Id = *payload.Id
	return s.repo.Create(&data)
}

func (s *MBiodataAttachmentServiceImpl) Update(payload *MBiodataAttachmentRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.BiodataId = payload.BiodataId
	existing.FileName = payload.FileName
	existing.FilePath = payload.FilePath
	existing.FileSize = payload.FileSize
	existing.File = payload.File
	existing.ModifiedBy = &mUserId
	existing.ModifiedOn = &dto.JSONTime{Time: time.Now()}
	existing.DeletedBy = nil
	existing.DeletedOn = nil

	if payload.IsDelete != nil && *payload.IsDelete {
		existing.DeletedBy = &mUserId
		existing.DeletedOn = &dto.JSONTime{Time: time.Now()}
		existing.IsDelete = *payload.IsDelete
	}

	return s.repo.Update(existing)
}

func (s *MBiodataAttachmentServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MBiodataAttachmentServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
