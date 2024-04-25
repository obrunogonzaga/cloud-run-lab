package viacep

import (
	"context"
	"github.com/obrunogonzaga/cloud-run-lab/internal/entity"
)

type GatewayInterface interface {
	FindLocation(ctx context.Context, cep string) (*entity.Location, error)
}
