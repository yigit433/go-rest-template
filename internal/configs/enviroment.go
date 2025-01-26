package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/yigit433/go-rest-template/internal/logging"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		logging.CustomLogger.Error("⚠️ '.env' file is missing or cannot be opened.", map[string]interface{}{"eventId": "config_load"})
		os.Exit(1)
		return
	} 

    logging.CustomLogger.Info("🔧 '.env' file loaded and settings applied seamlessly!", map[string]interface{}{"eventId": "config_load"})
}