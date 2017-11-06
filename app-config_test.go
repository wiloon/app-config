package config

import (
	"testing"
	"log"
)

func TestAppConfig(t *testing.T) {
	log.Println(GetString("k0"))
}

func TestAppConfigInt(t *testing.T) {
	log.Println(GetInt("k1"))
	log.Println(GetIntWithDefaultValue("k2", 9))
}

func TestAppConfigWithDefaultValue(t *testing.T) {
	log.Println(GetString("k0"))
	log.Println(GetStringWithDefaultValue("k1", ""))

}

func TestAppConfigBool(t *testing.T) {
	log.Println("result:", GetBool("k10"))
	log.Println("result:", GetBoolWithDefaultValue("k11", true))

}

func TestAppConfigTest(t *testing.T) {
	LoadLocalConfig("test.conf")
	log.Println(GetString("k0"))
}
