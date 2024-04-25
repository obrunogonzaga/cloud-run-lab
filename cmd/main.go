package main

import (
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/gateway/viacep/impl"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/web"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/web/webserver"
	"net/http"
)

func main() {

	client := &http.Client{}
	viaCEP := impl.NewViaCEP(client)
	handler := web.NewHandler(viaCEP)

	//TODO: Implementar a injeção de dependência com o wire
	//TODO: Implementar captura de variaveis de ambiente
	restServer := webserver.NewWebServer(":8080")
	restServer.AddHandler("/weather", handler.Execute)

	restServer.Start()

}
