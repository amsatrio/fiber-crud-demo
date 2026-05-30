package m_medical_item

import (
	"errors"
	"io"
	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type MMedicalItemService interface {
	Get(id uint) (*MMedicalItem, error)
	Create(payload *MMedicalItemRequest, mUserId uint) error
	Update(payload *MMedicalItemRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MMedicalItemServiceImpl struct {
	repo MMedicalItemRepository
}

func NewMMedicalItemService(repo MMedicalItemRepository) MMedicalItemService {
	return &MMedicalItemServiceImpl{
		repo: repo,
	}
}

func (s *MMedicalItemServiceImpl) Get(id uint) (*MMedicalItem, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *MMedicalItemServiceImpl) Create(payload *MMedicalItemRequest, mUserId uint) error {
	var dataBytes []byte
	if payload.Image != nil {
		file, err := payload.Image.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		dataBytes, err = io.ReadAll(file)
		if err != nil {
			return err
		}
	}

	data := ToMMedicalItemEntity(*payload, dataBytes, mUserId)

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

func (s *MMedicalItemServiceImpl) Update(payload *MMedicalItemRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	var dataBytes []byte
	if payload.Image != nil {
		file, err := payload.Image.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		dataBytes, err = io.ReadAll(file)
		if err != nil {
			return err
		}
	}

	existing.Name = payload.Name
	existing.MedicalItemCategoryId = payload.MedicalItemCategoryId
	existing.Composition = payload.Composition
	existing.MedicalItemSegmentationId = payload.MedicalItemSegmentationId
	existing.Manufacturer = payload.Manufacturer
	existing.Indication = payload.Indication
	existing.Dosage = payload.Dosage
	existing.Directions = payload.Directions
	existing.Contraindication = payload.Contraindication
	existing.Caution = payload.Caution
	existing.Packaging = payload.Packaging
	existing.PriceMax = payload.PriceMax
	existing.PriceMin = payload.PriceMin
	existing.ImagePath = payload.ImagePath
	existing.Image = &dataBytes
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

func (s *MMedicalItemServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MMedicalItemServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
