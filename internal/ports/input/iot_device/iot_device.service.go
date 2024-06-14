package ports

import entities "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/domain/entities/iot_device"

type IOTDeviceService interface {
	GetDeviceByMACAddress(macAddress string) (entities.IOTDeviceEntity, error)
}
