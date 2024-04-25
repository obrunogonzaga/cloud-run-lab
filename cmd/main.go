package main

import "github.com/obrunogonzaga/cloud-run-lab/internal/infra/web/webserver"

func main() {

	//TODO: Implementar a injeção de dependência com o wire
	//TODO: Implementar captura de variaveis de ambiente
	restServer := webserver.NewWebServer(":8080")

}
