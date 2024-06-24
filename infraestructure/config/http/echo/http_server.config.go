package config

import (
	"os"

	echoapirest "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/entrypoints/http/echo"
	routes "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/entrypoints/http/echo/routes"
	ports "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/rest"
)

func NewEchoAPIRestAdapter() ports.APIRestPorts {
	return echoapirest.NewAPIRestServer(
		os.Getenv("API_REST_PORT"),
		os.Getenv("API_REST_BASE_URL"),
		routes.EchoRoutes,
	)
}
