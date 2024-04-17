package entity

import "errors"

type Location struct {
	City       string
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

func NewLocation(city string, celsius float64) (*Location, error) {
	location := &Location{
		City:    city,
		Celsius: celsius,
	}
	err := location.isValid()
	if err != nil {
		return nil, err
	}

	location.ConvertFahrenheit()
	location.ConvertKelvin()

	return location, nil
}

func (l *Location) isValid() error {
	if l.City == "" {
		return errors.New("invalid city")
	}
	if l.Celsius <= -273.15 {
		return errors.New("invalid celsius")
	}
	return nil
}

func (l *Location) ConvertFahrenheit() {
	l.Fahrenheit = l.Celsius*1.8 + 32
}

func (l *Location) ConvertKelvin() {
	l.Kelvin = l.Celsius + 273.15
}
