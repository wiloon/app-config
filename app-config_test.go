package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestAppConfigBool(t *testing.T) {
	foo := GetBool("rssx.dev-mode")
	fmt.Println(foo)

}

func TestAppConfigString(t *testing.T) {
	log.Println("group0.k0: " + GetString("group0.k0", ""))
}

func TestAppConfigInt(t *testing.T) {
	log.Printf("group0.k1:  %v", GetInt("group0.k1"))
	//log.Println(GetIntWithDefaultValue("k2", 9))
}

func TestAppConfigWithDefaultValue(t *testing.T) {
	log.Println(GetString("k0", ""))
}

//func TestAppConfigBool(t *testing.T) {
//	log.Println("result:", GetBool("k10"))
//	log.Println("result:", GetBoolWithDefaultValue("k11", true))
//
//}

func TestAppConfigTest(t *testing.T) {
	LoadLocalConfig("test.conf")
	log.Println(GetString("k0", ""))
}

func TestToml(t *testing.T) {
	b, err := ioutil.ReadFile("/home/wiloon/projects/pingd-config/config.toml")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	config, _ := toml.Load(str)
	// retrieve data directly
	user := config.Get("group0.k0").(string)
	fmt.Println(user)
	// or using an intermediate object
	postgresConfig := config.Get("group0").(*toml.Tree)
	password := postgresConfig.Get("k1").(int64)
	fmt.Println(password)

	s := config.Get("group0.k9")
	ss := s.(string)
	fmt.Println(ss)
}

func TestSysEnv(t *testing.T) {
	_ = os.Setenv("k0", "v0")
	value := os.Getenv("k0")
	fmt.Println(value)
}
