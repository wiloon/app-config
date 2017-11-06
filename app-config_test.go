package config

import (
	"testing"
	"log"
	"fmt"
)

func TestAppConfig(t *testing.T) {
	InitConfig()
	fmt.Println(GetString("k0"))
}

func TestAppConfigInt(t *testing.T) {
	fmt.Println(GetInt("k1"))
	fmt.Println(GetIntWithDefaultValue("k2", 9))
}

func TestAppConfigWithDefaultValue(t *testing.T) {
	log.Println(GetString("k0"))
	log.Println(GetStringWithDefaultValue("k1", ""))

}

func TestAppConfigBool(t *testing.T) {
	log.Println("result:", GetBool("k10"))
	log.Println("result:", GetBoolWithDefaultValue("k11", true))

}
