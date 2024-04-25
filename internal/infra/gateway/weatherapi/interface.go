package weatherapi

import (
	"context"
	"github.com/obrunogonzaga/cloud-run-lab/internal/entity"
)

type GatewayInterface interface {
	GetWeather(ctx context.Context, city string) (*entity.Temperature, error)
}
