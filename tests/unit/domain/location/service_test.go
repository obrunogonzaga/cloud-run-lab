package location

import (
	"context"
	"errors"
	location2 "github.com/obrunogonzaga/cloud-run-lab/internal/domain/location"
	"github.com/obrunogonzaga/cloud-run-lab/internal/domain/location/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFindLocation_ValidCEP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"cep": "01001-006ga", "localidade": "São Paulo"}`))
	}))
	defer server.Close()

	client := server.Client()
	locationService := service.NewLocationService(client)
	location, err := locationService.FindLocationByZipCode(context.Background(), "01001-000")

	assert.NoError(t, err)
	assert.Equal(t, &location2.Location{CEP: "01001-000", City: "São Paulo"}, location)
}

func TestFindLocation_InvalidCEP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := server.Client()
	locationService := service.NewLocationService(client)
	_, err := locationService.FindLocationByZipCode(context.Background(), "invalid")

	assert.Error(t, err)
}

func TestFindLocation_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		time.Sleep(100 * time.Millisecond)
		rw.Write([]byte(`{"cep": "01001-000", "localidade": "São Paulo"}`))
	}))
	defer server.Close()

	client := server.Client()
	locationService := service.NewLocationService(client)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	_, err := locationService.FindLocationByZipCode(ctx, "01001-000")

	assert.True(t, errors.Is(err, context.DeadlineExceeded))
}
