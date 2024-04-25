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
	GatewayFindLocation viacep.GatewayInterface
}

func NewFindLocationUseCase(GatewayFindLocatoin viacep.GatewayInterface) *FindLocationUseCase {
	return &FindLocationUseCase{
		GatewayFindLocation: GatewayFindLocatoin,
	}
}

func (c *FindLocationUseCase) Execute(ctx context.Context, input Input) (Output, error) {
	location, err := c.GatewayFindLocation.FindLocation(ctx, input.CEP)
	if err != nil {
		return Output{}, err
	}

	return Output{
		CEP:  location.CEP,
		City: location.City,
	}, nil
}
