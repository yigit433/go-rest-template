package configs

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("⚠️ Failed to load .env file! Please check that the file exists and that read permissions are set correctly.")
		os.Exit(1)
		return
	}

	log.Info().Msg("🔧 .env file successfully loaded and configurations applied!")
}