package main

import (
	"github.com/obrunogonzaga/cloud-run-lab/configs"
	viacep "github.com/obrunogonzaga/cloud-run-lab/internal/infra/gateway/viacep/impl"
	weatherapi "github.com/obrunogonzaga/cloud-run-lab/internal/infra/gateway/weatherapi/impl"
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
	viaCEP := viacep.NewViaCEP(client)
	weather := weatherapi.NewWeatherAPI(client)
	handler := web.NewHandler(viaCEP, weather, config)

	//TODO: Implementar a injeção de dependência com o wire
	restServer := webserver.NewWebServer(config.WebServerPort)
	restServer.AddHandler("/weather", handler.Execute)

	restServer.Start()

}
