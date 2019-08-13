package main

import (
	"github.com/KennethanCeyer/tit/tit"
	"log"
)

// go run example/tagging/main.go
func InitConfig() error {
	var ymlConfig map[string]interface{}
	var jsonConfig map[string]interface{}
	var tomlConfig map[string]interface{}
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
	log.Println("tomlConfig", tomlConfig)

	return nil
}

func main() {
	err := InitConfig()
	if err != nil {
		log.Panic(err)
	}
}
