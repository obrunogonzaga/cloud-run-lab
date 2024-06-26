package usecase

import (
	"context"
	"github.com/obrunogonzaga/cloud-run-lab/configs"
	"github.com/obrunogonzaga/cloud-run-lab/internal/repository"
)

type CalculateWeatherInput struct {
	City string `json:"city"`
}

type CalculateWeatherOutput struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

type CalculateWeatherUseCase struct {
	Gateway repository.WeatherRepository
	Config  *configs.Config
}

func NewCalculateWeatherUseCase(Gateway repository.WeatherRepository, Config *configs.Config) *CalculateWeatherUseCase {
	return &CalculateWeatherUseCase{
		Gateway: Gateway,
		Config:  Config,
	}
}

func (c *CalculateWeatherUseCase) Execute(ctx context.Context, input CalculateWeatherInput) (CalculateWeatherOutput, error) {
	weather, err := c.Gateway.GetWeather(ctx, input.City, *c.Config)
	if err != nil {
		return CalculateWeatherOutput{}, err
	}

	return CalculateWeatherOutput{
		Celsius:    weather.Celsius,
		Fahrenheit: weather.Fahrenheit,
		Kelvin:     weather.Kelvin,
	}, nil
}
