package entrypoints

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	globalconstants "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/global/constants"
)

type APIRestServer struct {
	echoInstance *echo.Echo
	globalGroup  *echo.Group
	port         string
	routes       func(*echo.Group)
}

func NewAPIRestServer(port, globalPrefix string, routes func(*echo.Group)) *APIRestServer {
	e := echo.New()
	return &APIRestServer{
		echoInstance: e,
		globalGroup:  e.Group(globalPrefix),
		port:         port,
		routes:       routes,
	}
}

func (api *APIRestServer) RunServer() {

	api.echoInstance.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		CustomTimeFormat: globalconstants.DateAndHourFormat,
		Format:           "[${time_custom}] - method=${method}, uri=${uri}, status=${status}\n",
	}))

	api.routes(api.globalGroup)
	routes := api.echoInstance.Routes()
	for _, route := range routes {
		log.Info().Str("method", route.Method).Str("path", route.Path).Msg("echo route")
	}
	api.echoInstance.Logger.Fatal(api.echoInstance.Start(":" + api.port))
}
