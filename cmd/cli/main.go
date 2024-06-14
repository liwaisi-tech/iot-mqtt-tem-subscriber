package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	migrations "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/infraestructure/database/migrations/postgres"
	zerologpkg "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/pkg/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerologpkg.LoadLogger()
	migrations.RunMigrations()
	log.Info().Str("ENV", os.Getenv("ENV")).Msg("Loaded ENV variables")
}
