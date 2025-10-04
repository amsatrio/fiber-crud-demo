package application

import (
	"errors"
	"fiber-crud-demo/dto"
	"fiber-crud-demo/dto/request"
	"fiber-crud-demo/dto/response"
	"fiber-crud-demo/internal/domain"
	"fmt"
	"io"
	"strconv"
	"time"
)

type MBiodataService interface {
	Get(id uint) (*domain.MBiodata, error)
	Create(payload *domain.MBiodataRequest, mUserId uint) error
	Update(payload *domain.MBiodataRequest, mUserId uint) error
	Delete(id uint) error
	GetPage(
		sortRequest []request.Sort,
		filterRequest []request.Filter,
		searchRequest string,
		pageInt int,
		sizeInt64 int64,
		sizeInt int) (*response.Page, error)
}

type MBiodataServiceImpl struct {
	repo     domain.MBiodataRepository
	fileRepo domain.MFileRepository
}

func NewMBiodataService(repo domain.MBiodataRepository, fileRepo domain.MFileRepository) MBiodataService {
	return &MBiodataServiceImpl{
		repo:     repo,
		fileRepo: fileRepo,
	}
}

func (s *MBiodataServiceImpl) Get(id uint) (*domain.MBiodata, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if existing.ImagePath != nil && *existing.ImagePath != "" {
		parsed, err := strconv.ParseUint(*existing.ImagePath, 10, 0)
		if err != nil {
			return nil, err
		}
		fileId := uint(parsed)
		s, err := s.fileRepo.Stream(fileId)
		if err != nil {
			return nil, err
		}

		defer s.Close()
		existing.Image, err = io.ReadAll(s)
		if err != nil {
			return nil, err
		}
	}

	return existing, nil
}

func (s *MBiodataServiceImpl) Create(payload *domain.MBiodataRequest, mUserId uint) error {
	bool_true := false
	data := &domain.MBiodata{
		Id:          payload.Id,
		Fullname:    payload.Fullname,
		MobilePhone: payload.MobilePhone,
		Image:       nil,
		ImagePath:   payload.ImagePath,
		CreatedOn:   dto.JSONTime{Time: time.Now()},
		CreatedBy:   mUserId,
		IsDelete:    &bool_true,
	}

	id := time.Now().Unix()
	if payload.Image != nil {
		err := s.fileRepo.Upload(payload.Image, id)
		if err == nil {
			idString := fmt.Sprintf("%d", id)
			data.ImagePath = &idString
		}
	}

	if payload.Id == nil {
		return s.repo.Create(data)
	}

	_, err := s.repo.Get(*payload.Id)
	if err == nil {
		return errors.New("data exists")
	}

	data.Id = nil

	return s.repo.Create(data)
}

func (s *MBiodataServiceImpl) Update(payload *domain.MBiodataRequest, mUserId uint) error {

	if payload.Id == nil {
		return errors.New("invalid payload")
	}

	existing, err := s.repo.Get(*payload.Id)
	if err != nil {
		return err
	}

	id := time.Now().Unix()
	idString := ""
	if payload.Image != nil {
		// delete existing file
		if existing.ImagePath != nil && *existing.ImagePath != "" {
			err = s.fileRepo.Delete(*existing.ImagePath)
			if err != nil {
				fmt.Printf("%v", err)
				return err
			}
		}

		err := s.fileRepo.Upload(payload.Image, id)
		if err != nil {
			fmt.Printf("%v", err)
			return err
		}
		idString = fmt.Sprintf("%d", id)
		existing.ImagePath = &idString
	}

	existing.Fullname = payload.Fullname
	existing.MobilePhone = payload.MobilePhone
	existing.ModifiedBy = &mUserId
	existing.ModifiedOn = &dto.JSONTime{Time: time.Now()}
	if *payload.IsDelete {
		existing.DeletedBy = &mUserId
		existing.DeletedOn = &dto.JSONTime{Time: time.Now()}
		existing.IsDelete = payload.IsDelete
	}

	return s.repo.Update(existing)
}

func (s *MBiodataServiceImpl) Delete(id uint) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *MBiodataServiceImpl) GetPage(
	sortRequest []request.Sort,
	filterRequest []request.Filter,
	searchRequest string,
	pageInt int,
	sizeInt64 int64,
	sizeInt int) (*response.Page, error) {

	return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
