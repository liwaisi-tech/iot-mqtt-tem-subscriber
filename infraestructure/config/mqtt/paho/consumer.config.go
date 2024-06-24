package config

import (
	"context"
	"sync"

	mqtthandler "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/primary/handlers/mqtt/paho"
	climatedatarepo "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/repositories/postgres/climate_data"
	iotdevicerepo "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/repositories/postgres/iot_device"
	iotdevicesrv "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/services/postgres/iot_device"
	usecases "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/application"
	ports "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/events/mqtt"
	usecaseports "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/usecases"
	gormpkg "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/pkg/gorm/postgres"
)

var (
	consumer ports.MQTTPorts
	once     sync.Once
)

func GetMQTTConsumer() ports.MQTTPorts {
	once.Do(func() {
		consumer = mqtthandler.New(context.Background(), getClimateDataUseCase())
	})
	return consumer
}

func getClimateDataUseCase() usecaseports.SaveClimateDataUseCasePort {
	gormDB := gormpkg.NewPostgresDBConnection()

	climateDataRepo := climatedatarepo.New(gormDB)
	iotDeviceRepo := iotdevicerepo.New(gormDB)

	iotDeviceService := iotdevicesrv.New(gormDB)

	return usecases.New(climateDataRepo, iotDeviceRepo, iotDeviceService)
}
