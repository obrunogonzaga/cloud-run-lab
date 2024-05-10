package service

import (
	"context"
	"github.com/obrunogonzaga/cloud-run-lab/internal/domain/location"
)

type locationServiceImpl struct {
	repo location.LocationRepository
}

func NewLocationService(repo location.LocationRepository) location.LocationService {
	return &locationServiceImpl{
		repo: repo,
	}
}

func (v *locationServiceImpl) FindLocationByZipCode(ctx context.Context, cep string) (*location.Location, error) {
	return v.repo.FindCityByZipCode(ctx, cep)
}
