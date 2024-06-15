package zerolog

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoadLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
	})
}
