package routes

import (
	"github.com/labstack/echo/v4"
	getclimatedatahandler "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/entrypoints/http/echo/handlers/get/climate"
	climatedatasrv "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/services/postgres/climate_data"
	usecasegetlatest "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/application/usecases/get_climate_data"
	ports "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/usecases"
	gormpkg "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/pkg/gorm/postgres"
)

func AddClimateDataEchoRoutes(group *echo.Group) {

	group.GET("climate/latest", getclimatedatahandler.GetClimateDataHandler(getLatestClimateDataUsecase()))
}
func getLatestClimateDataUsecase() ports.GetLatestClimateDataUseCasePort {
	gormDB := gormpkg.NewPostgresDBConnection()
	climateDataSrv := climatedatasrv.New(gormDB)
	return usecasegetlatest.New(climateDataSrv)
}
