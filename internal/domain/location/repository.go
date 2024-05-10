package location

import (
	"context"
	"github.com/obrunogonzaga/cloud-run-lab/configs"
	"github.com/obrunogonzaga/cloud-run-lab/internal/domain/weather"
)

type WeatherRepository interface {
	GetWeather(ctx context.Context, city string, config configs.Config) (*weather.Weather, error)
}

type LocationRepository interface {
	FindCityByZipCode(ctx context.Context, cep string) (*Location, error)
}
