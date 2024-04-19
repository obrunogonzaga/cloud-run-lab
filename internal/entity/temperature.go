package entity

import "errors"

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

func NewTemperature(celsius float64) (*Temperature, error) {
	temperature := &Temperature{
		Celsius: celsius,
	}
	err := temperature.isValid()
	if err != nil {
		return nil, err
	}

	temperature.ConvertFahrenheit()
	temperature.ConvertKelvin()

	return temperature, nil
}

func (t *Temperature) isValid() error {
	if t.Celsius <= -273.15 {
		return errors.New("invalid celsius")
	}
	return nil
}

func (t *Temperature) ConvertFahrenheit() {
	t.Fahrenheit = t.Celsius*1.8 + 32
}

func (t *Temperature) ConvertKelvin() {
	t.Kelvin = t.Celsius + 273.15
}
