package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	global "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/global/constants"
)

func PingHandler() func(c echo.Context) (err error) {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"response":  "pong",
			"timestamp": time.Now().Format(global.TimestampWithTimeZoneFormat),
		})
	}
}
