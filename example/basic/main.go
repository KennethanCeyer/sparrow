package main

import (
	"github.com/KennethanCeyer/tit/tit"
	"log"
)

type Config struct {
	AppName    string
	AppVersion string
	Debug      bool
}

// go run example/basic/main.go
func InitConfig() error {
	var ymlConfig Config
	var jsonConfig Config
	var tomlConfig Config
	var err error

	err = tit.ReadFile("./config.yml", &ymlConfig)
	if err != nil {
		return err
	}

	err = tit.ReadFile("./config.json", &jsonConfig)
	if err != nil {
		return err
	}

	err = tit.ReadFile("./config.toml", &tomlConfig)
	if err != nil {
		return err
	}

	log.Println("ymlConfig", ymlConfig)
	log.Println("jsonConfig", jsonConfig)
	log.Println("tomlConfig", jsonConfig)

	return nil
}

func main() {
	err := InitConfig()
	if err != nil {
		log.Panic(err)
	}
}
