package usecase

import (
	"context"
	"github.com/obrunogonzaga/cloud-run-lab/internal/infra/gateway/viacep"
)

type Input struct {
	CEP string `json:"cep"`
}

type Output struct {
	CEP   string `json:"cep"`
	City  string `json:"localidade"`
	State string `json:"uf"`
}

type FindLocationUseCase struct {
	GatewayFindLocatoin viacep.GatewayInterface
}

func NewFindLocationUseCase(GatewayFindLocatoin viacep.GatewayInterface) *FindLocationUseCase {
	return &FindLocationUseCase{
		GatewayFindLocatoin: GatewayFindLocatoin,
	}
}

func (c *FindLocationUseCase) Execute(ctx context.Context, input Input) (Output, error) {
	location, err := c.GatewayFindLocatoin.FindLocation(ctx, input.CEP)
	if err != nil {
		return Output{}, err
	}

	return Output{
		CEP:  location.CEP,
		City: location.City,
	}, nil
}
