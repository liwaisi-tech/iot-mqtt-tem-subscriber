package routes

import (
	"github.com/labstack/echo/v4"

	pinghandler "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/entrypoints/http/echo/handlers/ping"
	climatedata "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/entrypoints/http/echo/routes/climate_data"
)

func EchoRoutes(group *echo.Group) {

	group.GET("ping", pinghandler.PingHandler())
	climatedata.AddClimateDataEchoRoutes(group)
}
