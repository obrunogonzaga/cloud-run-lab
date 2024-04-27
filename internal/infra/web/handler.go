package web

import (
	"encoding/json"
	"github.com/obrunogonzaga/cloud-run-lab/configs"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/gateway/viacep"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/gateway/weatherapi"
	"github.com/obrunogonzaga/cloud-run-lab/internal/usecase"
	"net/http"
)

type Handler struct {
	LocationService viacep.GatewayInterface
	WeatherService  weatherapi.GatewayInterface
	Config          *configs.Config
}

func NewHandler(LocationService viacep.GatewayInterface, WeatherService weatherapi.GatewayInterface, Config *configs.Config) *Handler {
	return &Handler{
		LocationService: LocationService,
		WeatherService:  WeatherService,
		Config:          Config,
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

	weatherUseCase := usecase.NewCalculateWeatherUseCase(h.WeatherService, h.Config)
	weatherOutput, err := weatherUseCase.Execute(r.Context(), usecase.CalculateWeatherInput{City: output.City})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(weatherOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
