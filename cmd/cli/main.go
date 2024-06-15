package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	config "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/config/mqtt"
	migrations "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/migrations/postgres"
	zerologpkg "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/pkg/zerolog"
)

func main() {
	zerologpkg.LoadLogger()
	migrations.RunMigrations()
	config.GetMQTTConsumer().RunConsumer(os.Getenv("MQTT_TOPIC"))
}
