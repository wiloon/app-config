package config

import (
	"log"
	"github.com/go-akka/configuration"
	"path/filepath"
	"os"
	"strconv"
	"strings"
)

const sysEnvKeyAppConfig = "app_conf"

var defaultFileName = "app.conf"

var configFilePath string
var conf *configuration.Config

func init() {
	LoadLocalConfig(defaultFileName)
}
func LoadLocalConfig(configFileName string) {
	defaultFileName = configFileName

	configFilePath = configPath()
	if !isFileExist(configFilePath) {
		log.Fatal("conifg file not found:", configFilePath)
		//os.Exit(0)
	}
	conf = configuration.LoadConfig(configFilePath)
}
func configPath() string {
	configPath := os.Getenv(sysEnvKeyAppConfig)
	if strings.EqualFold(configPath, "") || !isFileExist(getConfigFilePath(configPath)) {
		log.Printf("config file not found, sys env key:%v, value:%v", sysEnvKeyAppConfig, configPath)

		configPath = execPath()
		if strings.EqualFold(configPath, "") || !isFileExist(getConfigFilePath(configPath)) {
			configPath = currentPath()
		}
	}
	configPath = filepath.Join(configPath, defaultFileName)
	log.Println("config file path:", configPath)
	return configPath
}

func getConfigFilePath(configPath string) string {
	return filepath.Join(configPath, defaultFileName)
}

func GetInt(key string) int {
	return GetIntWithDefaultValue(key, -1)
}

func GetBool(key string) bool {
	return GetBoolWithDefaultValue(key, false)
}

func GetStringList(key string) []string {
	return conf.GetStringList(key)
}

func GetString(key string, def string) string {
	var value string

	value = conf.GetString(key);
	log.Printf("key: %s, value: %s", key, value)

	if value == "" {
		value = def
	}
	return value
}

func GetIntWithDefaultValue(key string, def int) int {
	var value string
	value = conf.GetString(key);
	log.Printf("key: %s, value: %s", key, value)

	value = conf.GetString(key);
	var result int
	if value == "" {
		result = def
	} else {
		intValue, _ := strconv.Atoi(value)
		result = intValue
	}
	return result
}

func GetBoolWithDefaultValue(key string, def bool) bool {
	var value string
	var result bool

	value = conf.GetString(key);
	if value == "" {
		result = def
	} else {
		intValue, _ := strconv.ParseBool(value)
		result = intValue
	}
	return result
}

func isFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	fileExist := err == nil || os.IsExist(err)

	log.Printf("file: %s, exist:%v", filePath, fileExist)
	return fileExist
}

func execPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func currentPath() string {
	currentPath, _ := os.Getwd()
	return currentPath
}
