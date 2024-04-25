package impl

import (
	"context"
	"errors"
	"github.com/obrunogonzaga/cloud-run-lab/internal/entity"
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
	viaCEP := NewViaCEP(client)
	location, err := viaCEP.FindLocation(context.Background(), "01001-000")

	assert.NoError(t, err)
	assert.Equal(t, &entity.Location{CEP: "01001-000", City: "São Paulo"}, location)
}

func TestFindLocation_InvalidCEP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := server.Client()
	viaCEP := NewViaCEP(client)
	_, err := viaCEP.FindLocation(context.Background(), "invalid")

	assert.Error(t, err)
}

func TestFindLocation_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		time.Sleep(100 * time.Millisecond)
		rw.Write([]byte(`{"cep": "01001-000", "localidade": "São Paulo"}`))
	}))
	defer server.Close()

	client := server.Client()
	viaCEP := NewViaCEP(client)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	_, err := viaCEP.FindLocation(ctx, "01001-000")

	assert.True(t, errors.Is(err, context.DeadlineExceeded))
}
