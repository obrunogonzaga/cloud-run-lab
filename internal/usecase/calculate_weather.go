package usecase

import (
	"context"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/gateway/weatherapi"
)

type CalculateWeatherInput struct {
	City string `json:"city"`
}

type CalculateWeatherOutput struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}

type CalculateWeatherUseCase struct {
	Gateway weatherapi.GatewayInterface
}

func NewCalculateWeatherUseCase(Gateway weatherapi.GatewayInterface) *CalculateWeatherUseCase {
	return &CalculateWeatherUseCase{
		Gateway: Gateway,
	}
}

func (c *CalculateWeatherUseCase) Execute(ctx context.Context, input CalculateWeatherInput) (CalculateWeatherOutput, error) {
	weather, err := c.Gateway.GetWeather(ctx, input.City)
	if err != nil {
		return CalculateWeatherOutput{}, err
	}

	return CalculateWeatherOutput{
		Celsius:    weather.Celsius,
		Fahrenheit: weather.Fahrenheit,
		Kelvin:     weather.Kelvin,
	}, nil
}
