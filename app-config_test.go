package config

import (
	"testing"
	"log"
	"fmt"
)

func TestAppConfig(t *testing.T) {
	fmt.Println(GetString("k0"))
}

func TestAppConfigWithDefaultValue(t *testing.T) {
	log.Println(GetString("k0"))
	log.Println(GetStringWithDefaultValue("k1", ""))

}
