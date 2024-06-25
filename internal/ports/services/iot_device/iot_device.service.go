package services

import (
	"context"

	entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"
)

type IOTDeviceServicePort interface {
	GetIOTDeviceByMACAddress(ctx context.Context, macAddress string) (iotDevice *entities.IOTDeviceEntity, err error)
}
