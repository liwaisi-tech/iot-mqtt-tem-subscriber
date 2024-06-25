package config

import (
	"context"
	"os"
	"sync"

	mqtthandler "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/primary/handlers/mqtt/paho"
	climatedatarepo "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/repositories/postgres/climate_data"
	iotdevicerepo "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/repositories/postgres/iot_device"
	iotdevicesrv "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/services/postgres/iot_device"
	usecasessave "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/application/usecases/save_climate_data"
	ports "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/events/mqtt"
	gormpkg "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/pkg/gorm/postgres"
)

var (
	mapConsumers map[string]ports.MQTTPorts
	once         sync.Once
)

func InitConsumers() {
	once.Do(func() {
		mapConsumers = map[string]ports.MQTTPorts{
			os.Getenv("MQTT_TOPIC"): getClimateDataConsumer(),
		}
	})
}

func RunConsumers() {
	for topic, consumer := range mapConsumers {
		go consumer.RunConsumer(topic)
	}
}

func getClimateDataConsumer() ports.MQTTPorts {
	gormDB := gormpkg.NewPostgresDBConnection()

	climateDataRepo := climatedatarepo.New(gormDB)
	iotDeviceRepo := iotdevicerepo.New(gormDB)

	iotDeviceService := iotdevicesrv.New(gormDB)

	useCase := usecasessave.New(climateDataRepo, iotDeviceRepo, iotDeviceService)
	return mqtthandler.New(context.Background(), useCase.Execute)
}
