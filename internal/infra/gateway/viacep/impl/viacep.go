package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/obrunogonzaga/cloud-run-lab/internal/entity"
	"net/http"
)

type Output struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ViaCEP struct {
	Client *http.Client
}

func NewViaCEP(client *http.Client) *ViaCEP {
	return &ViaCEP{
		Client: client,
	}
}

func (v *ViaCEP) FindLocation(ctx context.Context, cep string) (*entity.Location, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := v.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var output Output
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, err
	}

	return &entity.Location{
		CEP:  output.Cep,
		City: output.Localidade,
	}, nil
}
