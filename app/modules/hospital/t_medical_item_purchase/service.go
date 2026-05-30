package t_medical_item_purchase

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"

	"time"
)

type TMedicalItemPurchaseService interface {
	Get(id uint) (*TMedicalItemPurchase, error)
	Create(payload *TMedicalItemPurchaseRequest, mUserId uint) error
	Update(payload *TMedicalItemPurchaseRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type TMedicalItemPurchaseServiceImpl struct {
	repo TMedicalItemPurchaseRepository
}

func NewTMedicalItemPurchaseService(repo TMedicalItemPurchaseRepository) TMedicalItemPurchaseService {
	return &TMedicalItemPurchaseServiceImpl{
		repo: repo,
	}
}

func (s *TMedicalItemPurchaseServiceImpl) Get(id uint) (*TMedicalItemPurchase, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *TMedicalItemPurchaseServiceImpl) Create(payload *TMedicalItemPurchaseRequest, mUserId uint) error {

	data := ToTMedicalItemPurchaseEntity(*payload, nil, mUserId)

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

func (s *TMedicalItemPurchaseServiceImpl) Update(payload *TMedicalItemPurchaseRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.CustomerId = payload.CustomerId
	existing.PaymentMethodId = payload.PaymentMethodId
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

func (s *TMedicalItemPurchaseServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *TMedicalItemPurchaseServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
