package main

import (
	"github.com/KennethanCeyer/tit/tit"
	"log"
)

type Config struct {
	Host      string `tit:"host"`
	IPAddr    string `tit:"ip"`
	Port      int    `tit:"port"`
	Admin     string `tit:"admin"`
	Id        string `tit:"id"`
	Pw        string `tit:"pw"`
	Encrypt   string `tit:"encrypt"`
	Token     string `tit:"token"`
	Expires   int    `tit:"expires"`
	Anonymous bool   `tit:"anonymous"`
	Debug     bool   `tit:"debug"`
}

// go run example/tagging/main.go
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
