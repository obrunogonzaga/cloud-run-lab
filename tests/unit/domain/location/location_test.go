package location

import (
	"github.com/obrunogonzaga/cloud-run-lab/internal/domain/location"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnEmptyCityWhenNewLocationThenReturnError(t *testing.T) {
	// Given
	city := ""

	// When
	_, err := location.NewLocation(city)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, "invalid city", err.Error())
}

func TestGivenAValidCityWhenNewLocationThenReturnLocationWithAllParams(t *testing.T) {
	// Given
	city := "São Paulo"

	// When
	location, err := location.NewLocation(city)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, city, location.City)
}
