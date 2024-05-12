package service

import (
	"context"
	"github.com/obrunogonzaga/cloud-run-lab/internal/domain/location"
)

type LocationService interface {
	FindLocationByZipCode(ctx context.Context, cep string) (*location.Location, error)
}
