package main

import (
	"github.com/obrunogonzaga/cloud-run-lab/configs"
	"github.com/obrunogonzaga/cloud-run-lab/internal/adapters/repository"
	repository2 "github.com/obrunogonzaga/cloud-run-lab/internal/domain/location/repository"
	locatiionService "github.com/obrunogonzaga/cloud-run-lab/internal/domain/location/service"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/web"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/web/webserver"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	locationRepo := repository2.NewLocationRepository(client)
	locationService := locatiionService.NewLocationService(locationRepo)
	weather := repository.NewWeatherAPI(client)
	handler := web.NewHandler(locationService, weather, config)

	//TODO: Implementar a injeção de dependência com o wire
	restServer := webserver.NewWebServer(config.WebServerPort)
	restServer.AddHandler("/weather", handler.Execute)

	restServer.Start()

}
