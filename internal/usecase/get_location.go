package usecase

import (
	"github.com/obrunogonzaga/cloud-run-lab/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

type LocationInputDTO struct {
	CEP string `json:"cep"`
}

func TestGivenAnEmptyCityWhenNewLocationThenReturnError(t *testing.T) {
	// Given
	city := ""

	// When
	_, err := entity.NewLocation(city)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, "invalid city", err.Error())
}
