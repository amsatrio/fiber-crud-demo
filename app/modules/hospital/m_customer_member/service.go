package m_customer_member

import (
	"errors"

	"time"

	"github.com/amsatrio/fiber-crud-demo/app/dto"
	"github.com/amsatrio/fiber-crud-demo/app/dto/request"
	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

type MCustomerMemberService interface {
	Get(id uint) (*MCustomerMember, error)
	Create(payload *MCustomerMemberRequest, mUserId uint) error
	Update(payload *MCustomerMemberRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MCustomerMemberServiceImpl struct {
	repo MCustomerMemberRepository
}

func NewMCustomerMemberService(repo MCustomerMemberRepository) MCustomerMemberService {
	return &MCustomerMemberServiceImpl{
		repo: repo,
	}
}

func (s *MCustomerMemberServiceImpl) Get(id uint) (*MCustomerMember, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *MCustomerMemberServiceImpl) Create(payload *MCustomerMemberRequest, mUserId uint) error {

	data := ToMCustomerMemberEntity(*payload, nil, mUserId)

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

func (s *MCustomerMemberServiceImpl) Update(payload *MCustomerMemberRequest, mUserId uint) error {
	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	existing.ParentBiodataId = payload.ParentBiodataId
	existing.CustomerId = payload.CustomerId
	existing.CustomerRelationId = payload.CustomerRelationId
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

func (s *MCustomerMemberServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *MCustomerMemberServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {
	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
