package ports

import (
	"context"

	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/climate_data"
)

type ClimateDataRepository interface {
	CreateClimateData(ctx context.Context, entity *entities.ClimateDataEntity) (err error)
}
