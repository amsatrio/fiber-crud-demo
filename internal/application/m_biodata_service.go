package application

import (
	"fiber-crud-demo/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type MBiodataService struct {
	mBiodataRepo domain.MBiodataRepository
}

func NewMBiodataService(repo domain.MBiodataRepository) *MBiodataService {
	return &MBiodataService{
		mBiodataRepo: repo,
	}
}

func (s *MBiodataService) CreateMBiodata(c *fiber.Ctx) error {
	return s.mBiodataRepo.CreateMBiodata(c)
}
