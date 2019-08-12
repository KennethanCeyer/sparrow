<h1 align="center">tit</h1>
<p align="center">:bird: Golang env based configure management module which is supported yaml, toml, json</p>

## Getting Started

### What is tit

tit is configure managing tool on Golang. When you need to separate configure file as file (toml, yml, json). This module easily read and parse them to pass the content to Golang code.

And tit supports environment isolation, When you use multiple stage for production, alpha and local. tit will overwrite specific env configuration to general configuration. 

## Quick Start

### Installation

In your terminal

```bash
go get github.com/KennethanCeyer/tit
```

In your code

```go
import "github.com/KennethanCeyer/tit"

var config map[string]interface{}

err := tit.ReadFile("./your_config_path.yml", &config)
```

### Example


> YAML file
```yaml
appName: 'tit'
appVersion: '1.0.0'
debug: true

```

> JSON file
```json
{
  "appName": "tit",
  "appVersion": "1.0.0",
  "debug": true
}
```

> TOML file
```toml
appName = "tit"
appVersion = "1.0.0"
debug = true
```

```go
import (
	"github.com/KennethanCeyer/tit/tit"
	"log"
)

type Config struct {
	AppName    string
	AppVersion string
	Debug      bool
}

func main() {
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
```

**OUTPUT**

```bash
2019/08/12 16:54:02 ymlConfig {tit 1.0.0 true}
2019/08/12 16:54:02 jsonConfig {tit 1.0.0 true}
2019/08/12 16:54:02 tomlConfig {tit 1.0.0 true}
```

For your more information, Please check [the example](./tree/master/example)

## Roadmap

### Configuration parser

- [x] YAML
- [x] JSON
- [x] TOML

### Documentation

- [x] Example
- [ ] API Guide
