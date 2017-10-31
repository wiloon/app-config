package config

import (
	"log"
	"github.com/go-akka/configuration"
	"path/filepath"
	"os"
)

const sysEnvKeyAppConfig = "app_conf"
const defaultFileName = "app.conf"

func GetString(key string) string {
	return GetStringWithDefaultValue(key, "")

}

func GetStringWithDefaultValue(key string, defaultValue string) string {
	var value string

	appConfigPath := os.Getenv(sysEnvKeyAppConfig)

	var fullPath string
	if appConfigPath == "" {
		log.Printf("sys env key '%s' not found", sysEnvKeyAppConfig)
	}
	fullPath = filepath.Join(appConfigPath, defaultFileName)
	if isFileExist(fullPath) {
		conf := configuration.LoadConfig(fullPath)
		value := conf.GetString(key);
		log.Printf("key: %s, value: %s", key, value)
	} else {
		log.Printf("config file '%s' not found", fullPath)
	}

	if value == "" {
		value = defaultValue
	}
	return value

}

func isFileExist(filename string) bool {
	_, err := os.Stat(filename)
	fileExist := err == nil || os.IsExist(err)
	currentPath, _ := os.Getwd()
	log.Printf("file: %s, exist:%v", currentPath+string(os.PathSeparator)+filename, fileExist)
	return fileExist
}
