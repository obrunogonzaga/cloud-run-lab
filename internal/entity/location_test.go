package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnEmptyCityWhenNewLocationThenReturnError(t *testing.T) {
	// Given
	city := ""
	celsius := 10.0

	// When
	_, err := NewLocation(city, celsius)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, "invalid city", err.Error())
}

func TestGivenAnInvalidCelsiusWhenNewLocationThenReturnError(t *testing.T) {
	// Given
	city := "São Paulo"
	celsius := -300.0

	// When
	_, err := NewLocation(city, celsius)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, "invalid celsius", err.Error())
}

func TestGivenAValidParamsWhenNewLocationThenReturnLocationWithAllParams(t *testing.T) {
	// Given
	city := "São Paulo"
	celsius := 10.0

	// When
	location, err := NewLocation(city, celsius)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, city, location.City)
	assert.Equal(t, celsius, location.Celsius)
	assert.Equal(t, 50.0, location.Fahrenheit)
	assert.Equal(t, 283.15, location.Kelvin)
}

func TestGivenAValidCityAndCelsiusWhenConvertFahrenheitThenReturnFahrenheit(t *testing.T) {
	// Given
	city := "São Paulo"
	celsius := 10.0
	location, _ := NewLocation(city, celsius)

	// When
	location.ConvertFahrenheit()

	// Then
	assert.Equal(t, 50.0, location.Fahrenheit)
}

func TestGivenAValidCityAndCelsiusWhenConvertKelvinThenReturnKelvin(t *testing.T) {
	// Given
	city := "São Paulo"
	celsius := 10.0
	location, _ := NewLocation(city, celsius)

	// When
	location.ConvertKelvin()

	// Then
	assert.Equal(t, 283.15, location.Kelvin)
}
