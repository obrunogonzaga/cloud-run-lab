package service

import (
	"context"
	"github.com/obrunogonzaga/cloud-run-lab/internal/domain/location"
	"github.com/obrunogonzaga/cloud-run-lab/internal/repository"
)

type locationServiceImpl struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationServiceImpl{
		repo: repo,
	}
}

func (v *locationServiceImpl) FindLocationByZipCode(ctx context.Context, cep string) (*location.Location, error) {
	return v.repo.FindCityByZipCode(ctx, cep)
}
