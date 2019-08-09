package main

import (
	"log"
	"tit/tit"
)

type Config struct {
	AppName    string
	AppVersion string
	Debug      bool
}

// go run example/basic/main.go
func InitConfig() error {
	var ymlConfig Config
	err := tit.ReadFile("./config.yml", &ymlConfig)
	if err != nil {
		return err
	}
	log.Println("ymlConfig", ymlConfig)
	return nil
}

func main() {
	err := InitConfig()
	if err != nil {
		log.Panic(err)
	}
}
