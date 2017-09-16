package config

import (
	"testing"
	"fmt"
)

func TestAppConfig(t *testing.T) {
	fmt.Println(GetString("k0"))
	fmt.Println(GetStringWithDefaultValue("k1", ""))
}
