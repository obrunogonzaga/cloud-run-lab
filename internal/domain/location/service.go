package location

import "context"

type LocationService interface {
	FindLocationByZipCode(ctx context.Context, cep string) (*Location, error)
}
