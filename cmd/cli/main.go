package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	httpconfig "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/config/http/echo"
	mqttconfig "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/config/mqtt/paho"
	migrations "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/migrations/postgres"
	zerologpkg "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/pkg/zerolog"
)

func main() {
	zerologpkg.LoadLogger()
	migrations.RunMigrations()
	go mqttconfig.GetMQTTConsumer().RunConsumer(os.Getenv("MQTT_TOPIC"))
	httpconfig.NewEchoAPIRestAdapter().RunServer()
}
