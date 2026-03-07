package m_biodata

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    "io"
    "time"
)

type MBiodataService interface {
    Get(id uint) (*MBiodata, error)
    Create(payload *MBiodataRequest, mUserId uint) error
    Update(payload *MBiodataRequest, mUserId uint) error
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
    repo     MBiodataRepository
}

func NewMBiodataService(repo MBiodataRepository) MBiodataService {
    return &MBiodataServiceImpl{
        repo:     repo,
    }
}

func (s *MBiodataServiceImpl) Get(id uint) (*MBiodata, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *MBiodataServiceImpl) Create(payload *MBiodataRequest, mUserId uint) error {
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

    data := ToMBiodataEntity(*payload, dataBytes, mUserId)

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

func (s *MBiodataServiceImpl) Update(payload *MBiodataRequest, mUserId uint) error {
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

    existing.Fullname        = payload.Fullname
    existing.MobilePhone     = payload.MobilePhone
    existing.ImagePath       = payload.ImagePath
    existing.Image           = &dataBytes
    existing.ModifiedBy = &mUserId
    existing.ModifiedOn = &dto.JSONTime{Time: time.Now()}
    existing.DeletedBy  = nil
    existing.DeletedOn  = nil

    if payload.IsDelete != nil && *payload.IsDelete {
        existing.DeletedBy = &mUserId
        existing.DeletedOn = &dto.JSONTime{Time: time.Now()}
        existing.IsDelete = *payload.IsDelete
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
