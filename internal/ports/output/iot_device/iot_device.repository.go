package ports

import (
	"context"

	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"
)

type IOTDeviceRepository interface {
	CreateDevice(ctx context.Context, entity *entities.IOTDeviceEntity) (err error)
}
