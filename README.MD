# Sistema de Clima por CEP

## Descrição

Este projeto é um sistema desenvolvido em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). O sistema é hospedado no Google Cloud Run.

Acesse a URL https://cloudrun-goexpert-za5o6n5xla-uc.a.run.app/weather?zipcode=80320310 para acessar a aplicação no Google Cloud Run.

## Requisitos

- O sistema deve receber um CEP válido de 8 dígitos.
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formatá-las em: Celsius, Fahrenheit, Kelvin.
- O sistema deve responder adequadamente nos seguintes cenários:
  - Em caso de sucesso:
    - Código HTTP: 200
    - Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
  - Em caso de falha, caso o CEP não seja válido (com formato correto):
    - Código HTTP: 422
    - Mensagem: invalid zipcode
  - Em caso de falha, caso o CEP não seja encontrado:
    - Código HTTP: 404
    - Mensagem: can not find zipcode

## Instalação e Uso

Este projeto usa Docker para facilitar a instalação e execução. Siga as etapas abaixo para começar:

1. Primeiro, clone o repositório para a sua máquina local:

```bash
git clone git@github.com:obrunogonzaga/cloud-run-lab.git
```

2. Navegue até o diretório do projeto:
```bash
cd cloud-run-lab
```

3. Construa a imagem Docker:
```bash
docker build -t cloud-run-lab .
```

4. Execute o container Docker:
```bash
docker run -p 8000:8000 cloud-run-lab
```

Agora, o sistema deve estar rodando na porta 8000 do seu localhost.
