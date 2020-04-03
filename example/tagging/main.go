package main

import (
	"github.com/KennethanCeyer/sparrow/sparrow"
	"log"
)

type Server struct {
	Host Host `sparrow:"host"`
}

type Host struct {
	IPAddr string `sparrow:"ip"`
	Port   int    `sparrow:"port"`
}

type Credential struct {
	Admin Admin `sparrow:"admin"`
}

type Admin struct {
	Id      string `sparrow:"id"`
	Pw      string `sparrow:"pw"`
	Encrypt string `sparrow:"encrypt"`
}

type Auth struct {
	Token     string `sparrow:"token"`
	Expires   int    `sparrow:"expires"`
	Anonymous bool   `sparrow:"anonymous"`
}

type Config struct {
	Server     Server     `sparrow:"server"`
	Credential Credential `sparrow:"credential"`
	Auth       Auth       `sparrow:"auth"`
	Debug      bool       `sparrow:"debug"`
}

// go run example/tagging/main.go
func InitConfig() error {
	var ymlConfig Config
	var jsonConfig Config
	var tomlConfig Config
	var err error

	err = sparrow.ReadFile("./config.yml", &ymlConfig)
	if err != nil {
		return err
	}

	err = sparrow.ReadFile("./config.json", &jsonConfig)
	if err != nil {
		return err
	}

	err = sparrow.ReadFile("./config.toml", &tomlConfig)
	if err != nil {
		return err
	}

	log.Println("ymlConfig", ymlConfig)
	log.Println("jsonConfig", jsonConfig)
	log.Println("tomlConfig", tomlConfig)

	return nil
}

func main() {
	err := InitConfig()
	if err != nil {
		log.Panic(err)
	}
}
