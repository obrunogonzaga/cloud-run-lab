package repository

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/obrunogonzaga/cloud-run-lab/internal/domain/location"
	"github.com/stretchr/testify/assert"
)

func TestFindCityByZipCode(t *testing.T) {
	// Create a test server that returns a predefined response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"cep": "01001-000",
			"logradouro": "Praça da Sé",
			"complemento": "lado ímpar",
			"bairro": "Sé",
			"localidade": "São Paulo",
			"uf": "SP",
			"ibge": "3550308",
			"gia": "1004",
			"ddd": "11",
			"siafi": "7107"
		}`))
	}))
	defer ts.Close()

	// Create a new LocationRepository with a mocked HTTP client
	repo := NewLocationRepository(ts.Client())

	// Call the function with a test zip code
	loc, err := repo.FindCityByZipCode(context.Background(), "01001-000")

	// Assert that there was no error and the response was correctly parsed
	assert.NoError(t, err)
	assert.Equal(t, &location.Location{
		CEP:  "01001-000",
		City: "São Paulo",
	}, loc)
}
