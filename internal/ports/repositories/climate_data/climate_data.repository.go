package repositories

import (
	"context"

	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
)

type ClimateDataRepositoryPort interface {
	CreateClimateData(ctx context.Context, entity *entities.ClimateDataEntity) (newID string, err error)
}
