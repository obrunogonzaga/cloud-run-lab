package web

import (
	"encoding/json"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/gateway/viacep"
	"github.com/obrunogonzaga/cloud-run-lab/internal/usecase"
	"net/http"
)

type Handler struct {
	LocationService viacep.GatewayInterface
}

func NewHandler(LocationService viacep.GatewayInterface) *Handler {
	return &Handler{
		LocationService: LocationService,
	}
}

func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	//defer cancel()

	zipCodeDTO := usecase.Input{
		CEP: r.URL.Query().Get("zipcode"),
	}

	findLocation := usecase.NewFindLocationUseCase(h.LocationService)
	output, err := findLocation.Execute(r.Context(), zipCodeDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
