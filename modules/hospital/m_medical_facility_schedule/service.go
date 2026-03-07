package m_medical_facility_schedule

import (
    "errors"
    "fiber-crud-demo/dto"
    "fiber-crud-demo/dto/request"
    "fiber-crud-demo/dto/response"
    
    "time"
)

type MMedicalFacilityScheduleService interface {
    Get(id uint) (*MMedicalFacilitySchedule, error)
    Create(payload *MMedicalFacilityScheduleRequest, mUserId uint) error
    Update(payload *MMedicalFacilityScheduleRequest, mUserId uint) error
    Delete(id uint) error
    GetPage(
        sortRequest []request.Sort,
        filterRequest []request.Filter,
        searchRequest string,
        pageInt int,
        sizeInt64 int64,
        sizeInt int) (*response.Page, error)
}

type MMedicalFacilityScheduleServiceImpl struct {
    repo     MMedicalFacilityScheduleRepository
}

func NewMMedicalFacilityScheduleService(repo MMedicalFacilityScheduleRepository) MMedicalFacilityScheduleService {
    return &MMedicalFacilityScheduleServiceImpl{
        repo:     repo,
    }
}

func (s *MMedicalFacilityScheduleServiceImpl) Get(id uint) (*MMedicalFacilitySchedule, error) {
    existing, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return existing, nil
}

func (s *MMedicalFacilityScheduleServiceImpl) Create(payload *MMedicalFacilityScheduleRequest, mUserId uint) error {

    data := ToMMedicalFacilityScheduleEntity(*payload, nil, mUserId)

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

func (s *MMedicalFacilityScheduleServiceImpl) Update(payload *MMedicalFacilityScheduleRequest, mUserId uint) error {
    if payload.Id == nil {
        return errors.New("invalid payload")
    }

    existing, err := s.repo.Get(*payload.Id)
    if err != nil {
        return err
    }


    existing.MedicalFacilityId = payload.MedicalFacilityId
    existing.Day             = payload.Day
    existing.TimeScheduleStart = payload.TimeScheduleStart
    existing.TimeScheduleEnd = payload.TimeScheduleEnd
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

func (s *MMedicalFacilityScheduleServiceImpl) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}

func (s *MMedicalFacilityScheduleServiceImpl) GetPage(
    sortRequest []request.Sort,
    filterRequest []request.Filter,
    searchRequest string,
    pageInt int,
    sizeInt64 int64,
    sizeInt int) (*response.Page, error) {
    return s.repo.GetPage(sortRequest, filterRequest, searchRequest, pageInt, sizeInt64, sizeInt)
}
