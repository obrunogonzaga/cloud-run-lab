package main

import (
	"github.com/obrunogonzaga/cloud-run-lab/configs"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/web"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/web/webserver"
	"github.com/obrunogonzaga/cloud-run-lab/internal/repository"
	locatiionService "github.com/obrunogonzaga/cloud-run-lab/internal/service"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	locationRepo := repository.NewLocationRepository(client)
	locationService := locatiionService.NewLocationService(locationRepo)
	weather := repository.NewWeatherAPI(client)
	handler := web.NewHandler(locationService, weather, config)

	//TODO: Implementar a injeção de dependência com o wire
	restServer := webserver.NewWebServer(config.WebServerPort)
	restServer.AddHandler("/weather", handler.Execute)

	restServer.Start()

}
