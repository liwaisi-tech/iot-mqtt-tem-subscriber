package repositories

import (
	"context"

	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"
)

type IOTDeviceRepositoryPort interface {
	CreateIOTDevice(ctx context.Context, entity *entities.IOTDeviceEntity) (newID string, err error)
}
