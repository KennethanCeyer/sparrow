package main

import (
	"github.com/KennethanCeyer/tit/tit"
	"log"
)

type Server struct {
	Host Host `tit:"host"`
}

type Host struct {
	IPAddr string `tit:"ip"`
	Port   int    `tit:"port"`
}

type Credential struct {
	Admin Admin `tit:"admin"`
}

type Admin struct {
	Id      string `tit:"id"`
	Pw      string `tit:"pw"`
	Encrypt string `tit:"encrypt"`
}

type Auth struct {
	Token     string `tit:"token"`
	Expires   int    `tit:"expires"`
	Anonymous bool   `tit:"anonymous"`
}

type Config struct {
	Server     Server     `tit:"server"`
	Credential Credential `tit:"credential"`
	Auth       Auth       `tit:"auth"`
	Debug      bool       `tit:"debug"`
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
	log.Println("tomlConfig", tomlConfig)

	return nil
}

func main() {
	err := InitConfig()
	if err != nil {
		log.Panic(err)
	}
}
