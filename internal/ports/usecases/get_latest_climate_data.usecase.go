package usecases

import (
	"context"

	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
)

type GetLatestClimateDataUseCasePort interface {
	Execute(ctx context.Context) (climateData *entities.ClimateDataEntity, err error)
}
