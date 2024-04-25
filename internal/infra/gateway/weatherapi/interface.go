package weatherapi

type GatewayInterface interface {
	GetWeather(city string) (string, error)
}
