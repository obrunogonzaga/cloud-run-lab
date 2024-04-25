package web

import (
	"encoding/json"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/gateway/viacep"
	"github.com/obrunogonzaga/cloud-run-lab/internal/usecase"
	"net/http"
)

type GetLocationHandler struct {
	LocationService viacep.GatewayInterface
}

func NewGetLocationHandler(LocationService viacep.GatewayInterface) *GetLocationHandler {
	return &GetLocationHandler{
		LocationService: LocationService,
	}
}

func (h *GetLocationHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var dto usecase.Input
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	findLocation := usecase.NewFindLocationUseCase(h.LocationService)
	output, err := findLocation.Execute(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
