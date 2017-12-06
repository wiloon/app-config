package config

import (
	"log"
	"github.com/go-akka/configuration"
	"path/filepath"
	"os"
)

const SYS_ENV_KEY_APP_CONFIG = "app_config"
const DEFAULT_FILE_NAME = "app.config"

func GetString(key string, defaultValue string) string {
	var value string

	appConfigPath := os.Getenv(SYS_ENV_KEY_APP_CONFIG)
	if appConfigPath != "" {
		fullPath := filepath.Join(appConfigPath, DEFAULT_FILE_NAME)
		log.Println("app config full path:", fullPath)

		conf := configuration.LoadConfig(fullPath)
		value := conf.GetString(key);
		log.Printf("key: %s, value: %s", key, value)
	} else {
		log.Fatalf("sys env '%s' not found", SYS_ENV_KEY_APP_CONFIG)
	}

	if value == "" {
		value = defaultValue
	}
	return value

}
