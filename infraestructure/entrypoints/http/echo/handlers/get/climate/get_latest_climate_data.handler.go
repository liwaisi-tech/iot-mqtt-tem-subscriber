package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
	ports "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/usecases"
)

func GetClimateDataHandler(
	useCase ports.GetLatestClimateDataUseCasePort,
) func(c echo.Context) (err error) {
	return func(c echo.Context) error {
		entity, err := useCase.Execute(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, mapEntityResponse(entity))
	}
}

func mapEntityResponse(entity *entities.ClimateDataEntity) map[string]interface{} {
	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		loc = time.UTC
	}
	return map[string]interface{}{
		"temperature":   entity.Temperature,
		"humidity":      entity.Humidity,
		"last_datetime": entity.CreatedAt.In(loc).Format("2006-01-02 03:04:05 pm"),
	}
}
