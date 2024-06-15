package config

import (
	"sync"

	mqtthandler "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/primary/handlers/mqtt"
	repositories "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/adapters/secondary/repositories/postgres/climate_data"
	ports "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/event/mqtt"
	gormpkg "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/pkg/gorm/postgres"
)

var (
	consumer ports.MQTTPorts
	once     sync.Once
)

func GetMQTTConsumer() ports.MQTTPorts {
	once.Do(func() {

		climateDataRepository := repositories.New(gormpkg.NewPostgresDBConnection())
		consumer = mqtthandler.New(climateDataRepository)
	})
	return consumer
}
